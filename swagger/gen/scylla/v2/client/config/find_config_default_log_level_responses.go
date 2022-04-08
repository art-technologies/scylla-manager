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

// FindConfigDefaultLogLevelReader is a Reader for the FindConfigDefaultLogLevel structure.
type FindConfigDefaultLogLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDefaultLogLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDefaultLogLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDefaultLogLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDefaultLogLevelOK creates a FindConfigDefaultLogLevelOK with default headers values
func NewFindConfigDefaultLogLevelOK() *FindConfigDefaultLogLevelOK {
	return &FindConfigDefaultLogLevelOK{}
}

/*FindConfigDefaultLogLevelOK handles this case with default header values.

Config value
*/
type FindConfigDefaultLogLevelOK struct {
	Payload string
}

func (o *FindConfigDefaultLogLevelOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigDefaultLogLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDefaultLogLevelDefault creates a FindConfigDefaultLogLevelDefault with default headers values
func NewFindConfigDefaultLogLevelDefault(code int) *FindConfigDefaultLogLevelDefault {
	return &FindConfigDefaultLogLevelDefault{
		_statusCode: code,
	}
}

/*FindConfigDefaultLogLevelDefault handles this case with default header values.

unexpected error
*/
type FindConfigDefaultLogLevelDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config default log level default response
func (o *FindConfigDefaultLogLevelDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDefaultLogLevelDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDefaultLogLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDefaultLogLevelDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
