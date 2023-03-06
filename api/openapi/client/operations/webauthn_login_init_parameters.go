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

// NewWebauthnLoginInitParams creates a new WebauthnLoginInitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebauthnLoginInitParams() *WebauthnLoginInitParams {
	return &WebauthnLoginInitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebauthnLoginInitParamsWithTimeout creates a new WebauthnLoginInitParams object
// with the ability to set a timeout on a request.
func NewWebauthnLoginInitParamsWithTimeout(timeout time.Duration) *WebauthnLoginInitParams {
	return &WebauthnLoginInitParams{
		timeout: timeout,
	}
}

// NewWebauthnLoginInitParamsWithContext creates a new WebauthnLoginInitParams object
// with the ability to set a context for a request.
func NewWebauthnLoginInitParamsWithContext(ctx context.Context) *WebauthnLoginInitParams {
	return &WebauthnLoginInitParams{
		Context: ctx,
	}
}

// NewWebauthnLoginInitParamsWithHTTPClient creates a new WebauthnLoginInitParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebauthnLoginInitParamsWithHTTPClient(client *http.Client) *WebauthnLoginInitParams {
	return &WebauthnLoginInitParams{
		HTTPClient: client,
	}
}

/*
WebauthnLoginInitParams contains all the parameters to send to the API endpoint

	for the webauthn login init operation.

	Typically these are written to a http.Request.
*/
type WebauthnLoginInitParams struct {

	/* User.

	   Fields need to update a area
	*/
	User WebauthnLoginInitBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webauthn login init params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnLoginInitParams) WithDefaults() *WebauthnLoginInitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webauthn login init params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnLoginInitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webauthn login init params
func (o *WebauthnLoginInitParams) WithTimeout(timeout time.Duration) *WebauthnLoginInitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webauthn login init params
func (o *WebauthnLoginInitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webauthn login init params
func (o *WebauthnLoginInitParams) WithContext(ctx context.Context) *WebauthnLoginInitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webauthn login init params
func (o *WebauthnLoginInitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webauthn login init params
func (o *WebauthnLoginInitParams) WithHTTPClient(client *http.Client) *WebauthnLoginInitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webauthn login init params
func (o *WebauthnLoginInitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the webauthn login init params
func (o *WebauthnLoginInitParams) WithUser(user WebauthnLoginInitBody) *WebauthnLoginInitParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the webauthn login init params
func (o *WebauthnLoginInitParams) SetUser(user WebauthnLoginInitBody) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *WebauthnLoginInitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
