// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v2/models"
)

// FindConfigAlternatorAddressReader is a Reader for the FindConfigAlternatorAddress structure.
type FindConfigAlternatorAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAlternatorAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAlternatorAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAlternatorAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAlternatorAddressOK creates a FindConfigAlternatorAddressOK with default headers values
func NewFindConfigAlternatorAddressOK() *FindConfigAlternatorAddressOK {
	return &FindConfigAlternatorAddressOK{}
}

/*FindConfigAlternatorAddressOK handles this case with default header values.

Config value
*/
type FindConfigAlternatorAddressOK struct {
	Payload string
}

func (o *FindConfigAlternatorAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigAlternatorAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAlternatorAddressDefault creates a FindConfigAlternatorAddressDefault with default headers values
func NewFindConfigAlternatorAddressDefault(code int) *FindConfigAlternatorAddressDefault {
	return &FindConfigAlternatorAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigAlternatorAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigAlternatorAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config alternator address default response
func (o *FindConfigAlternatorAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAlternatorAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAlternatorAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAlternatorAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
