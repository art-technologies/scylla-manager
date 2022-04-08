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

// FindConfigMaxHintsDeliveryThreadsReader is a Reader for the FindConfigMaxHintsDeliveryThreads structure.
type FindConfigMaxHintsDeliveryThreadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMaxHintsDeliveryThreadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMaxHintsDeliveryThreadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMaxHintsDeliveryThreadsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMaxHintsDeliveryThreadsOK creates a FindConfigMaxHintsDeliveryThreadsOK with default headers values
func NewFindConfigMaxHintsDeliveryThreadsOK() *FindConfigMaxHintsDeliveryThreadsOK {
	return &FindConfigMaxHintsDeliveryThreadsOK{}
}

/*FindConfigMaxHintsDeliveryThreadsOK handles this case with default header values.

Config value
*/
type FindConfigMaxHintsDeliveryThreadsOK struct {
	Payload int64
}

func (o *FindConfigMaxHintsDeliveryThreadsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMaxHintsDeliveryThreadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMaxHintsDeliveryThreadsDefault creates a FindConfigMaxHintsDeliveryThreadsDefault with default headers values
func NewFindConfigMaxHintsDeliveryThreadsDefault(code int) *FindConfigMaxHintsDeliveryThreadsDefault {
	return &FindConfigMaxHintsDeliveryThreadsDefault{
		_statusCode: code,
	}
}

/*FindConfigMaxHintsDeliveryThreadsDefault handles this case with default header values.

unexpected error
*/
type FindConfigMaxHintsDeliveryThreadsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config max hints delivery threads default response
func (o *FindConfigMaxHintsDeliveryThreadsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMaxHintsDeliveryThreadsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMaxHintsDeliveryThreadsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMaxHintsDeliveryThreadsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
