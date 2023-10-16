// Code generated by go-swagger; DO NOT EDIT.

package datasources

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
)

// NewDeleteDataSourceByUIDParams creates a new DeleteDataSourceByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDataSourceByUIDParams() *DeleteDataSourceByUIDParams {
	return &DeleteDataSourceByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataSourceByUIDParamsWithTimeout creates a new DeleteDataSourceByUIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDataSourceByUIDParamsWithTimeout(timeout time.Duration) *DeleteDataSourceByUIDParams {
	return &DeleteDataSourceByUIDParams{
		timeout: timeout,
	}
}

// NewDeleteDataSourceByUIDParamsWithContext creates a new DeleteDataSourceByUIDParams object
// with the ability to set a context for a request.
func NewDeleteDataSourceByUIDParamsWithContext(ctx context.Context) *DeleteDataSourceByUIDParams {
	return &DeleteDataSourceByUIDParams{
		Context: ctx,
	}
}

// NewDeleteDataSourceByUIDParamsWithHTTPClient creates a new DeleteDataSourceByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDataSourceByUIDParamsWithHTTPClient(client *http.Client) *DeleteDataSourceByUIDParams {
	return &DeleteDataSourceByUIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDataSourceByUIDParams contains all the parameters to send to the API endpoint

	for the delete data source by UID operation.

	Typically these are written to a http.Request.
*/
type DeleteDataSourceByUIDParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete data source by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataSourceByUIDParams) WithDefaults() *DeleteDataSourceByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete data source by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataSourceByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) WithTimeout(timeout time.Duration) *DeleteDataSourceByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) WithContext(ctx context.Context) *DeleteDataSourceByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) WithHTTPClient(client *http.Client) *DeleteDataSourceByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) WithUID(uid string) *DeleteDataSourceByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete data source by UID params
func (o *DeleteDataSourceByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataSourceByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
