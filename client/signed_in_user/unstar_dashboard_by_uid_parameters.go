// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

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

// NewUnstarDashboardByUIDParams creates a new UnstarDashboardByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnstarDashboardByUIDParams() *UnstarDashboardByUIDParams {
	return &UnstarDashboardByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnstarDashboardByUIDParamsWithTimeout creates a new UnstarDashboardByUIDParams object
// with the ability to set a timeout on a request.
func NewUnstarDashboardByUIDParamsWithTimeout(timeout time.Duration) *UnstarDashboardByUIDParams {
	return &UnstarDashboardByUIDParams{
		timeout: timeout,
	}
}

// NewUnstarDashboardByUIDParamsWithContext creates a new UnstarDashboardByUIDParams object
// with the ability to set a context for a request.
func NewUnstarDashboardByUIDParamsWithContext(ctx context.Context) *UnstarDashboardByUIDParams {
	return &UnstarDashboardByUIDParams{
		Context: ctx,
	}
}

// NewUnstarDashboardByUIDParamsWithHTTPClient creates a new UnstarDashboardByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnstarDashboardByUIDParamsWithHTTPClient(client *http.Client) *UnstarDashboardByUIDParams {
	return &UnstarDashboardByUIDParams{
		HTTPClient: client,
	}
}

/*
UnstarDashboardByUIDParams contains all the parameters to send to the API endpoint

	for the unstar dashboard by UID operation.

	Typically these are written to a http.Request.
*/
type UnstarDashboardByUIDParams struct {

	// DashboardUID.
	DashboardUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unstar dashboard by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnstarDashboardByUIDParams) WithDefaults() *UnstarDashboardByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unstar dashboard by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnstarDashboardByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) WithTimeout(timeout time.Duration) *UnstarDashboardByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) WithContext(ctx context.Context) *UnstarDashboardByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) WithHTTPClient(client *http.Client) *UnstarDashboardByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardUID adds the dashboardUID to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) WithDashboardUID(dashboardUID string) *UnstarDashboardByUIDParams {
	o.SetDashboardUID(dashboardUID)
	return o
}

// SetDashboardUID adds the dashboardUid to the unstar dashboard by UID params
func (o *UnstarDashboardByUIDParams) SetDashboardUID(dashboardUID string) {
	o.DashboardUID = dashboardUID
}

// WriteToRequest writes these params to a swagger request
func (o *UnstarDashboardByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard_uid
	if err := r.SetPathParam("dashboard_uid", o.DashboardUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
