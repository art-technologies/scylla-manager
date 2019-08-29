// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Cluster cluster
// swagger:model cluster
type Cluster struct {

	// auth token
	AuthToken string `json:"auth_token,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ssl user cert file
	// Format: byte
	SslUserCertFile strfmt.Base64 `json:"ssl_user_cert_file,omitempty"`

	// ssl user key file
	// Format: byte
	SslUserKeyFile strfmt.Base64 `json:"ssl_user_key_file,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSslUserCertFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslUserKeyFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateSslUserCertFile(formats strfmt.Registry) error {

	if swag.IsZero(m.SslUserCertFile) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *Cluster) validateSslUserKeyFile(formats strfmt.Registry) error {

	if swag.IsZero(m.SslUserKeyFile) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
