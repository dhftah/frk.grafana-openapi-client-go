// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AdminDeleteUserReader is a Reader for the AdminDeleteUser structure.
type AdminDeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /admin/users/{user_id}] adminDeleteUser", response, response.Code())
	}
}

// NewAdminDeleteUserOK creates a AdminDeleteUserOK with default headers values
func NewAdminDeleteUserOK() *AdminDeleteUserOK {
	return &AdminDeleteUserOK{}
}

/*
AdminDeleteUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminDeleteUserOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin delete user o k response has a 2xx status code
func (o *AdminDeleteUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin delete user o k response has a 3xx status code
func (o *AdminDeleteUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin delete user o k response has a 4xx status code
func (o *AdminDeleteUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin delete user o k response has a 5xx status code
func (o *AdminDeleteUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin delete user o k response a status code equal to that given
func (o *AdminDeleteUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin delete user o k response
func (o *AdminDeleteUserOK) Code() int {
	return 200
}

func (o *AdminDeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserOK  %+v", 200, o.Payload)
}

func (o *AdminDeleteUserOK) String() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserOK  %+v", 200, o.Payload)
}

func (o *AdminDeleteUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserUnauthorized creates a AdminDeleteUserUnauthorized with default headers values
func NewAdminDeleteUserUnauthorized() *AdminDeleteUserUnauthorized {
	return &AdminDeleteUserUnauthorized{}
}

/*
AdminDeleteUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminDeleteUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin delete user unauthorized response has a 2xx status code
func (o *AdminDeleteUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin delete user unauthorized response has a 3xx status code
func (o *AdminDeleteUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin delete user unauthorized response has a 4xx status code
func (o *AdminDeleteUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin delete user unauthorized response has a 5xx status code
func (o *AdminDeleteUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin delete user unauthorized response a status code equal to that given
func (o *AdminDeleteUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin delete user unauthorized response
func (o *AdminDeleteUserUnauthorized) Code() int {
	return 401
}

func (o *AdminDeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminDeleteUserUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminDeleteUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserForbidden creates a AdminDeleteUserForbidden with default headers values
func NewAdminDeleteUserForbidden() *AdminDeleteUserForbidden {
	return &AdminDeleteUserForbidden{}
}

/*
AdminDeleteUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminDeleteUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin delete user forbidden response has a 2xx status code
func (o *AdminDeleteUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin delete user forbidden response has a 3xx status code
func (o *AdminDeleteUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin delete user forbidden response has a 4xx status code
func (o *AdminDeleteUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin delete user forbidden response has a 5xx status code
func (o *AdminDeleteUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin delete user forbidden response a status code equal to that given
func (o *AdminDeleteUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin delete user forbidden response
func (o *AdminDeleteUserForbidden) Code() int {
	return 403
}

func (o *AdminDeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminDeleteUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminDeleteUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserNotFound creates a AdminDeleteUserNotFound with default headers values
func NewAdminDeleteUserNotFound() *AdminDeleteUserNotFound {
	return &AdminDeleteUserNotFound{}
}

/*
AdminDeleteUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminDeleteUserNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin delete user not found response has a 2xx status code
func (o *AdminDeleteUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin delete user not found response has a 3xx status code
func (o *AdminDeleteUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin delete user not found response has a 4xx status code
func (o *AdminDeleteUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin delete user not found response has a 5xx status code
func (o *AdminDeleteUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin delete user not found response a status code equal to that given
func (o *AdminDeleteUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin delete user not found response
func (o *AdminDeleteUserNotFound) Code() int {
	return 404
}

func (o *AdminDeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminDeleteUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminDeleteUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserInternalServerError creates a AdminDeleteUserInternalServerError with default headers values
func NewAdminDeleteUserInternalServerError() *AdminDeleteUserInternalServerError {
	return &AdminDeleteUserInternalServerError{}
}

/*
AdminDeleteUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminDeleteUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin delete user internal server error response has a 2xx status code
func (o *AdminDeleteUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin delete user internal server error response has a 3xx status code
func (o *AdminDeleteUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin delete user internal server error response has a 4xx status code
func (o *AdminDeleteUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin delete user internal server error response has a 5xx status code
func (o *AdminDeleteUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin delete user internal server error response a status code equal to that given
func (o *AdminDeleteUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin delete user internal server error response
func (o *AdminDeleteUserInternalServerError) Code() int {
	return 500
}

func (o *AdminDeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDeleteUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDeleteUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}