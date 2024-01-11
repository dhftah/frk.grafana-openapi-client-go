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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewSetResourcePermissionsForUserParams creates a new SetResourcePermissionsForUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetResourcePermissionsForUserParams() *SetResourcePermissionsForUserParams {
	return &SetResourcePermissionsForUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetResourcePermissionsForUserParamsWithTimeout creates a new SetResourcePermissionsForUserParams object
// with the ability to set a timeout on a request.
func NewSetResourcePermissionsForUserParamsWithTimeout(timeout time.Duration) *SetResourcePermissionsForUserParams {
	return &SetResourcePermissionsForUserParams{
		timeout: timeout,
	}
}

// NewSetResourcePermissionsForUserParamsWithContext creates a new SetResourcePermissionsForUserParams object
// with the ability to set a context for a request.
func NewSetResourcePermissionsForUserParamsWithContext(ctx context.Context) *SetResourcePermissionsForUserParams {
	return &SetResourcePermissionsForUserParams{
		Context: ctx,
	}
}

// NewSetResourcePermissionsForUserParamsWithHTTPClient creates a new SetResourcePermissionsForUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetResourcePermissionsForUserParamsWithHTTPClient(client *http.Client) *SetResourcePermissionsForUserParams {
	return &SetResourcePermissionsForUserParams{
		HTTPClient: client,
	}
}

/*
SetResourcePermissionsForUserParams contains all the parameters to send to the API endpoint

	for the set resource permissions for user operation.

	Typically these are written to a http.Request.
*/
type SetResourcePermissionsForUserParams struct {

	// Body.
	Body *models.SetPermissionCommand

	// Resource.
	Resource string

	// ResourceID.
	ResourceID string

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set resource permissions for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetResourcePermissionsForUserParams) WithDefaults() *SetResourcePermissionsForUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set resource permissions for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetResourcePermissionsForUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithTimeout(timeout time.Duration) *SetResourcePermissionsForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithContext(ctx context.Context) *SetResourcePermissionsForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithHTTPClient(client *http.Client) *SetResourcePermissionsForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithBody(body *models.SetPermissionCommand) *SetResourcePermissionsForUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetBody(body *models.SetPermissionCommand) {
	o.Body = body
}

// WithResource adds the resource to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithResource(resource string) *SetResourcePermissionsForUserParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetResource(resource string) {
	o.Resource = resource
}

// WithResourceID adds the resourceID to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithResourceID(resourceID string) *SetResourcePermissionsForUserParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithUserID adds the userID to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) WithUserID(userID int64) *SetResourcePermissionsForUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set resource permissions for user params
func (o *SetResourcePermissionsForUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetResourcePermissionsForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource
	if err := r.SetPathParam("resource", o.Resource); err != nil {
		return err
	}

	// path param resourceID
	if err := r.SetPathParam("resourceID", o.ResourceID); err != nil {
		return err
	}

	// path param userID
	if err := r.SetPathParam("userID", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}