// Copyright (C) 2017 ScyllaDB

package backup

import (
	"context"
	"net/http"
	"path"
	"sort"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
	"github.com/scylladb/go-set/strset"
	"github.com/scylladb/scylla-manager/pkg/scyllaclient"
	. "github.com/scylladb/scylla-manager/pkg/service/backup/backupspec"
	"github.com/scylladb/scylla-manager/pkg/util/parallel"
	"github.com/scylladb/scylla-manager/pkg/util/timeutc"
	"github.com/scylladb/scylla-manager/pkg/util/uuid"
	"go.uber.org/atomic"
)

// listManifestsInAllLocations returns manifests for all nodes of a in all
// locations specified in hosts.
func listManifestsInAllLocations(ctx context.Context, client *scyllaclient.Client, hosts []hostInfo, clusterID uuid.UUID) ([]*RemoteManifest, error) {
	var (
		locations = make(map[Location]struct{})
		manifests []*RemoteManifest
	)

	for _, hi := range hosts {
		if _, ok := locations[hi.Location]; ok {
			continue
		}
		locations[hi.Location] = struct{}{}

		lm, err := listManifests(ctx, client, hi.IP, hi.Location, clusterID)
		if err != nil {
			return nil, err
		}
		manifests = append(manifests, lm...)
	}

	return manifests, nil
}

// listManifests returns manifests for all nodes of a given cluster in the location.
func listManifests(ctx context.Context, client *scyllaclient.Client, host string, location Location, clusterID uuid.UUID) ([]*RemoteManifest, error) {
	baseDir := RemoteMetaClusterDCDir(clusterID)
	opts := scyllaclient.RcloneListDirOpts{
		FilesOnly: true,
		Recurse:   true,
	}
	items, err := client.RcloneListDir(ctx, host, location.RemotePath(baseDir), &opts)
	if err != nil {
		return nil, err
	}

	var manifests []*RemoteManifest
	for _, item := range items {
		p := path.Join(baseDir, item.Path)
		m := &RemoteManifest{}
		if err := m.ParsePartialPath(p); err != nil {
			continue
		}
		m.Location = location
		manifests = append(manifests, m)
	}
	return manifests, nil
}

func groupManifestsByNode(manifests []*RemoteManifest) map[string][]*RemoteManifest {
	v := map[string][]*RemoteManifest{}
	for _, m := range manifests {
		v[m.NodeID] = append(v[m.NodeID], m)
	}
	return v
}

func groupManifestsByTask(manifests []*RemoteManifest) map[uuid.UUID][]*RemoteManifest {
	v := map[uuid.UUID][]*RemoteManifest{}
	for _, m := range manifests {
		v[m.TaskID] = append(v[m.TaskID], m)
	}
	return v
}

// popNodeIDManifestsForLocation returns a function that for a given location
// finds next node and it's manifests from that location.
func popNodeIDManifestsForLocation(manifests []*RemoteManifest) func(h hostInfo) (string, []*RemoteManifest) {
	var mu sync.Mutex
	nodeIDManifests := groupManifestsByNode(manifests)
	return func(h hostInfo) (string, []*RemoteManifest) {
		mu.Lock()
		defer mu.Unlock()

		// Fast path, get manifests for the current node
		if manifests, ok := nodeIDManifests[h.ID]; ok {
			delete(nodeIDManifests, h.ID)
			return h.ID, manifests
		}

		// Look for other nodes in the same location
		for nodeID, manifests := range nodeIDManifests {
			if manifests[0].Location == h.Location {
				delete(nodeIDManifests, nodeID)
				return nodeID, manifests
			}
		}

		return "", nil
	}
}

// staleTags returns collection of snapshot tags for manifests that are "stale".
// That is:
// - temporary manifests,
// - manifests over task retention policy,
// - manifests older than threshold if retention policy is unknown.
func staleTags(manifests []*RemoteManifest, policy map[uuid.UUID]int, threshold time.Time) *strset.Set {
	tags := strset.New()

	for taskID, taskManifests := range groupManifestsByTask(manifests) {
		taskPolicy := policy[taskID]
		taskTags := strset.New()
		for _, m := range taskManifests {
			switch {
			case m.Temporary:
				tags.Add(m.SnapshotTag)
			case taskPolicy > 0:
				taskTags.Add(m.SnapshotTag)
			default:
				t, _ := SnapshotTagTime(m.SnapshotTag) // nolint: errcheck
				if t.Before(threshold) {
					tags.Add(m.SnapshotTag)
				}
			}
		}
		if taskPolicy > 0 && taskTags.Size()-taskPolicy > 0 {
			l := taskTags.List()
			sort.Strings(l)
			tags.Add(l[:len(l)-taskPolicy]...)
		}
	}
	return tags
}

type purger struct {
	client *scyllaclient.Client
	host   string
	logger log.Logger
	// nodeIP maps node ID to IP based on information from read manifests.
	nodeIP map[string]string

	notifyEach int
	OnScan     func(scanned, orphaned int, orphanedBytes int64)
	OnDelete   func(total, success int)
}

func newPurger(client *scyllaclient.Client, host string, logger log.Logger) purger {
	return purger{
		client: client,
		host:   host,
		logger: logger,
		nodeIP: make(map[string]string),

		notifyEach: 1000,
	}
}

func (p purger) PurgeSnapshotTags(ctx context.Context, manifests []*RemoteManifest, tags *strset.Set) (int, error) {
	if len(manifests) == 0 {
		return 0, nil
	}

	var (
		files = make(fileSet)
		stale = 0

		c ManifestContent
	)

	for _, m := range manifests {
		if tags.Has(m.SnapshotTag) {
			stale++
			p.logger.Info(ctx, "Found manifest to remove",
				"task", m.TaskID,
				"snapshot_tag", m.SnapshotTag,
				"temporary", m.Temporary,
			)
			if err := p.loadManifestContentInto(ctx, m, &c); err != nil {
				return 0, errors.Wrapf(err, "load manifest %s", m.RemoteManifestFile())
			}
			p.forEachDir(m, &c, files.AddFiles)
		}
	}
	if stale == 0 {
		return 0, nil
	}
	for _, m := range manifests {
		if !tags.Has(m.SnapshotTag) {
			if err := p.loadManifestContentInto(ctx, m, &c); err != nil {
				return 0, errors.Wrapf(err, "load manifest %s", m.RemoteManifestFile())
			}
			p.forEachDir(m, &c, files.RemoveFiles)
		}
	}
	if _, err := p.deleteFiles(ctx, manifests[0].Location, files); err != nil {
		return 0, errors.Wrapf(err, "delete")
	}
	deletedManifests := 0
	for _, m := range manifests {
		if tags.Has(m.SnapshotTag) {
			if _, err := p.deleteFile(ctx, m.Location.RemotePath(m.RemoteSchemaFile())); err != nil {
				p.logger.Info(ctx, "Failed to remove schema file", "path", m.RemoteSchemaFile(), "error", err)
			}
			if _, err := p.deleteFile(ctx, m.Location.RemotePath(m.RemoteManifestFile())); err != nil {
				p.logger.Info(ctx, "Failed to remove manifest", "path", m.RemoteManifestFile(), "error", err)
			} else {
				deletedManifests++
			}
		}
	}
	return deletedManifests, nil
}

// ValidationResult is a summary generated by Validate.
type ValidationResult struct {
	ScannedFiles    int      `json:"scanned_files"`
	BrokenSnapshots []string `json:"broken_snapshots"`
	MissingFiles    int      `json:"missing_files"`
	OrphanedFiles   int      `json:"orphaned_files"`
	OrphanedBytes   int64    `json:"orphaned_bytes"`
	DeletedFiles    int      `json:"deleted_files"`
}

func (p purger) Validate(ctx context.Context, manifests []*RemoteManifest, deleteOrphanedFiles bool) (ValidationResult, error) {
	var result ValidationResult

	if len(manifests) == 0 {
		return result, nil
	}

	start := timeutc.Now()

	var (
		files             = make(fileSet)
		tempManifestFiles = make(fileSet)
		orphanedFiles     = make(fileSet)

		c ManifestContent
	)

	for _, m := range manifests {
		if err := p.loadManifestContentInto(ctx, m, &c); err != nil {
			return result, errors.Wrapf(err, "load manifest %s", m.RemoteManifestFile())
		}
		if m.Temporary {
			p.forEachDir(m, &c, tempManifestFiles.AddFiles)
		} else {
			p.forEachDir(m, &c, files.AddFiles)
		}
	}

	handler := func(item *scyllaclient.RcloneListDirItem) {
		result.ScannedFiles++

		// OK, file from manifest
		if files.Has(item.Path) {
			files.Remove(item.Path)
			return
		}
		// OK, file from temporary manifest i.e. running backup
		if tempManifestFiles.Has(item.Path) {
			tempManifestFiles.Remove(item.Path)
			return
		}
		// OK, file added after we started
		if time.Time(item.ModTime).After(start) {
			p.logger.Info(ctx, "Unexpected new file", "path", item.Path, "mod_time", time.Time(item.ModTime))
			return
		}

		// NOK, handle orphaned file
		result.OrphanedFiles++
		result.OrphanedBytes += item.Size
		orphanedFiles.Add(item.Path)

		if result.ScannedFiles%p.notifyEach == 0 {
			p.onScan(ctx, result)
		}
	}
	if err := p.forEachRemoteFile(ctx, manifests[0], handler); err != nil {
		return result, errors.Wrap(err, "list files")
	}

	result.MissingFiles = files.Size()
	p.onScan(ctx, result)

	if result.MissingFiles > 0 {
		p.logger.Info(ctx, "Found missing files, looking for affected manifests")
		if bs, err := p.findBrokenSnapshots(ctx, manifests, files); err != nil {
			p.logger.Error(ctx, "Error while finding broken snapshots", "error", err)
		} else {
			result.BrokenSnapshots = bs
		}
	}

	// Remove orphaned files
	if deleteOrphanedFiles {
		n, err := p.deleteFiles(ctx, manifests[0].Location, orphanedFiles)
		if err != nil {
			return result, errors.Wrapf(err, "delete")
		}
		result.DeletedFiles = n
	}

	return result, nil
}

func (p purger) onScan(ctx context.Context, result ValidationResult) {
	p.logger.Info(ctx, "Scanning files",
		"scanned_files", result.ScannedFiles,
		"orphaned_files", result.OrphanedFiles,
		"orphaned_bytes", result.OrphanedBytes,
	)
	if p.OnScan != nil {
		p.OnScan(result.ScannedFiles, result.OrphanedFiles, result.OrphanedBytes)
	}
}

func (p purger) findBrokenSnapshots(ctx context.Context, manifests []*RemoteManifest, missingFiles fileSet) ([]string, error) {
	if missingFiles.Size() == 0 {
		return nil, nil
	}

	var (
		s = strset.New()
		c ManifestContent
	)
	for _, m := range manifests {
		if m.Temporary {
			continue
		}
		if err := p.loadManifestContentInto(ctx, m, &c); err != nil {
			// Ignore manifests removed while validate was running
			if scyllaclient.StatusCodeOf(err) == http.StatusNotFound {
				continue
			}
			return nil, errors.Wrapf(err, "load manifest %s", m.RemoteManifestFile())
		}
		p.forEachDir(m, &c, func(dir string, files []string) {
			if missingFiles.HasAnyFiles(dir, files) {
				s.Add(m.SnapshotTag)
			}
		})
	}

	v := s.List()
	sort.Strings(v)
	return v, nil
}

func (p purger) loadManifestContentInto(ctx context.Context, m *RemoteManifest, c *ManifestContent) error {
	p.logger.Info(ctx, "Reading manifest",
		"task", m.TaskID,
		"snapshot_tag", m.SnapshotTag,
	)

	*c = ManifestContent{}
	r, err := p.client.RcloneOpen(ctx, p.host, m.Location.RemotePath(m.RemoteManifestFile()))
	if err != nil {
		return err
	}
	defer r.Close()

	defer func() {
		if c.IP != "" {
			p.nodeIP[m.NodeID] = c.IP
		}
	}()

	return c.Read(r)
}

func (p purger) forEachDir(m *RemoteManifest, c *ManifestContent, callback func(dir string, files []string)) {
	for _, fi := range c.Index {
		dir := RemoteSSTableVersionDir(m.ClusterID, m.DC, m.NodeID, fi.Keyspace, fi.Table, fi.Version)
		callback(dir, fi.Files)
	}
}

func (p purger) forEachRemoteFile(ctx context.Context, m *RemoteManifest, f func(*scyllaclient.RcloneListDirItem)) error {
	baseDir := RemoteSSTableBaseDir(m.ClusterID, m.DC, m.NodeID)
	wrapper := func(item *scyllaclient.RcloneListDirItem) {
		item.Path = path.Join(baseDir, item.Path)
		f(item)
	}
	opts := scyllaclient.RcloneListDirOpts{
		FilesOnly: true,
		Recurse:   true,
	}
	return p.client.RcloneListDirIter(ctx, p.host, m.Location.RemotePath(baseDir), &opts, wrapper)
}

func (p purger) deleteFiles(ctx context.Context, location Location, files fileSet) (int, error) {
	if files.Size() == 0 {
		return 0, nil
	}

	var (
		dirs    = files.Dirs()
		total   = files.Size()
		success atomic.Int64
		missing atomic.Int64
	)

	defer func() {
		s, m := int(success.Load()), int(missing.Load())
		p.onDelete(ctx, total, s, m)
	}()

	if err := parallel.Run(len(dirs), parallel.NoLimit, func(i int) error {
		var (
			dir  = dirs[i]
			rerr error
		)

		files.DirSet(dir).Each(func(file string) bool {
			f := path.Join(dir, file)
			ok, err := p.deleteFile(ctx, location.RemotePath(f))
			// On error exit iteration and report error
			if err != nil {
				rerr = errors.Wrapf(err, "file %s", f)
				return false
			}
			s, m := int(success.Inc()), 0
			if !ok {
				m = int(missing.Inc())
			}
			if s%p.notifyEach == 0 {
				if m == 0 {
					m = int(missing.Load())
				}
				p.onDelete(ctx, total, s, m)
			}

			return true
		})

		return rerr
	}); err != nil {
		return int(success.Load()), err
	}

	return int(success.Load()), nil
}

func (p purger) onDelete(ctx context.Context, total, success, missing int) {
	p.logger.Info(ctx, "Deleted files",
		"success", success,
		"missing", missing,
		"total", total,
	)
	if p.OnDelete != nil {
		p.OnDelete(total, success)
	}
}

func (p purger) deleteFile(ctx context.Context, path string) (bool, error) {
	p.logger.Debug(ctx, "Deleting file", "path", path)
	err := p.client.RcloneDeleteFile(ctx, p.host, path)
	if scyllaclient.StatusCodeOf(err) == http.StatusNotFound {
		return false, nil
	}
	return true, err
}

// Host can be called from OnPreDelete and OnDelete callbacks to convert node ID
// to IP for metrics purposes.
func (p purger) Host(nodeID string) string {
	return p.nodeIP[nodeID]
}
