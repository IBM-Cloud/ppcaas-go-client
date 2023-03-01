// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// NewPcloudPvminstancesNetworksDeleteParams creates a new PcloudPvminstancesNetworksDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesNetworksDeleteParams() *PcloudPvminstancesNetworksDeleteParams {
	return &PcloudPvminstancesNetworksDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesNetworksDeleteParamsWithTimeout creates a new PcloudPvminstancesNetworksDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesNetworksDeleteParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksDeleteParams {
	return &PcloudPvminstancesNetworksDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesNetworksDeleteParamsWithContext creates a new PcloudPvminstancesNetworksDeleteParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesNetworksDeleteParamsWithContext(ctx context.Context) *PcloudPvminstancesNetworksDeleteParams {
	return &PcloudPvminstancesNetworksDeleteParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesNetworksDeleteParamsWithHTTPClient creates a new PcloudPvminstancesNetworksDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesNetworksDeleteParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksDeleteParams {
	return &PcloudPvminstancesNetworksDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesNetworksDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances networks delete operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesNetworksDeleteParams struct {

	/* Body.

	   Remove a network from PVM Instance parameters
	*/
	Body *models.PVMInstanceRemoveNetwork

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances networks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesNetworksDeleteParams) WithDefaults() *PcloudPvminstancesNetworksDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances networks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesNetworksDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithContext(ctx context.Context) *PcloudPvminstancesNetworksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithBody(body *models.PVMInstanceRemoveNetwork) *PcloudPvminstancesNetworksDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetBody(body *models.PVMInstanceRemoveNetwork) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesNetworksDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithNetworkID(networkID string) *PcloudPvminstancesNetworksDeleteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesNetworksDeleteParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances networks delete params
func (o *PcloudPvminstancesNetworksDeleteParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesNetworksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
