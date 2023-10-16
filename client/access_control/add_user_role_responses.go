// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AddUserRoleReader is a Reader for the AddUserRole structure.
type AddUserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAddUserRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddUserRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUserRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access-control/users/{userId}/roles] addUserRole", response, response.Code())
	}
}

// NewAddUserRoleOK creates a AddUserRoleOK with default headers values
func NewAddUserRoleOK() *AddUserRoleOK {
	return &AddUserRoleOK{}
}

/*
AddUserRoleOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddUserRoleOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this add user role o k response has a 2xx status code
func (o *AddUserRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user role o k response has a 3xx status code
func (o *AddUserRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user role o k response has a 4xx status code
func (o *AddUserRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user role o k response has a 5xx status code
func (o *AddUserRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user role o k response a status code equal to that given
func (o *AddUserRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user role o k response
func (o *AddUserRoleOK) Code() int {
	return 200
}

func (o *AddUserRoleOK) Error() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleOK  %+v", 200, o.Payload)
}

func (o *AddUserRoleOK) String() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleOK  %+v", 200, o.Payload)
}

func (o *AddUserRoleOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddUserRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRoleForbidden creates a AddUserRoleForbidden with default headers values
func NewAddUserRoleForbidden() *AddUserRoleForbidden {
	return &AddUserRoleForbidden{}
}

/*
AddUserRoleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddUserRoleForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add user role forbidden response has a 2xx status code
func (o *AddUserRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user role forbidden response has a 3xx status code
func (o *AddUserRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user role forbidden response has a 4xx status code
func (o *AddUserRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user role forbidden response has a 5xx status code
func (o *AddUserRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add user role forbidden response a status code equal to that given
func (o *AddUserRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add user role forbidden response
func (o *AddUserRoleForbidden) Code() int {
	return 403
}

func (o *AddUserRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleForbidden  %+v", 403, o.Payload)
}

func (o *AddUserRoleForbidden) String() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleForbidden  %+v", 403, o.Payload)
}

func (o *AddUserRoleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddUserRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRoleNotFound creates a AddUserRoleNotFound with default headers values
func NewAddUserRoleNotFound() *AddUserRoleNotFound {
	return &AddUserRoleNotFound{}
}

/*
AddUserRoleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AddUserRoleNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add user role not found response has a 2xx status code
func (o *AddUserRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user role not found response has a 3xx status code
func (o *AddUserRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user role not found response has a 4xx status code
func (o *AddUserRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user role not found response has a 5xx status code
func (o *AddUserRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add user role not found response a status code equal to that given
func (o *AddUserRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add user role not found response
func (o *AddUserRoleNotFound) Code() int {
	return 404
}

func (o *AddUserRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleNotFound  %+v", 404, o.Payload)
}

func (o *AddUserRoleNotFound) String() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleNotFound  %+v", 404, o.Payload)
}

func (o *AddUserRoleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddUserRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserRoleInternalServerError creates a AddUserRoleInternalServerError with default headers values
func NewAddUserRoleInternalServerError() *AddUserRoleInternalServerError {
	return &AddUserRoleInternalServerError{}
}

/*
AddUserRoleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddUserRoleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add user role internal server error response has a 2xx status code
func (o *AddUserRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user role internal server error response has a 3xx status code
func (o *AddUserRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user role internal server error response has a 4xx status code
func (o *AddUserRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user role internal server error response has a 5xx status code
func (o *AddUserRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add user role internal server error response a status code equal to that given
func (o *AddUserRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add user role internal server error response
func (o *AddUserRoleInternalServerError) Code() int {
	return 500
}

func (o *AddUserRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserRoleInternalServerError) String() string {
	return fmt.Sprintf("[POST /access-control/users/{userId}/roles][%d] addUserRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserRoleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddUserRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
