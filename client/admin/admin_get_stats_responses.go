// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AdminGetStatsReader is a Reader for the AdminGetStats structure.
type AdminGetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminGetStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminGetStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/stats] adminGetStats", response, response.Code())
	}
}

// NewAdminGetStatsOK creates a AdminGetStatsOK with default headers values
func NewAdminGetStatsOK() *AdminGetStatsOK {
	return &AdminGetStatsOK{}
}

/*
AdminGetStatsOK describes a response with status code 200, with default header values.

(empty)
*/
type AdminGetStatsOK struct {
	Payload *models.AdminStats
}

// IsSuccess returns true when this admin get stats o k response has a 2xx status code
func (o *AdminGetStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin get stats o k response has a 3xx status code
func (o *AdminGetStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get stats o k response has a 4xx status code
func (o *AdminGetStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin get stats o k response has a 5xx status code
func (o *AdminGetStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin get stats o k response a status code equal to that given
func (o *AdminGetStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin get stats o k response
func (o *AdminGetStatsOK) Code() int {
	return 200
}

func (o *AdminGetStatsOK) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsOK  %+v", 200, o.Payload)
}

func (o *AdminGetStatsOK) String() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsOK  %+v", 200, o.Payload)
}

func (o *AdminGetStatsOK) GetPayload() *models.AdminStats {
	return o.Payload
}

func (o *AdminGetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsUnauthorized creates a AdminGetStatsUnauthorized with default headers values
func NewAdminGetStatsUnauthorized() *AdminGetStatsUnauthorized {
	return &AdminGetStatsUnauthorized{}
}

/*
AdminGetStatsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminGetStatsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin get stats unauthorized response has a 2xx status code
func (o *AdminGetStatsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin get stats unauthorized response has a 3xx status code
func (o *AdminGetStatsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get stats unauthorized response has a 4xx status code
func (o *AdminGetStatsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin get stats unauthorized response has a 5xx status code
func (o *AdminGetStatsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin get stats unauthorized response a status code equal to that given
func (o *AdminGetStatsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin get stats unauthorized response
func (o *AdminGetStatsUnauthorized) Code() int {
	return 401
}

func (o *AdminGetStatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminGetStatsUnauthorized) String() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminGetStatsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsForbidden creates a AdminGetStatsForbidden with default headers values
func NewAdminGetStatsForbidden() *AdminGetStatsForbidden {
	return &AdminGetStatsForbidden{}
}

/*
AdminGetStatsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminGetStatsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin get stats forbidden response has a 2xx status code
func (o *AdminGetStatsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin get stats forbidden response has a 3xx status code
func (o *AdminGetStatsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get stats forbidden response has a 4xx status code
func (o *AdminGetStatsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin get stats forbidden response has a 5xx status code
func (o *AdminGetStatsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin get stats forbidden response a status code equal to that given
func (o *AdminGetStatsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin get stats forbidden response
func (o *AdminGetStatsForbidden) Code() int {
	return 403
}

func (o *AdminGetStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsForbidden  %+v", 403, o.Payload)
}

func (o *AdminGetStatsForbidden) String() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsForbidden  %+v", 403, o.Payload)
}

func (o *AdminGetStatsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsInternalServerError creates a AdminGetStatsInternalServerError with default headers values
func NewAdminGetStatsInternalServerError() *AdminGetStatsInternalServerError {
	return &AdminGetStatsInternalServerError{}
}

/*
AdminGetStatsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminGetStatsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin get stats internal server error response has a 2xx status code
func (o *AdminGetStatsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin get stats internal server error response has a 3xx status code
func (o *AdminGetStatsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get stats internal server error response has a 4xx status code
func (o *AdminGetStatsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin get stats internal server error response has a 5xx status code
func (o *AdminGetStatsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin get stats internal server error response a status code equal to that given
func (o *AdminGetStatsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin get stats internal server error response
func (o *AdminGetStatsInternalServerError) Code() int {
	return 500
}

func (o *AdminGetStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminGetStatsInternalServerError) String() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminGetStatsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
