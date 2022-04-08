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

// FindConfigMaxHintWindowInMsReader is a Reader for the FindConfigMaxHintWindowInMs structure.
type FindConfigMaxHintWindowInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMaxHintWindowInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMaxHintWindowInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMaxHintWindowInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMaxHintWindowInMsOK creates a FindConfigMaxHintWindowInMsOK with default headers values
func NewFindConfigMaxHintWindowInMsOK() *FindConfigMaxHintWindowInMsOK {
	return &FindConfigMaxHintWindowInMsOK{}
}

/*FindConfigMaxHintWindowInMsOK handles this case with default header values.

Config value
*/
type FindConfigMaxHintWindowInMsOK struct {
	Payload int64
}

func (o *FindConfigMaxHintWindowInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMaxHintWindowInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMaxHintWindowInMsDefault creates a FindConfigMaxHintWindowInMsDefault with default headers values
func NewFindConfigMaxHintWindowInMsDefault(code int) *FindConfigMaxHintWindowInMsDefault {
	return &FindConfigMaxHintWindowInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigMaxHintWindowInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigMaxHintWindowInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config max hint window in ms default response
func (o *FindConfigMaxHintWindowInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMaxHintWindowInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMaxHintWindowInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMaxHintWindowInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
