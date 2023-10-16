// Code generated by go-swagger; DO NOT EDIT.

package dashboard_permissions

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

// NewUpdateDashboardPermissionsByIDParams creates a new UpdateDashboardPermissionsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardPermissionsByIDParams() *UpdateDashboardPermissionsByIDParams {
	return &UpdateDashboardPermissionsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardPermissionsByIDParamsWithTimeout creates a new UpdateDashboardPermissionsByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardPermissionsByIDParamsWithTimeout(timeout time.Duration) *UpdateDashboardPermissionsByIDParams {
	return &UpdateDashboardPermissionsByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardPermissionsByIDParamsWithContext creates a new UpdateDashboardPermissionsByIDParams object
// with the ability to set a context for a request.
func NewUpdateDashboardPermissionsByIDParamsWithContext(ctx context.Context) *UpdateDashboardPermissionsByIDParams {
	return &UpdateDashboardPermissionsByIDParams{
		Context: ctx,
	}
}

// NewUpdateDashboardPermissionsByIDParamsWithHTTPClient creates a new UpdateDashboardPermissionsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardPermissionsByIDParamsWithHTTPClient(client *http.Client) *UpdateDashboardPermissionsByIDParams {
	return &UpdateDashboardPermissionsByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDashboardPermissionsByIDParams contains all the parameters to send to the API endpoint

	for the update dashboard permissions by ID operation.

	Typically these are written to a http.Request.
*/
type UpdateDashboardPermissionsByIDParams struct {

	// Body.
	Body *models.UpdateDashboardACLCommand

	// DashboardID.
	//
	// Format: int64
	DashboardID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard permissions by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardPermissionsByIDParams) WithDefaults() *UpdateDashboardPermissionsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard permissions by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardPermissionsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) WithTimeout(timeout time.Duration) *UpdateDashboardPermissionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) WithContext(ctx context.Context) *UpdateDashboardPermissionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) WithHTTPClient(client *http.Client) *UpdateDashboardPermissionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) WithBody(body *models.UpdateDashboardACLCommand) *UpdateDashboardPermissionsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) SetBody(body *models.UpdateDashboardACLCommand) {
	o.Body = body
}

// WithDashboardID adds the dashboardID to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) WithDashboardID(dashboardID int64) *UpdateDashboardPermissionsByIDParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the update dashboard permissions by ID params
func (o *UpdateDashboardPermissionsByIDParams) SetDashboardID(dashboardID int64) {
	o.DashboardID = dashboardID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardPermissionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param DashboardID
	if err := r.SetPathParam("DashboardID", swag.FormatInt64(o.DashboardID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
