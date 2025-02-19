// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetClusterClusterIDTaskRestoreTaskIDRunIDParams creates a new GetClusterClusterIDTaskRestoreTaskIDRunIDParams object
// with the default values initialized.
func NewGetClusterClusterIDTaskRestoreTaskIDRunIDParams() *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRestoreTaskIDRunIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithTimeout creates a new GetClusterClusterIDTaskRestoreTaskIDRunIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRestoreTaskIDRunIDParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithContext creates a new GetClusterClusterIDTaskRestoreTaskIDRunIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithContext(ctx context.Context) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRestoreTaskIDRunIDParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithHTTPClient creates a new GetClusterClusterIDTaskRestoreTaskIDRunIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDTaskRestoreTaskIDRunIDParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	var ()
	return &GetClusterClusterIDTaskRestoreTaskIDRunIDParams{
		HTTPClient: client,
	}
}

/*GetClusterClusterIDTaskRestoreTaskIDRunIDParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID task restore task ID run ID operation typically these are written to a http.Request
*/
type GetClusterClusterIDTaskRestoreTaskIDRunIDParams struct {

	/*ClusterID*/
	ClusterID string
	/*RunID*/
	RunID string
	/*TaskID*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithContext(ctx context.Context) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithClusterID(clusterID string) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithRunID adds the runID to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithRunID(runID string) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetRunID(runID string) {
	o.RunID = runID
}

// WithTaskID adds the taskID to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WithTaskID(taskID string) *GetClusterClusterIDTaskRestoreTaskIDRunIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get cluster cluster ID task restore task ID run ID params
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDTaskRestoreTaskIDRunIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
