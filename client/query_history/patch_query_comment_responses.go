// Code generated by go-swagger; DO NOT EDIT.

package query_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// PatchQueryCommentReader is a Reader for the PatchQueryComment structure.
type PatchQueryCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchQueryCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchQueryCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchQueryCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchQueryCommentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchQueryCommentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /query-history/{query_history_uid}] patchQueryComment", response, response.Code())
	}
}

// NewPatchQueryCommentOK creates a PatchQueryCommentOK with default headers values
func NewPatchQueryCommentOK() *PatchQueryCommentOK {
	return &PatchQueryCommentOK{}
}

/*
PatchQueryCommentOK describes a response with status code 200, with default header values.

(empty)
*/
type PatchQueryCommentOK struct {
	Payload *models.QueryHistoryResponse
}

// IsSuccess returns true when this patch query comment o k response has a 2xx status code
func (o *PatchQueryCommentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch query comment o k response has a 3xx status code
func (o *PatchQueryCommentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch query comment o k response has a 4xx status code
func (o *PatchQueryCommentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch query comment o k response has a 5xx status code
func (o *PatchQueryCommentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch query comment o k response a status code equal to that given
func (o *PatchQueryCommentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch query comment o k response
func (o *PatchQueryCommentOK) Code() int {
	return 200
}

func (o *PatchQueryCommentOK) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentOK  %+v", 200, o.Payload)
}

func (o *PatchQueryCommentOK) String() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentOK  %+v", 200, o.Payload)
}

func (o *PatchQueryCommentOK) GetPayload() *models.QueryHistoryResponse {
	return o.Payload
}

func (o *PatchQueryCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentBadRequest creates a PatchQueryCommentBadRequest with default headers values
func NewPatchQueryCommentBadRequest() *PatchQueryCommentBadRequest {
	return &PatchQueryCommentBadRequest{}
}

/*
PatchQueryCommentBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type PatchQueryCommentBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch query comment bad request response has a 2xx status code
func (o *PatchQueryCommentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch query comment bad request response has a 3xx status code
func (o *PatchQueryCommentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch query comment bad request response has a 4xx status code
func (o *PatchQueryCommentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch query comment bad request response has a 5xx status code
func (o *PatchQueryCommentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch query comment bad request response a status code equal to that given
func (o *PatchQueryCommentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch query comment bad request response
func (o *PatchQueryCommentBadRequest) Code() int {
	return 400
}

func (o *PatchQueryCommentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchQueryCommentBadRequest) String() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchQueryCommentBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentUnauthorized creates a PatchQueryCommentUnauthorized with default headers values
func NewPatchQueryCommentUnauthorized() *PatchQueryCommentUnauthorized {
	return &PatchQueryCommentUnauthorized{}
}

/*
PatchQueryCommentUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PatchQueryCommentUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch query comment unauthorized response has a 2xx status code
func (o *PatchQueryCommentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch query comment unauthorized response has a 3xx status code
func (o *PatchQueryCommentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch query comment unauthorized response has a 4xx status code
func (o *PatchQueryCommentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch query comment unauthorized response has a 5xx status code
func (o *PatchQueryCommentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch query comment unauthorized response a status code equal to that given
func (o *PatchQueryCommentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch query comment unauthorized response
func (o *PatchQueryCommentUnauthorized) Code() int {
	return 401
}

func (o *PatchQueryCommentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchQueryCommentUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchQueryCommentUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentInternalServerError creates a PatchQueryCommentInternalServerError with default headers values
func NewPatchQueryCommentInternalServerError() *PatchQueryCommentInternalServerError {
	return &PatchQueryCommentInternalServerError{}
}

/*
PatchQueryCommentInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PatchQueryCommentInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch query comment internal server error response has a 2xx status code
func (o *PatchQueryCommentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch query comment internal server error response has a 3xx status code
func (o *PatchQueryCommentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch query comment internal server error response has a 4xx status code
func (o *PatchQueryCommentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch query comment internal server error response has a 5xx status code
func (o *PatchQueryCommentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch query comment internal server error response a status code equal to that given
func (o *PatchQueryCommentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch query comment internal server error response
func (o *PatchQueryCommentInternalServerError) Code() int {
	return 500
}

func (o *PatchQueryCommentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchQueryCommentInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchQueryCommentInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
