// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// PostContactpointsReader is a Reader for the PostContactpoints structure.
type PostContactpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContactpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostContactpointsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContactpointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/provisioning/contact-points] PostContactpoints", response, response.Code())
	}
}

// NewPostContactpointsAccepted creates a PostContactpointsAccepted with default headers values
func NewPostContactpointsAccepted() *PostContactpointsAccepted {
	return &PostContactpointsAccepted{}
}

/*
PostContactpointsAccepted describes a response with status code 202, with default header values.

EmbeddedContactPoint
*/
type PostContactpointsAccepted struct {
	Payload *models.EmbeddedContactPoint
}

// IsSuccess returns true when this post contactpoints accepted response has a 2xx status code
func (o *PostContactpointsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post contactpoints accepted response has a 3xx status code
func (o *PostContactpointsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post contactpoints accepted response has a 4xx status code
func (o *PostContactpointsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post contactpoints accepted response has a 5xx status code
func (o *PostContactpointsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post contactpoints accepted response a status code equal to that given
func (o *PostContactpointsAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the post contactpoints accepted response
func (o *PostContactpointsAccepted) Code() int {
	return 202
}

func (o *PostContactpointsAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/provisioning/contact-points][%d] postContactpointsAccepted  %+v", 202, o.Payload)
}

func (o *PostContactpointsAccepted) String() string {
	return fmt.Sprintf("[POST /v1/provisioning/contact-points][%d] postContactpointsAccepted  %+v", 202, o.Payload)
}

func (o *PostContactpointsAccepted) GetPayload() *models.EmbeddedContactPoint {
	return o.Payload
}

func (o *PostContactpointsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmbeddedContactPoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContactpointsBadRequest creates a PostContactpointsBadRequest with default headers values
func NewPostContactpointsBadRequest() *PostContactpointsBadRequest {
	return &PostContactpointsBadRequest{}
}

/*
PostContactpointsBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type PostContactpointsBadRequest struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this post contactpoints bad request response has a 2xx status code
func (o *PostContactpointsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post contactpoints bad request response has a 3xx status code
func (o *PostContactpointsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post contactpoints bad request response has a 4xx status code
func (o *PostContactpointsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post contactpoints bad request response has a 5xx status code
func (o *PostContactpointsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post contactpoints bad request response a status code equal to that given
func (o *PostContactpointsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post contactpoints bad request response
func (o *PostContactpointsBadRequest) Code() int {
	return 400
}

func (o *PostContactpointsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/provisioning/contact-points][%d] postContactpointsBadRequest  %+v", 400, o.Payload)
}

func (o *PostContactpointsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/provisioning/contact-points][%d] postContactpointsBadRequest  %+v", 400, o.Payload)
}

func (o *PostContactpointsBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PostContactpointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}