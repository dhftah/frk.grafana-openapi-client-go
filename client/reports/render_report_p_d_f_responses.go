// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// RenderReportPDFReader is a Reader for the RenderReportPDF structure.
type RenderReportPDFReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderReportPDFReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderReportPDFOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenderReportPDFBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenderReportPDFUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenderReportPDFInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reports/render/pdf/{dashboardID}] renderReportPDF", response, response.Code())
	}
}

// NewRenderReportPDFOK creates a RenderReportPDFOK with default headers values
func NewRenderReportPDFOK() *RenderReportPDFOK {
	return &RenderReportPDFOK{}
}

/*
RenderReportPDFOK describes a response with status code 200, with default header values.

(empty)
*/
type RenderReportPDFOK struct {
	Payload []uint8
}

// IsSuccess returns true when this render report p d f o k response has a 2xx status code
func (o *RenderReportPDFOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this render report p d f o k response has a 3xx status code
func (o *RenderReportPDFOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report p d f o k response has a 4xx status code
func (o *RenderReportPDFOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this render report p d f o k response has a 5xx status code
func (o *RenderReportPDFOK) IsServerError() bool {
	return false
}

// IsCode returns true when this render report p d f o k response a status code equal to that given
func (o *RenderReportPDFOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the render report p d f o k response
func (o *RenderReportPDFOK) Code() int {
	return 200
}

func (o *RenderReportPDFOK) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFOK  %+v", 200, o.Payload)
}

func (o *RenderReportPDFOK) String() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFOK  %+v", 200, o.Payload)
}

func (o *RenderReportPDFOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *RenderReportPDFOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFBadRequest creates a RenderReportPDFBadRequest with default headers values
func NewRenderReportPDFBadRequest() *RenderReportPDFBadRequest {
	return &RenderReportPDFBadRequest{}
}

/*
RenderReportPDFBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RenderReportPDFBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report p d f bad request response has a 2xx status code
func (o *RenderReportPDFBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report p d f bad request response has a 3xx status code
func (o *RenderReportPDFBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report p d f bad request response has a 4xx status code
func (o *RenderReportPDFBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this render report p d f bad request response has a 5xx status code
func (o *RenderReportPDFBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this render report p d f bad request response a status code equal to that given
func (o *RenderReportPDFBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the render report p d f bad request response
func (o *RenderReportPDFBadRequest) Code() int {
	return 400
}

func (o *RenderReportPDFBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFBadRequest  %+v", 400, o.Payload)
}

func (o *RenderReportPDFBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFBadRequest  %+v", 400, o.Payload)
}

func (o *RenderReportPDFBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFUnauthorized creates a RenderReportPDFUnauthorized with default headers values
func NewRenderReportPDFUnauthorized() *RenderReportPDFUnauthorized {
	return &RenderReportPDFUnauthorized{}
}

/*
RenderReportPDFUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RenderReportPDFUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report p d f unauthorized response has a 2xx status code
func (o *RenderReportPDFUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report p d f unauthorized response has a 3xx status code
func (o *RenderReportPDFUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report p d f unauthorized response has a 4xx status code
func (o *RenderReportPDFUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this render report p d f unauthorized response has a 5xx status code
func (o *RenderReportPDFUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this render report p d f unauthorized response a status code equal to that given
func (o *RenderReportPDFUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the render report p d f unauthorized response
func (o *RenderReportPDFUnauthorized) Code() int {
	return 401
}

func (o *RenderReportPDFUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFUnauthorized  %+v", 401, o.Payload)
}

func (o *RenderReportPDFUnauthorized) String() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFUnauthorized  %+v", 401, o.Payload)
}

func (o *RenderReportPDFUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFInternalServerError creates a RenderReportPDFInternalServerError with default headers values
func NewRenderReportPDFInternalServerError() *RenderReportPDFInternalServerError {
	return &RenderReportPDFInternalServerError{}
}

/*
RenderReportPDFInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RenderReportPDFInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report p d f internal server error response has a 2xx status code
func (o *RenderReportPDFInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report p d f internal server error response has a 3xx status code
func (o *RenderReportPDFInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report p d f internal server error response has a 4xx status code
func (o *RenderReportPDFInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this render report p d f internal server error response has a 5xx status code
func (o *RenderReportPDFInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this render report p d f internal server error response a status code equal to that given
func (o *RenderReportPDFInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the render report p d f internal server error response
func (o *RenderReportPDFInternalServerError) Code() int {
	return 500
}

func (o *RenderReportPDFInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFInternalServerError  %+v", 500, o.Payload)
}

func (o *RenderReportPDFInternalServerError) String() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFInternalServerError  %+v", 500, o.Payload)
}

func (o *RenderReportPDFInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
