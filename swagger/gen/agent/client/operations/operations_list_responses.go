// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/agent/models"
)

// OperationsListReader is a Reader for the OperationsList structure.
type OperationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsListOK creates a OperationsListOK with default headers values
func NewOperationsListOK() *OperationsListOK {
	return &OperationsListOK{}
}

/*OperationsListOK handles this case with default header values.

List of items
*/
type OperationsListOK struct {
	Payload *OperationsListOKBody
	JobID   int64
}

func (o *OperationsListOK) GetPayload() *OperationsListOKBody {
	return o.Payload
}

func (o *OperationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(OperationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsListDefault creates a OperationsListDefault with default headers values
func NewOperationsListDefault(code int) *OperationsListDefault {
	return &OperationsListDefault{
		_statusCode: code,
	}
}

/*OperationsListDefault handles this case with default header values.

Server error
*/
type OperationsListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations list default response
func (o *OperationsListDefault) Code() int {
	return o._statusCode
}

func (o *OperationsListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsListDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

/*OperationsListOKBody operations list o k body
swagger:model OperationsListOKBody
*/
type OperationsListOKBody struct {

	// list
	List []*models.ListItem `json:"list"`
}

// Validate validates this operations list o k body
func (o *OperationsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OperationsListOKBody) validateList(formats strfmt.Registry) error {

	if swag.IsZero(o.List) { // not required
		return nil
	}

	for i := 0; i < len(o.List); i++ {
		if swag.IsZero(o.List[i]) { // not required
			continue
		}

		if o.List[i] != nil {
			if err := o.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operationsListOK" + "." + "list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *OperationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OperationsListOKBody) UnmarshalBinary(b []byte) error {
	var res OperationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
