// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetIdentityParams creates a new GetIdentityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityParams() *GetIdentityParams {
	return &GetIdentityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityParamsWithTimeout creates a new GetIdentityParams object
// with the ability to set a timeout on a request.
func NewGetIdentityParamsWithTimeout(timeout time.Duration) *GetIdentityParams {
	return &GetIdentityParams{
		timeout: timeout,
	}
}

// NewGetIdentityParamsWithContext creates a new GetIdentityParams object
// with the ability to set a context for a request.
func NewGetIdentityParamsWithContext(ctx context.Context) *GetIdentityParams {
	return &GetIdentityParams{
		Context: ctx,
	}
}

// NewGetIdentityParamsWithHTTPClient creates a new GetIdentityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityParamsWithHTTPClient(client *http.Client) *GetIdentityParams {
	return &GetIdentityParams{
		HTTPClient: client,
	}
}

/*
GetIdentityParams contains all the parameters to send to the API endpoint

	for the get identity operation.

	Typically these are written to a http.Request.
*/
type GetIdentityParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityParams) WithDefaults() *GetIdentityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) WithTimeout(timeout time.Duration) *GetIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity params
func (o *GetIdentityParams) WithContext(ctx context.Context) *GetIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity params
func (o *GetIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) WithHTTPClient(client *http.Client) *GetIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
