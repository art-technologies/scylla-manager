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

// FindConfigWriteRequestTimeoutInMsReader is a Reader for the FindConfigWriteRequestTimeoutInMs structure.
type FindConfigWriteRequestTimeoutInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigWriteRequestTimeoutInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigWriteRequestTimeoutInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigWriteRequestTimeoutInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigWriteRequestTimeoutInMsOK creates a FindConfigWriteRequestTimeoutInMsOK with default headers values
func NewFindConfigWriteRequestTimeoutInMsOK() *FindConfigWriteRequestTimeoutInMsOK {
	return &FindConfigWriteRequestTimeoutInMsOK{}
}

/*FindConfigWriteRequestTimeoutInMsOK handles this case with default header values.

Config value
*/
type FindConfigWriteRequestTimeoutInMsOK struct {
	Payload int64
}

func (o *FindConfigWriteRequestTimeoutInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigWriteRequestTimeoutInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigWriteRequestTimeoutInMsDefault creates a FindConfigWriteRequestTimeoutInMsDefault with default headers values
func NewFindConfigWriteRequestTimeoutInMsDefault(code int) *FindConfigWriteRequestTimeoutInMsDefault {
	return &FindConfigWriteRequestTimeoutInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigWriteRequestTimeoutInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigWriteRequestTimeoutInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config write request timeout in ms default response
func (o *FindConfigWriteRequestTimeoutInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigWriteRequestTimeoutInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigWriteRequestTimeoutInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigWriteRequestTimeoutInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
