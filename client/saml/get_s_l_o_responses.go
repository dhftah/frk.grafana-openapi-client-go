// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetSLOReader is a Reader for the GetSLO structure.
type GetSLOReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSLOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetSLOFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSLOBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSLOForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSLOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /saml/slo] getSLO", response, response.Code())
	}
}

// NewGetSLOFound creates a GetSLOFound with default headers values
func NewGetSLOFound() *GetSLOFound {
	return &GetSLOFound{}
}

/*
GetSLOFound describes a response with status code 302, with default header values.

(empty)
*/
type GetSLOFound struct {
}

// IsSuccess returns true when this get s l o found response has a 2xx status code
func (o *GetSLOFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s l o found response has a 3xx status code
func (o *GetSLOFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get s l o found response has a 4xx status code
func (o *GetSLOFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s l o found response has a 5xx status code
func (o *GetSLOFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get s l o found response a status code equal to that given
func (o *GetSLOFound) IsCode(code int) bool {
	return code == 302
}

// Code gets the status code for the get s l o found response
func (o *GetSLOFound) Code() int {
	return 302
}

func (o *GetSLOFound) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOFound ", 302)
}

func (o *GetSLOFound) String() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOFound ", 302)
}

func (o *GetSLOFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSLOBadRequest creates a GetSLOBadRequest with default headers values
func NewGetSLOBadRequest() *GetSLOBadRequest {
	return &GetSLOBadRequest{}
}

/*
GetSLOBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetSLOBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get s l o bad request response has a 2xx status code
func (o *GetSLOBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s l o bad request response has a 3xx status code
func (o *GetSLOBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s l o bad request response has a 4xx status code
func (o *GetSLOBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s l o bad request response has a 5xx status code
func (o *GetSLOBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get s l o bad request response a status code equal to that given
func (o *GetSLOBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get s l o bad request response
func (o *GetSLOBadRequest) Code() int {
	return 400
}

func (o *GetSLOBadRequest) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOBadRequest  %+v", 400, o.Payload)
}

func (o *GetSLOBadRequest) String() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOBadRequest  %+v", 400, o.Payload)
}

func (o *GetSLOBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSLOForbidden creates a GetSLOForbidden with default headers values
func NewGetSLOForbidden() *GetSLOForbidden {
	return &GetSLOForbidden{}
}

/*
GetSLOForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSLOForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get s l o forbidden response has a 2xx status code
func (o *GetSLOForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s l o forbidden response has a 3xx status code
func (o *GetSLOForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s l o forbidden response has a 4xx status code
func (o *GetSLOForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s l o forbidden response has a 5xx status code
func (o *GetSLOForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get s l o forbidden response a status code equal to that given
func (o *GetSLOForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get s l o forbidden response
func (o *GetSLOForbidden) Code() int {
	return 403
}

func (o *GetSLOForbidden) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOForbidden  %+v", 403, o.Payload)
}

func (o *GetSLOForbidden) String() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOForbidden  %+v", 403, o.Payload)
}

func (o *GetSLOForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSLOInternalServerError creates a GetSLOInternalServerError with default headers values
func NewGetSLOInternalServerError() *GetSLOInternalServerError {
	return &GetSLOInternalServerError{}
}

/*
GetSLOInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSLOInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get s l o internal server error response has a 2xx status code
func (o *GetSLOInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s l o internal server error response has a 3xx status code
func (o *GetSLOInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s l o internal server error response has a 4xx status code
func (o *GetSLOInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s l o internal server error response has a 5xx status code
func (o *GetSLOInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get s l o internal server error response a status code equal to that given
func (o *GetSLOInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get s l o internal server error response
func (o *GetSLOInternalServerError) Code() int {
	return 500
}

func (o *GetSLOInternalServerError) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSLOInternalServerError) String() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSLOInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
