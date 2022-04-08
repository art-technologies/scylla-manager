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

// StorageServiceRescheduleFailedDeletionsPostReader is a Reader for the StorageServiceRescheduleFailedDeletionsPost structure.
type StorageServiceRescheduleFailedDeletionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRescheduleFailedDeletionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRescheduleFailedDeletionsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRescheduleFailedDeletionsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRescheduleFailedDeletionsPostOK creates a StorageServiceRescheduleFailedDeletionsPostOK with default headers values
func NewStorageServiceRescheduleFailedDeletionsPostOK() *StorageServiceRescheduleFailedDeletionsPostOK {
	return &StorageServiceRescheduleFailedDeletionsPostOK{}
}

/*StorageServiceRescheduleFailedDeletionsPostOK handles this case with default header values.

StorageServiceRescheduleFailedDeletionsPostOK storage service reschedule failed deletions post o k
*/
type StorageServiceRescheduleFailedDeletionsPostOK struct {
}

func (o *StorageServiceRescheduleFailedDeletionsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceRescheduleFailedDeletionsPostDefault creates a StorageServiceRescheduleFailedDeletionsPostDefault with default headers values
func NewStorageServiceRescheduleFailedDeletionsPostDefault(code int) *StorageServiceRescheduleFailedDeletionsPostDefault {
	return &StorageServiceRescheduleFailedDeletionsPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceRescheduleFailedDeletionsPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceRescheduleFailedDeletionsPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service reschedule failed deletions post default response
func (o *StorageServiceRescheduleFailedDeletionsPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRescheduleFailedDeletionsPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRescheduleFailedDeletionsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRescheduleFailedDeletionsPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
