// Code generated by go-swagger; DO NOT EDIT.

package internal_transit_gateway

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

// NewInternalV1TransitgatewayGetParams creates a new InternalV1TransitgatewayGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1TransitgatewayGetParams() *InternalV1TransitgatewayGetParams {
	return &InternalV1TransitgatewayGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1TransitgatewayGetParamsWithTimeout creates a new InternalV1TransitgatewayGetParams object
// with the ability to set a timeout on a request.
func NewInternalV1TransitgatewayGetParamsWithTimeout(timeout time.Duration) *InternalV1TransitgatewayGetParams {
	return &InternalV1TransitgatewayGetParams{
		timeout: timeout,
	}
}

// NewInternalV1TransitgatewayGetParamsWithContext creates a new InternalV1TransitgatewayGetParams object
// with the ability to set a context for a request.
func NewInternalV1TransitgatewayGetParamsWithContext(ctx context.Context) *InternalV1TransitgatewayGetParams {
	return &InternalV1TransitgatewayGetParams{
		Context: ctx,
	}
}

// NewInternalV1TransitgatewayGetParamsWithHTTPClient creates a new InternalV1TransitgatewayGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1TransitgatewayGetParamsWithHTTPClient(client *http.Client) *InternalV1TransitgatewayGetParams {
	return &InternalV1TransitgatewayGetParams{
		HTTPClient: client,
	}
}

/*
InternalV1TransitgatewayGetParams contains all the parameters to send to the API endpoint

	for the internal v1 transitgateway get operation.

	Typically these are written to a http.Request.
*/
type InternalV1TransitgatewayGetParams struct {

	/* IBMUserAuthorization.

	   Authentication of the user that initiated the request from the TGW Platform
	*/
	IBMUserAuthorization string

	/* PowervsServiceCrn.

	   The Power Private Cloud Service Instance CRN (Targeting)
	*/
	PowervsServiceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 transitgateway get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1TransitgatewayGetParams) WithDefaults() *InternalV1TransitgatewayGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 transitgateway get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1TransitgatewayGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) WithTimeout(timeout time.Duration) *InternalV1TransitgatewayGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) WithContext(ctx context.Context) *InternalV1TransitgatewayGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) WithHTTPClient(client *http.Client) *InternalV1TransitgatewayGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) WithIBMUserAuthorization(iBMUserAuthorization string) *InternalV1TransitgatewayGetParams {
	o.SetIBMUserAuthorization(iBMUserAuthorization)
	return o
}

// SetIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) SetIBMUserAuthorization(iBMUserAuthorization string) {
	o.IBMUserAuthorization = iBMUserAuthorization
}

// WithPowervsServiceCrn adds the powervsServiceCrn to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) WithPowervsServiceCrn(powervsServiceCrn string) *InternalV1TransitgatewayGetParams {
	o.SetPowervsServiceCrn(powervsServiceCrn)
	return o
}

// SetPowervsServiceCrn adds the powervsServiceCrn to the internal v1 transitgateway get params
func (o *InternalV1TransitgatewayGetParams) SetPowervsServiceCrn(powervsServiceCrn string) {
	o.PowervsServiceCrn = powervsServiceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1TransitgatewayGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param IBM-UserAuthorization
	if err := r.SetHeaderParam("IBM-UserAuthorization", o.IBMUserAuthorization); err != nil {
		return err
	}

	// path param powervs_service_crn
	if err := r.SetPathParam("powervs_service_crn", o.PowervsServiceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
