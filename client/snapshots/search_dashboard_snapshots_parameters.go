// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// NewSearchDashboardSnapshotsParams creates a new SearchDashboardSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchDashboardSnapshotsParams() *SearchDashboardSnapshotsParams {
	return &SearchDashboardSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchDashboardSnapshotsParamsWithTimeout creates a new SearchDashboardSnapshotsParams object
// with the ability to set a timeout on a request.
func NewSearchDashboardSnapshotsParamsWithTimeout(timeout time.Duration) *SearchDashboardSnapshotsParams {
	return &SearchDashboardSnapshotsParams{
		timeout: timeout,
	}
}

// NewSearchDashboardSnapshotsParamsWithContext creates a new SearchDashboardSnapshotsParams object
// with the ability to set a context for a request.
func NewSearchDashboardSnapshotsParamsWithContext(ctx context.Context) *SearchDashboardSnapshotsParams {
	return &SearchDashboardSnapshotsParams{
		Context: ctx,
	}
}

// NewSearchDashboardSnapshotsParamsWithHTTPClient creates a new SearchDashboardSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchDashboardSnapshotsParamsWithHTTPClient(client *http.Client) *SearchDashboardSnapshotsParams {
	return &SearchDashboardSnapshotsParams{
		HTTPClient: client,
	}
}

/*
SearchDashboardSnapshotsParams contains all the parameters to send to the API endpoint

	for the search dashboard snapshots operation.

	Typically these are written to a http.Request.
*/
type SearchDashboardSnapshotsParams struct {

	/* Limit.

	   Limit the number of returned results

	   Format: int64
	   Default: 1000
	*/
	Limit *int64

	/* Query.

	   Search Query
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search dashboard snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDashboardSnapshotsParams) WithDefaults() *SearchDashboardSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search dashboard snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDashboardSnapshotsParams) SetDefaults() {
	var (
		limitDefault = int64(1000)
	)

	val := SearchDashboardSnapshotsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) WithTimeout(timeout time.Duration) *SearchDashboardSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) WithContext(ctx context.Context) *SearchDashboardSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) WithHTTPClient(client *http.Client) *SearchDashboardSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) WithLimit(limit *int64) *SearchDashboardSnapshotsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithQuery adds the query to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) WithQuery(query *string) *SearchDashboardSnapshotsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search dashboard snapshots params
func (o *SearchDashboardSnapshotsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SearchDashboardSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}