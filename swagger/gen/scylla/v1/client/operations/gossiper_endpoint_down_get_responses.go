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

// GossiperEndpointDownGetReader is a Reader for the GossiperEndpointDownGet structure.
type GossiperEndpointDownGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GossiperEndpointDownGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGossiperEndpointDownGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGossiperEndpointDownGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGossiperEndpointDownGetOK creates a GossiperEndpointDownGetOK with default headers values
func NewGossiperEndpointDownGetOK() *GossiperEndpointDownGetOK {
	return &GossiperEndpointDownGetOK{}
}

/*GossiperEndpointDownGetOK handles this case with default header values.

GossiperEndpointDownGetOK gossiper endpoint down get o k
*/
type GossiperEndpointDownGetOK struct {
	Payload []string
}

func (o *GossiperEndpointDownGetOK) GetPayload() []string {
	return o.Payload
}

func (o *GossiperEndpointDownGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGossiperEndpointDownGetDefault creates a GossiperEndpointDownGetDefault with default headers values
func NewGossiperEndpointDownGetDefault(code int) *GossiperEndpointDownGetDefault {
	return &GossiperEndpointDownGetDefault{
		_statusCode: code,
	}
}

/*GossiperEndpointDownGetDefault handles this case with default header values.

internal server error
*/
type GossiperEndpointDownGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the gossiper endpoint down get default response
func (o *GossiperEndpointDownGetDefault) Code() int {
	return o._statusCode
}

func (o *GossiperEndpointDownGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GossiperEndpointDownGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *GossiperEndpointDownGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
