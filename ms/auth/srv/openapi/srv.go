// Package openapi implements OpenAPI server.
package openapi

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/hellohq/hqservice/api/openapi/restapi"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/structlog"
)

type (
	// Log is a synonym for convenience.
	Log = *structlog.Logger
	// Config contains configuration for OpenAPI server.
	Config struct {
		Addr netx.Addr
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

	log := structlog.New(structlog.KeyUnit, "swagger").SetDefaultKeyvals(structlog.KeyApp, config.ServiceName)
	log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	log.Info("Base path", "base", swaggerSpec.BasePath())
	api := op.NewHqServiceAPI(swaggerSpec)
	api.Logger = log.Printf

	bindOAIHandlers(api, srv)
	server := restapi.NewServer(api)
	server.Host = cfg.Addr.Host()
	server.Port = cfg.Addr.Port()
	bindMiddlewares(api, server, swaggerSpec.BasePath())

	return server, nil
}

func fromRequest(r *http.Request) (Ctx, Log) {
	ctx := r.Context()
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ctx = NewContextWithRemoteIP(ctx, remoteIP)
	log := structlog.FromContext(ctx, nil)
	log.SetDefaultKeyvals(def.LogUserID, "userID")

	return ctx, log
}

func (c CustomResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}
