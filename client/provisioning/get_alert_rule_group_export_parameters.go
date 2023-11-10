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
	"github.com/go-openapi/swag"
)

// NewGetAlertRuleGroupExportParams creates a new GetAlertRuleGroupExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertRuleGroupExportParams() *GetAlertRuleGroupExportParams {
	return &GetAlertRuleGroupExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertRuleGroupExportParamsWithTimeout creates a new GetAlertRuleGroupExportParams object
// with the ability to set a timeout on a request.
func NewGetAlertRuleGroupExportParamsWithTimeout(timeout time.Duration) *GetAlertRuleGroupExportParams {
	return &GetAlertRuleGroupExportParams{
		timeout: timeout,
	}
}

// NewGetAlertRuleGroupExportParamsWithContext creates a new GetAlertRuleGroupExportParams object
// with the ability to set a context for a request.
func NewGetAlertRuleGroupExportParamsWithContext(ctx context.Context) *GetAlertRuleGroupExportParams {
	return &GetAlertRuleGroupExportParams{
		Context: ctx,
	}
}

// NewGetAlertRuleGroupExportParamsWithHTTPClient creates a new GetAlertRuleGroupExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertRuleGroupExportParamsWithHTTPClient(client *http.Client) *GetAlertRuleGroupExportParams {
	return &GetAlertRuleGroupExportParams{
		HTTPClient: client,
	}
}

/*
GetAlertRuleGroupExportParams contains all the parameters to send to the API endpoint

	for the get alert rule group export operation.

	Typically these are written to a http.Request.
*/
type GetAlertRuleGroupExportParams struct {

	// FolderUID.
	FolderUID string

	// Group.
	Group string

	/* Download.

	   Whether to initiate a download of the file or not.
	*/
	Download *bool

	/* Format.

	   Format of the downloaded file, either yaml or json. Accept header can also be used, but the query parameter will take precedence.

	   Default: "yaml"
	*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert rule group export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRuleGroupExportParams) WithDefaults() *GetAlertRuleGroupExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert rule group export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRuleGroupExportParams) SetDefaults() {
	var (
		downloadDefault = bool(false)

		formatDefault = string("yaml")
	)

	val := GetAlertRuleGroupExportParams{
		Download: &downloadDefault,
		Format:   &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithTimeout(timeout time.Duration) *GetAlertRuleGroupExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithContext(ctx context.Context) *GetAlertRuleGroupExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithHTTPClient(client *http.Client) *GetAlertRuleGroupExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderUID adds the folderUID to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithFolderUID(folderUID string) *GetAlertRuleGroupExportParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WithGroup adds the group to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithGroup(group string) *GetAlertRuleGroupExportParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetGroup(group string) {
	o.Group = group
}

// WithDownload adds the download to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithDownload(download *bool) *GetAlertRuleGroupExportParams {
	o.SetDownload(download)
	return o
}

// SetDownload adds the download to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetDownload(download *bool) {
	o.Download = download
}

// WithFormat adds the format to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) WithFormat(format *string) *GetAlertRuleGroupExportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get alert rule group export params
func (o *GetAlertRuleGroupExportParams) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertRuleGroupExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FolderUID
	if err := r.SetPathParam("FolderUID", o.FolderUID); err != nil {
		return err
	}

	// path param Group
	if err := r.SetPathParam("Group", o.Group); err != nil {
		return err
	}

	if o.Download != nil {

		// query param download
		var qrDownload bool

		if o.Download != nil {
			qrDownload = *o.Download
		}
		qDownload := swag.FormatBool(qrDownload)
		if qDownload != "" {

			if err := r.SetQueryParam("download", qDownload); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
