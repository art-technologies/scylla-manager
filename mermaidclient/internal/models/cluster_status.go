// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatus cluster status
// swagger:model clusterStatus
type ClusterStatus []*ClusterStatusItems0

// Validate validates this cluster status
func (m ClusterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ClusterStatusItems0 cluster status items0
// swagger:model ClusterStatusItems0
type ClusterStatusItems0 struct {

	// api rtt ms
	APIRttMs float32 `json:"api_rtt_ms,omitempty"`

	// api status
	APIStatus string `json:"api_status,omitempty"`

	// cql rtt ms
	CqlRttMs float32 `json:"cql_rtt_ms,omitempty"`

	// cql status
	CqlStatus string `json:"cql_status,omitempty"`

	// dc
	Dc string `json:"dc,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`
}

// Validate validates this cluster status items0
func (m *ClusterStatusItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatusItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatusItems0) UnmarshalBinary(b []byte) error {
	var res ClusterStatusItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
