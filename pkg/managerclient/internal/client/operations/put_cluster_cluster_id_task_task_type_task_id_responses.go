// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/managerclient/internal/models"
)

// PutClusterClusterIDTaskTaskTypeTaskIDReader is a Reader for the PutClusterClusterIDTaskTaskTypeTaskID structure.
type PutClusterClusterIDTaskTaskTypeTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDTaskTaskTypeTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDOK creates a PutClusterClusterIDTaskTaskTypeTaskIDOK with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDOK() *PutClusterClusterIDTaskTaskTypeTaskIDOK {
	return &PutClusterClusterIDTaskTaskTypeTaskIDOK{}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDOK handles this case with default header values.

Updated task info
*/
type PutClusterClusterIDTaskTaskTypeTaskIDOK struct {
	Payload *models.Task
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] putClusterClusterIdTaskTaskTypeTaskIdOK  %+v", 200, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDDefault creates a PutClusterClusterIDTaskTaskTypeTaskIDDefault with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDDefault(code int) *PutClusterClusterIDTaskTaskTypeTaskIDDefault {
	return &PutClusterClusterIDTaskTaskTypeTaskIDDefault{
		_statusCode: code,
	}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDDefault handles this case with default header values.

Error
*/
type PutClusterClusterIDTaskTaskTypeTaskIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put cluster cluster ID task task type task ID default response
func (o *PutClusterClusterIDTaskTaskTypeTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDDefault) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] PutClusterClusterIDTaskTaskTypeTaskID default  %+v", o._statusCode, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
