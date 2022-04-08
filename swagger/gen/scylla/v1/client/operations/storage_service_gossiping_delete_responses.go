// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceGossipingDeleteReader is a Reader for the StorageServiceGossipingDelete structure.
type StorageServiceGossipingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceGossipingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceGossipingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceGossipingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceGossipingDeleteOK creates a StorageServiceGossipingDeleteOK with default headers values
func NewStorageServiceGossipingDeleteOK() *StorageServiceGossipingDeleteOK {
	return &StorageServiceGossipingDeleteOK{}
}

/*StorageServiceGossipingDeleteOK handles this case with default header values.

StorageServiceGossipingDeleteOK storage service gossiping delete o k
*/
type StorageServiceGossipingDeleteOK struct {
}

func (o *StorageServiceGossipingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceGossipingDeleteDefault creates a StorageServiceGossipingDeleteDefault with default headers values
func NewStorageServiceGossipingDeleteDefault(code int) *StorageServiceGossipingDeleteDefault {
	return &StorageServiceGossipingDeleteDefault{
		_statusCode: code,
	}
}

/*StorageServiceGossipingDeleteDefault handles this case with default header values.

internal server error
*/
type StorageServiceGossipingDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service gossiping delete default response
func (o *StorageServiceGossipingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceGossipingDeleteDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceGossipingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceGossipingDeleteDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
