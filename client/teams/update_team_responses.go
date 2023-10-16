// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTeamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /teams/{team_id}] updateTeam", response, response.Code())
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*
UpdateTeamOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateTeamOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this update team o k response has a 2xx status code
func (o *UpdateTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team o k response has a 3xx status code
func (o *UpdateTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team o k response has a 4xx status code
func (o *UpdateTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team o k response has a 5xx status code
func (o *UpdateTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update team o k response a status code equal to that given
func (o *UpdateTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update team o k response
func (o *UpdateTeamOK) Code() int {
	return 200
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamUnauthorized creates a UpdateTeamUnauthorized with default headers values
func NewUpdateTeamUnauthorized() *UpdateTeamUnauthorized {
	return &UpdateTeamUnauthorized{}
}

/*
UpdateTeamUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateTeamUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update team unauthorized response has a 2xx status code
func (o *UpdateTeamUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team unauthorized response has a 3xx status code
func (o *UpdateTeamUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team unauthorized response has a 4xx status code
func (o *UpdateTeamUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team unauthorized response has a 5xx status code
func (o *UpdateTeamUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update team unauthorized response a status code equal to that given
func (o *UpdateTeamUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update team unauthorized response
func (o *UpdateTeamUnauthorized) Code() int {
	return 401
}

func (o *UpdateTeamUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTeamUnauthorized) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTeamUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamForbidden creates a UpdateTeamForbidden with default headers values
func NewUpdateTeamForbidden() *UpdateTeamForbidden {
	return &UpdateTeamForbidden{}
}

/*
UpdateTeamForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateTeamForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update team forbidden response has a 2xx status code
func (o *UpdateTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team forbidden response has a 3xx status code
func (o *UpdateTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team forbidden response has a 4xx status code
func (o *UpdateTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team forbidden response has a 5xx status code
func (o *UpdateTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update team forbidden response a status code equal to that given
func (o *UpdateTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update team forbidden response
func (o *UpdateTeamForbidden) Code() int {
	return 403
}

func (o *UpdateTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamForbidden) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamNotFound creates a UpdateTeamNotFound with default headers values
func NewUpdateTeamNotFound() *UpdateTeamNotFound {
	return &UpdateTeamNotFound{}
}

/*
UpdateTeamNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateTeamNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update team not found response has a 2xx status code
func (o *UpdateTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team not found response has a 3xx status code
func (o *UpdateTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team not found response has a 4xx status code
func (o *UpdateTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team not found response has a 5xx status code
func (o *UpdateTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update team not found response a status code equal to that given
func (o *UpdateTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update team not found response
func (o *UpdateTeamNotFound) Code() int {
	return 404
}

func (o *UpdateTeamNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamConflict creates a UpdateTeamConflict with default headers values
func NewUpdateTeamConflict() *UpdateTeamConflict {
	return &UpdateTeamConflict{}
}

/*
UpdateTeamConflict describes a response with status code 409, with default header values.

ConflictError
*/
type UpdateTeamConflict struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update team conflict response has a 2xx status code
func (o *UpdateTeamConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team conflict response has a 3xx status code
func (o *UpdateTeamConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team conflict response has a 4xx status code
func (o *UpdateTeamConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update team conflict response has a 5xx status code
func (o *UpdateTeamConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update team conflict response a status code equal to that given
func (o *UpdateTeamConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update team conflict response
func (o *UpdateTeamConflict) Code() int {
	return 409
}

func (o *UpdateTeamConflict) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamConflict  %+v", 409, o.Payload)
}

func (o *UpdateTeamConflict) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamConflict  %+v", 409, o.Payload)
}

func (o *UpdateTeamConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamInternalServerError creates a UpdateTeamInternalServerError with default headers values
func NewUpdateTeamInternalServerError() *UpdateTeamInternalServerError {
	return &UpdateTeamInternalServerError{}
}

/*
UpdateTeamInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateTeamInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update team internal server error response has a 2xx status code
func (o *UpdateTeamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update team internal server error response has a 3xx status code
func (o *UpdateTeamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team internal server error response has a 4xx status code
func (o *UpdateTeamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team internal server error response has a 5xx status code
func (o *UpdateTeamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update team internal server error response a status code equal to that given
func (o *UpdateTeamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update team internal server error response
func (o *UpdateTeamInternalServerError) Code() int {
	return 500
}

func (o *UpdateTeamInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTeamInternalServerError) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTeamInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
