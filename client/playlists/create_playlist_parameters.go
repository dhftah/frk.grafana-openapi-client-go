// Code generated by go-swagger; DO NOT EDIT.

package playlists

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

// NewCreatePlaylistParams creates a new CreatePlaylistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePlaylistParams() *CreatePlaylistParams {
	return &CreatePlaylistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePlaylistParamsWithTimeout creates a new CreatePlaylistParams object
// with the ability to set a timeout on a request.
func NewCreatePlaylistParamsWithTimeout(timeout time.Duration) *CreatePlaylistParams {
	return &CreatePlaylistParams{
		timeout: timeout,
	}
}

// NewCreatePlaylistParamsWithContext creates a new CreatePlaylistParams object
// with the ability to set a context for a request.
func NewCreatePlaylistParamsWithContext(ctx context.Context) *CreatePlaylistParams {
	return &CreatePlaylistParams{
		Context: ctx,
	}
}

// NewCreatePlaylistParamsWithHTTPClient creates a new CreatePlaylistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePlaylistParamsWithHTTPClient(client *http.Client) *CreatePlaylistParams {
	return &CreatePlaylistParams{
		HTTPClient: client,
	}
}

/*
CreatePlaylistParams contains all the parameters to send to the API endpoint

	for the create playlist operation.

	Typically these are written to a http.Request.
*/
type CreatePlaylistParams struct {

	// Body.
	Body *models.CreatePlaylistCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlaylistParams) WithDefaults() *CreatePlaylistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create playlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlaylistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create playlist params
func (o *CreatePlaylistParams) WithTimeout(timeout time.Duration) *CreatePlaylistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create playlist params
func (o *CreatePlaylistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create playlist params
func (o *CreatePlaylistParams) WithContext(ctx context.Context) *CreatePlaylistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create playlist params
func (o *CreatePlaylistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create playlist params
func (o *CreatePlaylistParams) WithHTTPClient(client *http.Client) *CreatePlaylistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create playlist params
func (o *CreatePlaylistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create playlist params
func (o *CreatePlaylistParams) WithBody(body *models.CreatePlaylistCommand) *CreatePlaylistParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create playlist params
func (o *CreatePlaylistParams) SetBody(body *models.CreatePlaylistCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePlaylistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
