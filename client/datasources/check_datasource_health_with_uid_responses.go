// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// CheckDatasourceHealthWithUIDReader is a Reader for the CheckDatasourceHealthWithUID structure.
type CheckDatasourceHealthWithUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckDatasourceHealthWithUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckDatasourceHealthWithUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckDatasourceHealthWithUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckDatasourceHealthWithUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckDatasourceHealthWithUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckDatasourceHealthWithUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/uid/{uid}/health] checkDatasourceHealthWithUID", response, response.Code())
	}
}

// NewCheckDatasourceHealthWithUIDOK creates a CheckDatasourceHealthWithUIDOK with default headers values
func NewCheckDatasourceHealthWithUIDOK() *CheckDatasourceHealthWithUIDOK {
	return &CheckDatasourceHealthWithUIDOK{}
}

/*
CheckDatasourceHealthWithUIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type CheckDatasourceHealthWithUIDOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this check datasource health with Uid o k response has a 2xx status code
func (o *CheckDatasourceHealthWithUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this check datasource health with Uid o k response has a 3xx status code
func (o *CheckDatasourceHealthWithUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check datasource health with Uid o k response has a 4xx status code
func (o *CheckDatasourceHealthWithUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this check datasource health with Uid o k response has a 5xx status code
func (o *CheckDatasourceHealthWithUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this check datasource health with Uid o k response a status code equal to that given
func (o *CheckDatasourceHealthWithUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the check datasource health with Uid o k response
func (o *CheckDatasourceHealthWithUIDOK) Code() int {
	return 200
}

func (o *CheckDatasourceHealthWithUIDOK) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidOK  %+v", 200, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDOK) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidOK  %+v", 200, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *CheckDatasourceHealthWithUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDatasourceHealthWithUIDBadRequest creates a CheckDatasourceHealthWithUIDBadRequest with default headers values
func NewCheckDatasourceHealthWithUIDBadRequest() *CheckDatasourceHealthWithUIDBadRequest {
	return &CheckDatasourceHealthWithUIDBadRequest{}
}

/*
CheckDatasourceHealthWithUIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CheckDatasourceHealthWithUIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this check datasource health with Uid bad request response has a 2xx status code
func (o *CheckDatasourceHealthWithUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check datasource health with Uid bad request response has a 3xx status code
func (o *CheckDatasourceHealthWithUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check datasource health with Uid bad request response has a 4xx status code
func (o *CheckDatasourceHealthWithUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this check datasource health with Uid bad request response has a 5xx status code
func (o *CheckDatasourceHealthWithUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this check datasource health with Uid bad request response a status code equal to that given
func (o *CheckDatasourceHealthWithUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the check datasource health with Uid bad request response
func (o *CheckDatasourceHealthWithUIDBadRequest) Code() int {
	return 400
}

func (o *CheckDatasourceHealthWithUIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidBadRequest  %+v", 400, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDBadRequest) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidBadRequest  %+v", 400, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CheckDatasourceHealthWithUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDatasourceHealthWithUIDUnauthorized creates a CheckDatasourceHealthWithUIDUnauthorized with default headers values
func NewCheckDatasourceHealthWithUIDUnauthorized() *CheckDatasourceHealthWithUIDUnauthorized {
	return &CheckDatasourceHealthWithUIDUnauthorized{}
}

/*
CheckDatasourceHealthWithUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CheckDatasourceHealthWithUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this check datasource health with Uid unauthorized response has a 2xx status code
func (o *CheckDatasourceHealthWithUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check datasource health with Uid unauthorized response has a 3xx status code
func (o *CheckDatasourceHealthWithUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check datasource health with Uid unauthorized response has a 4xx status code
func (o *CheckDatasourceHealthWithUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this check datasource health with Uid unauthorized response has a 5xx status code
func (o *CheckDatasourceHealthWithUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this check datasource health with Uid unauthorized response a status code equal to that given
func (o *CheckDatasourceHealthWithUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the check datasource health with Uid unauthorized response
func (o *CheckDatasourceHealthWithUIDUnauthorized) Code() int {
	return 401
}

func (o *CheckDatasourceHealthWithUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CheckDatasourceHealthWithUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDatasourceHealthWithUIDForbidden creates a CheckDatasourceHealthWithUIDForbidden with default headers values
func NewCheckDatasourceHealthWithUIDForbidden() *CheckDatasourceHealthWithUIDForbidden {
	return &CheckDatasourceHealthWithUIDForbidden{}
}

/*
CheckDatasourceHealthWithUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CheckDatasourceHealthWithUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this check datasource health with Uid forbidden response has a 2xx status code
func (o *CheckDatasourceHealthWithUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check datasource health with Uid forbidden response has a 3xx status code
func (o *CheckDatasourceHealthWithUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check datasource health with Uid forbidden response has a 4xx status code
func (o *CheckDatasourceHealthWithUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this check datasource health with Uid forbidden response has a 5xx status code
func (o *CheckDatasourceHealthWithUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this check datasource health with Uid forbidden response a status code equal to that given
func (o *CheckDatasourceHealthWithUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the check datasource health with Uid forbidden response
func (o *CheckDatasourceHealthWithUIDForbidden) Code() int {
	return 403
}

func (o *CheckDatasourceHealthWithUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidForbidden  %+v", 403, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDForbidden) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidForbidden  %+v", 403, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CheckDatasourceHealthWithUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDatasourceHealthWithUIDInternalServerError creates a CheckDatasourceHealthWithUIDInternalServerError with default headers values
func NewCheckDatasourceHealthWithUIDInternalServerError() *CheckDatasourceHealthWithUIDInternalServerError {
	return &CheckDatasourceHealthWithUIDInternalServerError{}
}

/*
CheckDatasourceHealthWithUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CheckDatasourceHealthWithUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this check datasource health with Uid internal server error response has a 2xx status code
func (o *CheckDatasourceHealthWithUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check datasource health with Uid internal server error response has a 3xx status code
func (o *CheckDatasourceHealthWithUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check datasource health with Uid internal server error response has a 4xx status code
func (o *CheckDatasourceHealthWithUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this check datasource health with Uid internal server error response has a 5xx status code
func (o *CheckDatasourceHealthWithUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this check datasource health with Uid internal server error response a status code equal to that given
func (o *CheckDatasourceHealthWithUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the check datasource health with Uid internal server error response
func (o *CheckDatasourceHealthWithUIDInternalServerError) Code() int {
	return 500
}

func (o *CheckDatasourceHealthWithUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/health][%d] checkDatasourceHealthWithUidInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckDatasourceHealthWithUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CheckDatasourceHealthWithUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
