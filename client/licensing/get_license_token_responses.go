// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetLicenseTokenReader is a Reader for the GetLicenseToken structure.
type GetLicenseTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /licensing/token] getLicenseToken", response, response.Code())
	}
}

// NewGetLicenseTokenOK creates a GetLicenseTokenOK with default headers values
func NewGetLicenseTokenOK() *GetLicenseTokenOK {
	return &GetLicenseTokenOK{}
}

/*
GetLicenseTokenOK describes a response with status code 200, with default header values.

(empty)
*/
type GetLicenseTokenOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this get license token o k response has a 2xx status code
func (o *GetLicenseTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get license token o k response has a 3xx status code
func (o *GetLicenseTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license token o k response has a 4xx status code
func (o *GetLicenseTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license token o k response has a 5xx status code
func (o *GetLicenseTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get license token o k response a status code equal to that given
func (o *GetLicenseTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get license token o k response
func (o *GetLicenseTokenOK) Code() int {
	return 200
}

func (o *GetLicenseTokenOK) Error() string {
	return fmt.Sprintf("[GET /licensing/token][%d] getLicenseTokenOK  %+v", 200, o.Payload)
}

func (o *GetLicenseTokenOK) String() string {
	return fmt.Sprintf("[GET /licensing/token][%d] getLicenseTokenOK  %+v", 200, o.Payload)
}

func (o *GetLicenseTokenOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *GetLicenseTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
