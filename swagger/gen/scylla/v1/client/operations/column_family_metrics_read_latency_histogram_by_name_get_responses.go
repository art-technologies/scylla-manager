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

// ColumnFamilyMetricsReadLatencyHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsReadLatencyHistogramByNameGet structure.
type ColumnFamilyMetricsReadLatencyHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsReadLatencyHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK creates a ColumnFamilyMetricsReadLatencyHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK() *ColumnFamilyMetricsReadLatencyHistogramByNameGetOK {
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsReadLatencyHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyHistogramByNameGetOK column family metrics read latency histogram by name get o k
*/
type ColumnFamilyMetricsReadLatencyHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetDefault creates a ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetDefault(code int) *ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault {
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics read latency histogram by name get default response
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
