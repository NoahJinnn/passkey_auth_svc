// Package openapi implements OpenAPI server.
package openapi

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/hellohq/hqservice/internal/apix"
	"github.com/hellohq/hqservice/ms/hq/app"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
)

type (
	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
	// Config contains configuration for OpenAPI server.
	Config struct {
		Addr     netx.Addr
		BasePath string
	}
	httpServer struct {
		app app.Appl
		cfg Config
	}
	CustomResponder func(http.ResponseWriter, runtime.Producer)
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, cfg Config) (*restapi.Server, error) {
	srv := &httpServer{
		app: appl,
		cfg: cfg,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load embedded swagger spec: %w", err)
	}
	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := op.NewPlaidConnectorAPI(swaggerSpec)
	api.Logger = structlog.New(structlog.KeyUnit, "swagger").SetDefaultKeyvals(structlog.KeyApp, config.ServiceName).Printf

	api.HealthCheckHandler = op.HealthCheckHandlerFunc(srv.HealthCheck)

	// TODO: Only for testing, remove this route on production
	api.GetInfoHandler = op.GetInfoHandlerFunc(srv.GetInfo)
	api.GetSandboxAccessTokenHandler = op.GetSandboxAccessTokenHandlerFunc(srv.GetSandboxAccessToken)

	// // TODO: Separate into other files
	api.LinkTokenCreateHandler = op.LinkTokenCreateHandlerFunc(srv.LinkTokenCreate)
	api.GetAuthAccountHandler = op.GetAuthAccountHandlerFunc(srv.GetAuthAccount)
	api.GetAccessTokenHandler = op.GetAccessTokenHandlerFunc(srv.GetAccessToken)
	api.GetTransactionsHandler = op.GetTransactionsHandlerFunc(srv.GetTransactions)
	api.GetIdentityHandler = op.GetIdentityHandlerFunc(srv.GetIdentity)
	api.GetBalanceHandler = op.GetBalanceHandlerFunc(srv.GetBalance)
	api.GetAccountsHandler = op.GetAccountsHandlerFunc(srv.GetAccounts)

	server := restapi.NewServer(api)
	server.Host = cfg.Addr.Host()
	server.Port = cfg.Addr.Port()

	// The middleware executes before anything.
	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath)
		accesslog := makeAccessLog(cfg.BasePath)
		return noCache(
			xffmw.Handler(
				logger(
					recovery(
						accesslog( //FIXME: middleware log error
							middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON, cors(handler)),
						),
					),
				),
			),
		)
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))

	log := structlog.New(structlog.KeyApp, config.ServiceName)
	log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	return server, nil
}

func fromRequest(r *http.Request) (Ctx, Log) {
	ctx := r.Context()
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ctx = apix.NewContextWithRemoteIP(ctx, remoteIP)
	log := structlog.FromContext(ctx, nil)

	// TODO: Add token or userId to log
	// userID := ""
	// if auth != nil {
	// 	userID = auth.UserID
	// }
	log.SetDefaultKeyvals(def.LogUserID, "userID")

	return ctx, log
}

func NewCustomResponder(r *http.Request, h http.Handler) middleware.Responder {
	return CustomResponder(func(w http.ResponseWriter, _ runtime.Producer) {
		h.ServeHTTP(w, r)
	})
}
func (c CustomResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}
