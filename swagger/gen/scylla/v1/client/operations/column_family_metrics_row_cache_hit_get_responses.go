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

// ColumnFamilyMetricsRowCacheHitGetReader is a Reader for the ColumnFamilyMetricsRowCacheHitGet structure.
type ColumnFamilyMetricsRowCacheHitGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRowCacheHitGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRowCacheHitGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsRowCacheHitGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsRowCacheHitGetOK creates a ColumnFamilyMetricsRowCacheHitGetOK with default headers values
func NewColumnFamilyMetricsRowCacheHitGetOK() *ColumnFamilyMetricsRowCacheHitGetOK {
	return &ColumnFamilyMetricsRowCacheHitGetOK{}
}

/*ColumnFamilyMetricsRowCacheHitGetOK handles this case with default header values.

ColumnFamilyMetricsRowCacheHitGetOK column family metrics row cache hit get o k
*/
type ColumnFamilyMetricsRowCacheHitGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsRowCacheHitGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheHitGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsRowCacheHitGetDefault creates a ColumnFamilyMetricsRowCacheHitGetDefault with default headers values
func NewColumnFamilyMetricsRowCacheHitGetDefault(code int) *ColumnFamilyMetricsRowCacheHitGetDefault {
	return &ColumnFamilyMetricsRowCacheHitGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsRowCacheHitGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsRowCacheHitGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics row cache hit get default response
func (o *ColumnFamilyMetricsRowCacheHitGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsRowCacheHitGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheHitGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsRowCacheHitGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
