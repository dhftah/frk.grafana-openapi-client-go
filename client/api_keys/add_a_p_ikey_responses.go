// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AddAPIkeyReader is a Reader for the AddAPIkey structure.
type AddAPIkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAPIkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAPIkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddAPIkeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddAPIkeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddAPIkeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddAPIkeyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddAPIkeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/keys] addAPIkey", response, response.Code())
	}
}

// NewAddAPIkeyOK creates a AddAPIkeyOK with default headers values
func NewAddAPIkeyOK() *AddAPIkeyOK {
	return &AddAPIkeyOK{}
}

/*
AddAPIkeyOK describes a response with status code 200, with default header values.

(empty)
*/
type AddAPIkeyOK struct {
	Payload *models.NewAPIKeyResult
}

// IsSuccess returns true when this add a p ikey o k response has a 2xx status code
func (o *AddAPIkeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add a p ikey o k response has a 3xx status code
func (o *AddAPIkeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey o k response has a 4xx status code
func (o *AddAPIkeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add a p ikey o k response has a 5xx status code
func (o *AddAPIkeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add a p ikey o k response a status code equal to that given
func (o *AddAPIkeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add a p ikey o k response
func (o *AddAPIkeyOK) Code() int {
	return 200
}

func (o *AddAPIkeyOK) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyOK  %+v", 200, o.Payload)
}

func (o *AddAPIkeyOK) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyOK  %+v", 200, o.Payload)
}

func (o *AddAPIkeyOK) GetPayload() *models.NewAPIKeyResult {
	return o.Payload
}

func (o *AddAPIkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewAPIKeyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPIkeyBadRequest creates a AddAPIkeyBadRequest with default headers values
func NewAddAPIkeyBadRequest() *AddAPIkeyBadRequest {
	return &AddAPIkeyBadRequest{}
}

/*
AddAPIkeyBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AddAPIkeyBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add a p ikey bad request response has a 2xx status code
func (o *AddAPIkeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add a p ikey bad request response has a 3xx status code
func (o *AddAPIkeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey bad request response has a 4xx status code
func (o *AddAPIkeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add a p ikey bad request response has a 5xx status code
func (o *AddAPIkeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add a p ikey bad request response a status code equal to that given
func (o *AddAPIkeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add a p ikey bad request response
func (o *AddAPIkeyBadRequest) Code() int {
	return 400
}

func (o *AddAPIkeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyBadRequest  %+v", 400, o.Payload)
}

func (o *AddAPIkeyBadRequest) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyBadRequest  %+v", 400, o.Payload)
}

func (o *AddAPIkeyBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddAPIkeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPIkeyUnauthorized creates a AddAPIkeyUnauthorized with default headers values
func NewAddAPIkeyUnauthorized() *AddAPIkeyUnauthorized {
	return &AddAPIkeyUnauthorized{}
}

/*
AddAPIkeyUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddAPIkeyUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add a p ikey unauthorized response has a 2xx status code
func (o *AddAPIkeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add a p ikey unauthorized response has a 3xx status code
func (o *AddAPIkeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey unauthorized response has a 4xx status code
func (o *AddAPIkeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add a p ikey unauthorized response has a 5xx status code
func (o *AddAPIkeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add a p ikey unauthorized response a status code equal to that given
func (o *AddAPIkeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add a p ikey unauthorized response
func (o *AddAPIkeyUnauthorized) Code() int {
	return 401
}

func (o *AddAPIkeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyUnauthorized  %+v", 401, o.Payload)
}

func (o *AddAPIkeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyUnauthorized  %+v", 401, o.Payload)
}

func (o *AddAPIkeyUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddAPIkeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPIkeyForbidden creates a AddAPIkeyForbidden with default headers values
func NewAddAPIkeyForbidden() *AddAPIkeyForbidden {
	return &AddAPIkeyForbidden{}
}

/*
AddAPIkeyForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddAPIkeyForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add a p ikey forbidden response has a 2xx status code
func (o *AddAPIkeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add a p ikey forbidden response has a 3xx status code
func (o *AddAPIkeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey forbidden response has a 4xx status code
func (o *AddAPIkeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add a p ikey forbidden response has a 5xx status code
func (o *AddAPIkeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add a p ikey forbidden response a status code equal to that given
func (o *AddAPIkeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add a p ikey forbidden response
func (o *AddAPIkeyForbidden) Code() int {
	return 403
}

func (o *AddAPIkeyForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyForbidden  %+v", 403, o.Payload)
}

func (o *AddAPIkeyForbidden) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyForbidden  %+v", 403, o.Payload)
}

func (o *AddAPIkeyForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddAPIkeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPIkeyConflict creates a AddAPIkeyConflict with default headers values
func NewAddAPIkeyConflict() *AddAPIkeyConflict {
	return &AddAPIkeyConflict{}
}

/*
AddAPIkeyConflict describes a response with status code 409, with default header values.

ConflictError
*/
type AddAPIkeyConflict struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add a p ikey conflict response has a 2xx status code
func (o *AddAPIkeyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add a p ikey conflict response has a 3xx status code
func (o *AddAPIkeyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey conflict response has a 4xx status code
func (o *AddAPIkeyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add a p ikey conflict response has a 5xx status code
func (o *AddAPIkeyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add a p ikey conflict response a status code equal to that given
func (o *AddAPIkeyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add a p ikey conflict response
func (o *AddAPIkeyConflict) Code() int {
	return 409
}

func (o *AddAPIkeyConflict) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyConflict  %+v", 409, o.Payload)
}

func (o *AddAPIkeyConflict) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyConflict  %+v", 409, o.Payload)
}

func (o *AddAPIkeyConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddAPIkeyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPIkeyInternalServerError creates a AddAPIkeyInternalServerError with default headers values
func NewAddAPIkeyInternalServerError() *AddAPIkeyInternalServerError {
	return &AddAPIkeyInternalServerError{}
}

/*
AddAPIkeyInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddAPIkeyInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add a p ikey internal server error response has a 2xx status code
func (o *AddAPIkeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add a p ikey internal server error response has a 3xx status code
func (o *AddAPIkeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add a p ikey internal server error response has a 4xx status code
func (o *AddAPIkeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add a p ikey internal server error response has a 5xx status code
func (o *AddAPIkeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add a p ikey internal server error response a status code equal to that given
func (o *AddAPIkeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add a p ikey internal server error response
func (o *AddAPIkeyInternalServerError) Code() int {
	return 500
}

func (o *AddAPIkeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyInternalServerError  %+v", 500, o.Payload)
}

func (o *AddAPIkeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/keys][%d] addAPIkeyInternalServerError  %+v", 500, o.Payload)
}

func (o *AddAPIkeyInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddAPIkeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
