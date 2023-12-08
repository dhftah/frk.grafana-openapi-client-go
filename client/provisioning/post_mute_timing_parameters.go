// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewPostMuteTimingParams creates a new PostMuteTimingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostMuteTimingParams() *PostMuteTimingParams {
	return &PostMuteTimingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostMuteTimingParamsWithTimeout creates a new PostMuteTimingParams object
// with the ability to set a timeout on a request.
func NewPostMuteTimingParamsWithTimeout(timeout time.Duration) *PostMuteTimingParams {
	return &PostMuteTimingParams{
		timeout: timeout,
	}
}

// NewPostMuteTimingParamsWithContext creates a new PostMuteTimingParams object
// with the ability to set a context for a request.
func NewPostMuteTimingParamsWithContext(ctx context.Context) *PostMuteTimingParams {
	return &PostMuteTimingParams{
		Context: ctx,
	}
}

// NewPostMuteTimingParamsWithHTTPClient creates a new PostMuteTimingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostMuteTimingParamsWithHTTPClient(client *http.Client) *PostMuteTimingParams {
	return &PostMuteTimingParams{
		HTTPClient: client,
	}
}

/*
PostMuteTimingParams contains all the parameters to send to the API endpoint

	for the post mute timing operation.

	Typically these are written to a http.Request.
*/
type PostMuteTimingParams struct {

	// Body.
	Body *models.MuteTimeInterval

	// XDisableProvenance.
	XDisableProvenance *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post mute timing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMuteTimingParams) WithDefaults() *PostMuteTimingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post mute timing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMuteTimingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post mute timing params
func (o *PostMuteTimingParams) WithTimeout(timeout time.Duration) *PostMuteTimingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post mute timing params
func (o *PostMuteTimingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post mute timing params
func (o *PostMuteTimingParams) WithContext(ctx context.Context) *PostMuteTimingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post mute timing params
func (o *PostMuteTimingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post mute timing params
func (o *PostMuteTimingParams) WithHTTPClient(client *http.Client) *PostMuteTimingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post mute timing params
func (o *PostMuteTimingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post mute timing params
func (o *PostMuteTimingParams) WithBody(body *models.MuteTimeInterval) *PostMuteTimingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post mute timing params
func (o *PostMuteTimingParams) SetBody(body *models.MuteTimeInterval) {
	o.Body = body
}

// WithXDisableProvenance adds the xDisableProvenance to the post mute timing params
func (o *PostMuteTimingParams) WithXDisableProvenance(xDisableProvenance *string) *PostMuteTimingParams {
	o.SetXDisableProvenance(xDisableProvenance)
	return o
}

// SetXDisableProvenance adds the xDisableProvenance to the post mute timing params
func (o *PostMuteTimingParams) SetXDisableProvenance(xDisableProvenance *string) {
	o.XDisableProvenance = xDisableProvenance
}

// WriteToRequest writes these params to a swagger request
func (o *PostMuteTimingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.XDisableProvenance != nil {

		// header param X-Disable-Provenance
		if err := r.SetHeaderParam("X-Disable-Provenance", *o.XDisableProvenance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
