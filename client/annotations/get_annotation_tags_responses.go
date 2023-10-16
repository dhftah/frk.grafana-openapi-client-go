// Code generated by go-swagger; DO NOT EDIT.

package annotations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetAnnotationTagsReader is a Reader for the GetAnnotationTags structure.
type GetAnnotationTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnnotationTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnnotationTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAnnotationTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnnotationTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /annotations/tags] getAnnotationTags", response, response.Code())
	}
}

// NewGetAnnotationTagsOK creates a GetAnnotationTagsOK with default headers values
func NewGetAnnotationTagsOK() *GetAnnotationTagsOK {
	return &GetAnnotationTagsOK{}
}

/*
GetAnnotationTagsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAnnotationTagsOK struct {
	Payload *models.GetAnnotationTagsResponse
}

// IsSuccess returns true when this get annotation tags o k response has a 2xx status code
func (o *GetAnnotationTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get annotation tags o k response has a 3xx status code
func (o *GetAnnotationTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation tags o k response has a 4xx status code
func (o *GetAnnotationTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get annotation tags o k response has a 5xx status code
func (o *GetAnnotationTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotation tags o k response a status code equal to that given
func (o *GetAnnotationTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get annotation tags o k response
func (o *GetAnnotationTagsOK) Code() int {
	return 200
}

func (o *GetAnnotationTagsOK) Error() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsOK  %+v", 200, o.Payload)
}

func (o *GetAnnotationTagsOK) String() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsOK  %+v", 200, o.Payload)
}

func (o *GetAnnotationTagsOK) GetPayload() *models.GetAnnotationTagsResponse {
	return o.Payload
}

func (o *GetAnnotationTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAnnotationTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnnotationTagsUnauthorized creates a GetAnnotationTagsUnauthorized with default headers values
func NewGetAnnotationTagsUnauthorized() *GetAnnotationTagsUnauthorized {
	return &GetAnnotationTagsUnauthorized{}
}

/*
GetAnnotationTagsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAnnotationTagsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get annotation tags unauthorized response has a 2xx status code
func (o *GetAnnotationTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotation tags unauthorized response has a 3xx status code
func (o *GetAnnotationTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation tags unauthorized response has a 4xx status code
func (o *GetAnnotationTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get annotation tags unauthorized response has a 5xx status code
func (o *GetAnnotationTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotation tags unauthorized response a status code equal to that given
func (o *GetAnnotationTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get annotation tags unauthorized response
func (o *GetAnnotationTagsUnauthorized) Code() int {
	return 401
}

func (o *GetAnnotationTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnnotationTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnnotationTagsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAnnotationTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnnotationTagsInternalServerError creates a GetAnnotationTagsInternalServerError with default headers values
func NewGetAnnotationTagsInternalServerError() *GetAnnotationTagsInternalServerError {
	return &GetAnnotationTagsInternalServerError{}
}

/*
GetAnnotationTagsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAnnotationTagsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get annotation tags internal server error response has a 2xx status code
func (o *GetAnnotationTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotation tags internal server error response has a 3xx status code
func (o *GetAnnotationTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation tags internal server error response has a 4xx status code
func (o *GetAnnotationTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get annotation tags internal server error response has a 5xx status code
func (o *GetAnnotationTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get annotation tags internal server error response a status code equal to that given
func (o *GetAnnotationTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get annotation tags internal server error response
func (o *GetAnnotationTagsInternalServerError) Code() int {
	return 500
}

func (o *GetAnnotationTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnnotationTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /annotations/tags][%d] getAnnotationTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnnotationTagsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAnnotationTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
