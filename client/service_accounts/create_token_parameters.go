// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

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
	"github.com/go-openapi/swag"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewCreateTokenParams creates a new CreateTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTokenParams() *CreateTokenParams {
	return &CreateTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenParamsWithTimeout creates a new CreateTokenParams object
// with the ability to set a timeout on a request.
func NewCreateTokenParamsWithTimeout(timeout time.Duration) *CreateTokenParams {
	return &CreateTokenParams{
		timeout: timeout,
	}
}

// NewCreateTokenParamsWithContext creates a new CreateTokenParams object
// with the ability to set a context for a request.
func NewCreateTokenParamsWithContext(ctx context.Context) *CreateTokenParams {
	return &CreateTokenParams{
		Context: ctx,
	}
}

// NewCreateTokenParamsWithHTTPClient creates a new CreateTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTokenParamsWithHTTPClient(client *http.Client) *CreateTokenParams {
	return &CreateTokenParams{
		HTTPClient: client,
	}
}

/*
CreateTokenParams contains all the parameters to send to the API endpoint

	for the create token operation.

	Typically these are written to a http.Request.
*/
type CreateTokenParams struct {

	// Body.
	Body *models.AddServiceAccountTokenCommand

	// ServiceAccountID.
	//
	// Format: int64
	ServiceAccountID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) WithDefaults() *CreateTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create token params
func (o *CreateTokenParams) WithTimeout(timeout time.Duration) *CreateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token params
func (o *CreateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token params
func (o *CreateTokenParams) WithContext(ctx context.Context) *CreateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token params
func (o *CreateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) WithHTTPClient(client *http.Client) *CreateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create token params
func (o *CreateTokenParams) WithBody(body *models.AddServiceAccountTokenCommand) *CreateTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create token params
func (o *CreateTokenParams) SetBody(body *models.AddServiceAccountTokenCommand) {
	o.Body = body
}

// WithServiceAccountID adds the serviceAccountID to the create token params
func (o *CreateTokenParams) WithServiceAccountID(serviceAccountID int64) *CreateTokenParams {
	o.SetServiceAccountID(serviceAccountID)
	return o
}

// SetServiceAccountID adds the serviceAccountId to the create token params
func (o *CreateTokenParams) SetServiceAccountID(serviceAccountID int64) {
	o.ServiceAccountID = serviceAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param serviceAccountId
	if err := r.SetPathParam("serviceAccountId", swag.FormatInt64(o.ServiceAccountID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
