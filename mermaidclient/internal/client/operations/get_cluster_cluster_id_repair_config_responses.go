// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/mermaidclient/internal/models"
)

// GetClusterClusterIDRepairConfigReader is a Reader for the GetClusterClusterIDRepairConfig structure.
type GetClusterClusterIDRepairConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDRepairConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterClusterIDRepairConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		body := response.Body()
		defer body.Close()

		var m json.RawMessage
		if err := json.NewDecoder(body).Decode(&m); err != nil {
			return nil, err
		}

		return nil, runtime.NewAPIError("API error", m, response.Code())
	}
}

// NewGetClusterClusterIDRepairConfigOK creates a GetClusterClusterIDRepairConfigOK with default headers values
func NewGetClusterClusterIDRepairConfigOK() *GetClusterClusterIDRepairConfigOK {
	return &GetClusterClusterIDRepairConfigOK{}
}

/*GetClusterClusterIDRepairConfigOK handles this case with default header values.

cluster config fields
*/
type GetClusterClusterIDRepairConfigOK struct {
	Payload *models.RepairConfig
}

func (o *GetClusterClusterIDRepairConfigOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/repair/config][%d] getClusterClusterIdRepairConfigOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDRepairConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepairConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
