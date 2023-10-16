// Code generated by go-swagger; DO NOT EDIT.

package access_control

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
)

// NewListRolesParams creates a new ListRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRolesParams() *ListRolesParams {
	return &ListRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRolesParamsWithTimeout creates a new ListRolesParams object
// with the ability to set a timeout on a request.
func NewListRolesParamsWithTimeout(timeout time.Duration) *ListRolesParams {
	return &ListRolesParams{
		timeout: timeout,
	}
}

// NewListRolesParamsWithContext creates a new ListRolesParams object
// with the ability to set a context for a request.
func NewListRolesParamsWithContext(ctx context.Context) *ListRolesParams {
	return &ListRolesParams{
		Context: ctx,
	}
}

// NewListRolesParamsWithHTTPClient creates a new ListRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRolesParamsWithHTTPClient(client *http.Client) *ListRolesParams {
	return &ListRolesParams{
		HTTPClient: client,
	}
}

/*
ListRolesParams contains all the parameters to send to the API endpoint

	for the list roles operation.

	Typically these are written to a http.Request.
*/
type ListRolesParams struct {

	// Delegatable.
	Delegatable *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRolesParams) WithDefaults() *ListRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list roles params
func (o *ListRolesParams) WithTimeout(timeout time.Duration) *ListRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list roles params
func (o *ListRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list roles params
func (o *ListRolesParams) WithContext(ctx context.Context) *ListRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list roles params
func (o *ListRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list roles params
func (o *ListRolesParams) WithHTTPClient(client *http.Client) *ListRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list roles params
func (o *ListRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatable adds the delegatable to the list roles params
func (o *ListRolesParams) WithDelegatable(delegatable *bool) *ListRolesParams {
	o.SetDelegatable(delegatable)
	return o
}

// SetDelegatable adds the delegatable to the list roles params
func (o *ListRolesParams) SetDelegatable(delegatable *bool) {
	o.Delegatable = delegatable
}

// WriteToRequest writes these params to a swagger request
func (o *ListRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Delegatable != nil {

		// query param delegatable
		var qrDelegatable bool

		if o.Delegatable != nil {
			qrDelegatable = *o.Delegatable
		}
		qDelegatable := swag.FormatBool(qrDelegatable)
		if qDelegatable != "" {

			if err := r.SetQueryParam("delegatable", qDelegatable); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
