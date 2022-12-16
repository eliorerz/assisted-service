// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewV2GetPresignedForClusterFilesParams creates a new V2GetPresignedForClusterFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2GetPresignedForClusterFilesParams() *V2GetPresignedForClusterFilesParams {
	return &V2GetPresignedForClusterFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetPresignedForClusterFilesParamsWithTimeout creates a new V2GetPresignedForClusterFilesParams object
// with the ability to set a timeout on a request.
func NewV2GetPresignedForClusterFilesParamsWithTimeout(timeout time.Duration) *V2GetPresignedForClusterFilesParams {
	return &V2GetPresignedForClusterFilesParams{
		timeout: timeout,
	}
}

// NewV2GetPresignedForClusterFilesParamsWithContext creates a new V2GetPresignedForClusterFilesParams object
// with the ability to set a context for a request.
func NewV2GetPresignedForClusterFilesParamsWithContext(ctx context.Context) *V2GetPresignedForClusterFilesParams {
	return &V2GetPresignedForClusterFilesParams{
		Context: ctx,
	}
}

// NewV2GetPresignedForClusterFilesParamsWithHTTPClient creates a new V2GetPresignedForClusterFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2GetPresignedForClusterFilesParamsWithHTTPClient(client *http.Client) *V2GetPresignedForClusterFilesParams {
	return &V2GetPresignedForClusterFilesParams{
		HTTPClient: client,
	}
}

/*
V2GetPresignedForClusterFilesParams contains all the parameters to send to the API endpoint

	for the v2 get presigned for cluster files operation.

	Typically these are written to a http.Request.
*/
type V2GetPresignedForClusterFilesParams struct {

	/* AdditionalName.

	   If downloading a manifest, the file name, prefaced with folder name, for example, openshift/99-openshift-xyz.yaml.
	*/
	AdditionalName *string

	/* ClusterID.

	   The cluster that owns the file that should be downloaded.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* FileName.

	   The file to be downloaded.
	*/
	FileName string

	/* HostID.

	   If downloading a file related to a host, the relevant host.

	   Format: uuid
	*/
	HostID *strfmt.UUID

	/* LogsType.

	   If downloading logs, the type of logs to download.
	*/
	LogsType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 get presigned for cluster files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetPresignedForClusterFilesParams) WithDefaults() *V2GetPresignedForClusterFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 get presigned for cluster files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetPresignedForClusterFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithTimeout(timeout time.Duration) *V2GetPresignedForClusterFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithContext(ctx context.Context) *V2GetPresignedForClusterFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithHTTPClient(client *http.Client) *V2GetPresignedForClusterFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdditionalName adds the additionalName to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithAdditionalName(additionalName *string) *V2GetPresignedForClusterFilesParams {
	o.SetAdditionalName(additionalName)
	return o
}

// SetAdditionalName adds the additionalName to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetAdditionalName(additionalName *string) {
	o.AdditionalName = additionalName
}

// WithClusterID adds the clusterID to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithClusterID(clusterID strfmt.UUID) *V2GetPresignedForClusterFilesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithFileName adds the fileName to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithFileName(fileName string) *V2GetPresignedForClusterFilesParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithHostID adds the hostID to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithHostID(hostID *strfmt.UUID) *V2GetPresignedForClusterFilesParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetHostID(hostID *strfmt.UUID) {
	o.HostID = hostID
}

// WithLogsType adds the logsType to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) WithLogsType(logsType *string) *V2GetPresignedForClusterFilesParams {
	o.SetLogsType(logsType)
	return o
}

// SetLogsType adds the logsType to the v2 get presigned for cluster files params
func (o *V2GetPresignedForClusterFilesParams) SetLogsType(logsType *string) {
	o.LogsType = logsType
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetPresignedForClusterFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdditionalName != nil {

		// query param additional_name
		var qrAdditionalName string

		if o.AdditionalName != nil {
			qrAdditionalName = *o.AdditionalName
		}
		qAdditionalName := qrAdditionalName
		if qAdditionalName != "" {

			if err := r.SetQueryParam("additional_name", qAdditionalName); err != nil {
				return err
			}
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {

		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
	}

	if o.HostID != nil {

		// query param host_id
		var qrHostID strfmt.UUID

		if o.HostID != nil {
			qrHostID = *o.HostID
		}
		qHostID := qrHostID.String()
		if qHostID != "" {

			if err := r.SetQueryParam("host_id", qHostID); err != nil {
				return err
			}
		}
	}

	if o.LogsType != nil {

		// query param logs_type
		var qrLogsType string

		if o.LogsType != nil {
			qrLogsType = *o.LogsType
		}
		qLogsType := qrLogsType
		if qLogsType != "" {

			if err := r.SetQueryParam("logs_type", qLogsType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
