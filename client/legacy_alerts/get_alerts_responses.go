// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetAlertsReader is a Reader for the GetAlerts structure.
type GetAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /alerts] getAlerts", response, response.Code())
	}
}

// NewGetAlertsOK creates a GetAlertsOK with default headers values
func NewGetAlertsOK() *GetAlertsOK {
	return &GetAlertsOK{}
}

/*
GetAlertsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAlertsOK struct {
	Payload []*models.AlertListItemDTO
}

// IsSuccess returns true when this get alerts o k response has a 2xx status code
func (o *GetAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerts o k response has a 3xx status code
func (o *GetAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerts o k response has a 4xx status code
func (o *GetAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerts o k response has a 5xx status code
func (o *GetAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerts o k response a status code equal to that given
func (o *GetAlertsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alerts o k response
func (o *GetAlertsOK) Code() int {
	return 200
}

func (o *GetAlertsOK) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertsOK) String() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertsOK) GetPayload() []*models.AlertListItemDTO {
	return o.Payload
}

func (o *GetAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsUnauthorized creates a GetAlertsUnauthorized with default headers values
func NewGetAlertsUnauthorized() *GetAlertsUnauthorized {
	return &GetAlertsUnauthorized{}
}

/*
GetAlertsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAlertsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alerts unauthorized response has a 2xx status code
func (o *GetAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerts unauthorized response has a 3xx status code
func (o *GetAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerts unauthorized response has a 4xx status code
func (o *GetAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerts unauthorized response has a 5xx status code
func (o *GetAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerts unauthorized response a status code equal to that given
func (o *GetAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get alerts unauthorized response
func (o *GetAlertsUnauthorized) Code() int {
	return 401
}

func (o *GetAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertsUnauthorized) String() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsInternalServerError creates a GetAlertsInternalServerError with default headers values
func NewGetAlertsInternalServerError() *GetAlertsInternalServerError {
	return &GetAlertsInternalServerError{}
}

/*
GetAlertsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAlertsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alerts internal server error response has a 2xx status code
func (o *GetAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerts internal server error response has a 3xx status code
func (o *GetAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerts internal server error response has a 4xx status code
func (o *GetAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerts internal server error response has a 5xx status code
func (o *GetAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerts internal server error response a status code equal to that given
func (o *GetAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get alerts internal server error response
func (o *GetAlertsInternalServerError) Code() int {
	return 500
}

func (o *GetAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertsInternalServerError) String() string {
	return fmt.Sprintf("[GET /alerts][%d] getAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
