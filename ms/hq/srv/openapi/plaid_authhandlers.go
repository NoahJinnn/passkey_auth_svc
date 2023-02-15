package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
)

func (srv *httpServer) HealthCheck(params op.HealthCheckParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	status, err := srv.app.HealthCheck(ctx)
	switch {
	default:
		return errHealthCheck(log, err, codeInternal)
	case err == nil:
		return op.NewHealthCheckOK().WithPayload(status)
	}
}

// // Sandbox testing
func (srv *httpServer) GetInfo(params op.GetInfoParams) middleware.Responder {
	inf := srv.app.Info()
	return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
		if err := producer.Produce(w, inf); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	})

}

func (srv *httpServer) GetSandboxAccessToken(params op.GetSandboxAccessTokenParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	t, err := srv.app.GetSandboxAccessToken(ctx, params.InstitutionID)

	switch {
	default:
		return errGetSandboxAccessToken(log, err, codeInternal)
	case err == nil:
		return op.NewGetSandboxAccessTokenOK().WithPayload(apiGetSandboxAccessToken(t))
	}

}

// // For OAuth flows, the process looks as follows.
// // 1. Create a link token with the redirectURI (as white listed at https://dashboard.plaid.com/team/api).
// // 2. Once the flow succeeds, Plaid Link will redirect to redirectURI with
// // additional parameters (as required by OAuth standards and Plaid).
// // 3. Re-initialize with the link token (from step 1) and the full received redirect URI
// // from step 2.

func (srv *httpServer) LinkTokenCreate(params op.LinkTokenCreateParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	// TODO: Implement link token create for dev,prod
	l, err := srv.app.LinkTokenCreate(ctx, nil)

	switch {
	default:
		return errLinkTokenCreate(log, err, codeInternal)
	case err == nil:
		return op.NewLinkTokenCreateOK().WithPayload(apiLinkTokenCreate(l))
	}
}

func (srv *httpServer) GetAccessToken(params op.GetAccessTokenParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	t, err := srv.app.GetAccessToken(ctx, params.PublicToken)

	switch {
	default:
		return errGetAccessToken(log, err, codeInternal)
	case err == nil:
		return op.NewGetAccessTokenOK().WithPayload(apiGetAccessToken(t))
	}
}
