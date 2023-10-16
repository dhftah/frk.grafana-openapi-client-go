// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UserSetUsingOrgReader is a Reader for the UserSetUsingOrg structure.
type UserSetUsingOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSetUsingOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserSetUsingOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserSetUsingOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserSetUsingOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserSetUsingOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserSetUsingOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user/using/{org_id}] userSetUsingOrg", response, response.Code())
	}
}

// NewUserSetUsingOrgOK creates a UserSetUsingOrgOK with default headers values
func NewUserSetUsingOrgOK() *UserSetUsingOrgOK {
	return &UserSetUsingOrgOK{}
}

/*
UserSetUsingOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UserSetUsingOrgOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this user set using org o k response has a 2xx status code
func (o *UserSetUsingOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user set using org o k response has a 3xx status code
func (o *UserSetUsingOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user set using org o k response has a 4xx status code
func (o *UserSetUsingOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user set using org o k response has a 5xx status code
func (o *UserSetUsingOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user set using org o k response a status code equal to that given
func (o *UserSetUsingOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user set using org o k response
func (o *UserSetUsingOrgOK) Code() int {
	return 200
}

func (o *UserSetUsingOrgOK) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgOK  %+v", 200, o.Payload)
}

func (o *UserSetUsingOrgOK) String() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgOK  %+v", 200, o.Payload)
}

func (o *UserSetUsingOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgBadRequest creates a UserSetUsingOrgBadRequest with default headers values
func NewUserSetUsingOrgBadRequest() *UserSetUsingOrgBadRequest {
	return &UserSetUsingOrgBadRequest{}
}

/*
UserSetUsingOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UserSetUsingOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this user set using org bad request response has a 2xx status code
func (o *UserSetUsingOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user set using org bad request response has a 3xx status code
func (o *UserSetUsingOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user set using org bad request response has a 4xx status code
func (o *UserSetUsingOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user set using org bad request response has a 5xx status code
func (o *UserSetUsingOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user set using org bad request response a status code equal to that given
func (o *UserSetUsingOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user set using org bad request response
func (o *UserSetUsingOrgBadRequest) Code() int {
	return 400
}

func (o *UserSetUsingOrgBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgBadRequest  %+v", 400, o.Payload)
}

func (o *UserSetUsingOrgBadRequest) String() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgBadRequest  %+v", 400, o.Payload)
}

func (o *UserSetUsingOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgUnauthorized creates a UserSetUsingOrgUnauthorized with default headers values
func NewUserSetUsingOrgUnauthorized() *UserSetUsingOrgUnauthorized {
	return &UserSetUsingOrgUnauthorized{}
}

/*
UserSetUsingOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UserSetUsingOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this user set using org unauthorized response has a 2xx status code
func (o *UserSetUsingOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user set using org unauthorized response has a 3xx status code
func (o *UserSetUsingOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user set using org unauthorized response has a 4xx status code
func (o *UserSetUsingOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user set using org unauthorized response has a 5xx status code
func (o *UserSetUsingOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user set using org unauthorized response a status code equal to that given
func (o *UserSetUsingOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user set using org unauthorized response
func (o *UserSetUsingOrgUnauthorized) Code() int {
	return 401
}

func (o *UserSetUsingOrgUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *UserSetUsingOrgUnauthorized) String() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *UserSetUsingOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgForbidden creates a UserSetUsingOrgForbidden with default headers values
func NewUserSetUsingOrgForbidden() *UserSetUsingOrgForbidden {
	return &UserSetUsingOrgForbidden{}
}

/*
UserSetUsingOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UserSetUsingOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this user set using org forbidden response has a 2xx status code
func (o *UserSetUsingOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user set using org forbidden response has a 3xx status code
func (o *UserSetUsingOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user set using org forbidden response has a 4xx status code
func (o *UserSetUsingOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user set using org forbidden response has a 5xx status code
func (o *UserSetUsingOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user set using org forbidden response a status code equal to that given
func (o *UserSetUsingOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user set using org forbidden response
func (o *UserSetUsingOrgForbidden) Code() int {
	return 403
}

func (o *UserSetUsingOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgForbidden  %+v", 403, o.Payload)
}

func (o *UserSetUsingOrgForbidden) String() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgForbidden  %+v", 403, o.Payload)
}

func (o *UserSetUsingOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgInternalServerError creates a UserSetUsingOrgInternalServerError with default headers values
func NewUserSetUsingOrgInternalServerError() *UserSetUsingOrgInternalServerError {
	return &UserSetUsingOrgInternalServerError{}
}

/*
UserSetUsingOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UserSetUsingOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this user set using org internal server error response has a 2xx status code
func (o *UserSetUsingOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user set using org internal server error response has a 3xx status code
func (o *UserSetUsingOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user set using org internal server error response has a 4xx status code
func (o *UserSetUsingOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user set using org internal server error response has a 5xx status code
func (o *UserSetUsingOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user set using org internal server error response a status code equal to that given
func (o *UserSetUsingOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user set using org internal server error response
func (o *UserSetUsingOrgInternalServerError) Code() int {
	return 500
}

func (o *UserSetUsingOrgInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *UserSetUsingOrgInternalServerError) String() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *UserSetUsingOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
