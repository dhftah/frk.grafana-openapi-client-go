// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// PostRenewLicenseTokenReader is a Reader for the PostRenewLicenseToken structure.
type PostRenewLicenseTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRenewLicenseTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRenewLicenseTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRenewLicenseTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRenewLicenseTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /licensing/token/renew] postRenewLicenseToken", response, response.Code())
	}
}

// NewPostRenewLicenseTokenOK creates a PostRenewLicenseTokenOK with default headers values
func NewPostRenewLicenseTokenOK() *PostRenewLicenseTokenOK {
	return &PostRenewLicenseTokenOK{}
}

/*
PostRenewLicenseTokenOK describes a response with status code 200, with default header values.

(empty)
*/
type PostRenewLicenseTokenOK struct {
}

// IsSuccess returns true when this post renew license token o k response has a 2xx status code
func (o *PostRenewLicenseTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post renew license token o k response has a 3xx status code
func (o *PostRenewLicenseTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post renew license token o k response has a 4xx status code
func (o *PostRenewLicenseTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post renew license token o k response has a 5xx status code
func (o *PostRenewLicenseTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post renew license token o k response a status code equal to that given
func (o *PostRenewLicenseTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post renew license token o k response
func (o *PostRenewLicenseTokenOK) Code() int {
	return 200
}

func (o *PostRenewLicenseTokenOK) Error() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenOK ", 200)
}

func (o *PostRenewLicenseTokenOK) String() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenOK ", 200)
}

func (o *PostRenewLicenseTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRenewLicenseTokenUnauthorized creates a PostRenewLicenseTokenUnauthorized with default headers values
func NewPostRenewLicenseTokenUnauthorized() *PostRenewLicenseTokenUnauthorized {
	return &PostRenewLicenseTokenUnauthorized{}
}

/*
PostRenewLicenseTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PostRenewLicenseTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post renew license token unauthorized response has a 2xx status code
func (o *PostRenewLicenseTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post renew license token unauthorized response has a 3xx status code
func (o *PostRenewLicenseTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post renew license token unauthorized response has a 4xx status code
func (o *PostRenewLicenseTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post renew license token unauthorized response has a 5xx status code
func (o *PostRenewLicenseTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post renew license token unauthorized response a status code equal to that given
func (o *PostRenewLicenseTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post renew license token unauthorized response
func (o *PostRenewLicenseTokenUnauthorized) Code() int {
	return 401
}

func (o *PostRenewLicenseTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRenewLicenseTokenUnauthorized) String() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRenewLicenseTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostRenewLicenseTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenewLicenseTokenNotFound creates a PostRenewLicenseTokenNotFound with default headers values
func NewPostRenewLicenseTokenNotFound() *PostRenewLicenseTokenNotFound {
	return &PostRenewLicenseTokenNotFound{}
}

/*
PostRenewLicenseTokenNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type PostRenewLicenseTokenNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post renew license token not found response has a 2xx status code
func (o *PostRenewLicenseTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post renew license token not found response has a 3xx status code
func (o *PostRenewLicenseTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post renew license token not found response has a 4xx status code
func (o *PostRenewLicenseTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post renew license token not found response has a 5xx status code
func (o *PostRenewLicenseTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post renew license token not found response a status code equal to that given
func (o *PostRenewLicenseTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post renew license token not found response
func (o *PostRenewLicenseTokenNotFound) Code() int {
	return 404
}

func (o *PostRenewLicenseTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenNotFound  %+v", 404, o.Payload)
}

func (o *PostRenewLicenseTokenNotFound) String() string {
	return fmt.Sprintf("[POST /licensing/token/renew][%d] postRenewLicenseTokenNotFound  %+v", 404, o.Payload)
}

func (o *PostRenewLicenseTokenNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostRenewLicenseTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
