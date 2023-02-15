// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSandboxAccessTokenHandlerFunc turns a function with the right signature into a get sandbox access token handler
type GetSandboxAccessTokenHandlerFunc func(GetSandboxAccessTokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSandboxAccessTokenHandlerFunc) Handle(params GetSandboxAccessTokenParams) middleware.Responder {
	return fn(params)
}

// GetSandboxAccessTokenHandler interface for that can handle valid get sandbox access token params
type GetSandboxAccessTokenHandler interface {
	Handle(GetSandboxAccessTokenParams) middleware.Responder
}

// NewGetSandboxAccessToken creates a new http.Handler for the get sandbox access token operation
func NewGetSandboxAccessToken(ctx *middleware.Context, handler GetSandboxAccessTokenHandler) *GetSandboxAccessToken {
	return &GetSandboxAccessToken{Context: ctx, Handler: handler}
}

/*
	GetSandboxAccessToken swagger:route GET /sandbox_access_token/{institution_id} getSandboxAccessToken

GetSandboxAccessToken get sandbox access token API
*/
type GetSandboxAccessToken struct {
	Context *middleware.Context
	Handler GetSandboxAccessTokenHandler
}

func (o *GetSandboxAccessToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSandboxAccessTokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
