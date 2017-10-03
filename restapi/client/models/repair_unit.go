// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RepairUnit repair unit
// swagger:model repairUnit

type RepairUnit struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// keyspace
	Keyspace string `json:"keyspace,omitempty"`

	// tables
	Tables []string `json:"tables"`
}

/* polymorph repairUnit cluster_id false */

/* polymorph repairUnit id false */

/* polymorph repairUnit keyspace false */

/* polymorph repairUnit tables false */

// Validate validates this repair unit
func (m *RepairUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTables(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepairUnit) validateTables(formats strfmt.Registry) error {

	if swag.IsZero(m.Tables) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepairUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepairUnit) UnmarshalBinary(b []byte) error {
	var res RepairUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
