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

// AddTeamMemberReader is a Reader for the AddTeamMember structure.
type AddTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddTeamMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTeamMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /teams/{team_id}/members] addTeamMember", response, response.Code())
	}
}

// NewAddTeamMemberOK creates a AddTeamMemberOK with default headers values
func NewAddTeamMemberOK() *AddTeamMemberOK {
	return &AddTeamMemberOK{}
}

/*
AddTeamMemberOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddTeamMemberOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this add team member o k response has a 2xx status code
func (o *AddTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add team member o k response has a 3xx status code
func (o *AddTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team member o k response has a 4xx status code
func (o *AddTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add team member o k response has a 5xx status code
func (o *AddTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add team member o k response a status code equal to that given
func (o *AddTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add team member o k response
func (o *AddTeamMemberOK) Code() int {
	return 200
}

func (o *AddTeamMemberOK) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberOK  %+v", 200, o.Payload)
}

func (o *AddTeamMemberOK) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberOK  %+v", 200, o.Payload)
}

func (o *AddTeamMemberOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamMemberUnauthorized creates a AddTeamMemberUnauthorized with default headers values
func NewAddTeamMemberUnauthorized() *AddTeamMemberUnauthorized {
	return &AddTeamMemberUnauthorized{}
}

/*
AddTeamMemberUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddTeamMemberUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team member unauthorized response has a 2xx status code
func (o *AddTeamMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team member unauthorized response has a 3xx status code
func (o *AddTeamMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team member unauthorized response has a 4xx status code
func (o *AddTeamMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team member unauthorized response has a 5xx status code
func (o *AddTeamMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add team member unauthorized response a status code equal to that given
func (o *AddTeamMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add team member unauthorized response
func (o *AddTeamMemberUnauthorized) Code() int {
	return 401
}

func (o *AddTeamMemberUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *AddTeamMemberUnauthorized) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *AddTeamMemberUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamMemberForbidden creates a AddTeamMemberForbidden with default headers values
func NewAddTeamMemberForbidden() *AddTeamMemberForbidden {
	return &AddTeamMemberForbidden{}
}

/*
AddTeamMemberForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddTeamMemberForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team member forbidden response has a 2xx status code
func (o *AddTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team member forbidden response has a 3xx status code
func (o *AddTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team member forbidden response has a 4xx status code
func (o *AddTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team member forbidden response has a 5xx status code
func (o *AddTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add team member forbidden response a status code equal to that given
func (o *AddTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add team member forbidden response
func (o *AddTeamMemberForbidden) Code() int {
	return 403
}

func (o *AddTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *AddTeamMemberForbidden) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *AddTeamMemberForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamMemberNotFound creates a AddTeamMemberNotFound with default headers values
func NewAddTeamMemberNotFound() *AddTeamMemberNotFound {
	return &AddTeamMemberNotFound{}
}

/*
AddTeamMemberNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AddTeamMemberNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team member not found response has a 2xx status code
func (o *AddTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team member not found response has a 3xx status code
func (o *AddTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team member not found response has a 4xx status code
func (o *AddTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add team member not found response has a 5xx status code
func (o *AddTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add team member not found response a status code equal to that given
func (o *AddTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add team member not found response
func (o *AddTeamMemberNotFound) Code() int {
	return 404
}

func (o *AddTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *AddTeamMemberNotFound) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *AddTeamMemberNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamMemberInternalServerError creates a AddTeamMemberInternalServerError with default headers values
func NewAddTeamMemberInternalServerError() *AddTeamMemberInternalServerError {
	return &AddTeamMemberInternalServerError{}
}

/*
AddTeamMemberInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddTeamMemberInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add team member internal server error response has a 2xx status code
func (o *AddTeamMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add team member internal server error response has a 3xx status code
func (o *AddTeamMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add team member internal server error response has a 4xx status code
func (o *AddTeamMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add team member internal server error response has a 5xx status code
func (o *AddTeamMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add team member internal server error response a status code equal to that given
func (o *AddTeamMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add team member internal server error response
func (o *AddTeamMemberInternalServerError) Code() int {
	return 500
}

func (o *AddTeamMemberInternalServerError) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTeamMemberInternalServerError) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/members][%d] addTeamMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTeamMemberInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
