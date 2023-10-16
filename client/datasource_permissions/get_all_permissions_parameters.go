// Code generated by go-swagger; DO NOT EDIT.

package datasource_permissions

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

// NewGetAllPermissionsParams creates a new GetAllPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllPermissionsParams() *GetAllPermissionsParams {
	return &GetAllPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllPermissionsParamsWithTimeout creates a new GetAllPermissionsParams object
// with the ability to set a timeout on a request.
func NewGetAllPermissionsParamsWithTimeout(timeout time.Duration) *GetAllPermissionsParams {
	return &GetAllPermissionsParams{
		timeout: timeout,
	}
}

// NewGetAllPermissionsParamsWithContext creates a new GetAllPermissionsParams object
// with the ability to set a context for a request.
func NewGetAllPermissionsParamsWithContext(ctx context.Context) *GetAllPermissionsParams {
	return &GetAllPermissionsParams{
		Context: ctx,
	}
}

// NewGetAllPermissionsParamsWithHTTPClient creates a new GetAllPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllPermissionsParamsWithHTTPClient(client *http.Client) *GetAllPermissionsParams {
	return &GetAllPermissionsParams{
		HTTPClient: client,
	}
}

/*
GetAllPermissionsParams contains all the parameters to send to the API endpoint

	for the get all permissions operation.

	Typically these are written to a http.Request.
*/
type GetAllPermissionsParams struct {

	// DatasourceID.
	DatasourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllPermissionsParams) WithDefaults() *GetAllPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all permissions params
func (o *GetAllPermissionsParams) WithTimeout(timeout time.Duration) *GetAllPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all permissions params
func (o *GetAllPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all permissions params
func (o *GetAllPermissionsParams) WithContext(ctx context.Context) *GetAllPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all permissions params
func (o *GetAllPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all permissions params
func (o *GetAllPermissionsParams) WithHTTPClient(client *http.Client) *GetAllPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all permissions params
func (o *GetAllPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasourceID adds the datasourceID to the get all permissions params
func (o *GetAllPermissionsParams) WithDatasourceID(datasourceID string) *GetAllPermissionsParams {
	o.SetDatasourceID(datasourceID)
	return o
}

// SetDatasourceID adds the datasourceId to the get all permissions params
func (o *GetAllPermissionsParams) SetDatasourceID(datasourceID string) {
	o.DatasourceID = datasourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasourceId
	if err := r.SetPathParam("datasourceId", o.DatasourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}