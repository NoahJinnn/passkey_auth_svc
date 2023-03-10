package server

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/labstack/echo/v4"
	"github.com/powerman/structlog"
)

type (
	// Log is a synonym for convenience.
	Log = *structlog.Logger
	// Config contains configuration for OpenAPI server.
	Config struct {
		Addr netx.Addr
	}
	HttpServer struct {
		app app.Appl
		cfg Config
	}
	CustomResponder func(http.ResponseWriter, runtime.Producer)
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, cfg Config) (*echo.Echo, error) {
	srv := &HttpServer{
		app: appl,
		cfg: cfg,
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// log := structlog.New(structlog.KeyUnit, "swagger").SetDefaultKeyvals(structlog.KeyApp, config.ServiceName)
	// log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	// log.Info("Base path", "base", swaggerSpec.BasePath())

	bindOAIHandlers(e, srv)
	// bindMiddlewares(api, server, swaggerSpec.BasePath())

	e.Logger.Fatal(e.Start(cfg.Addr.String()))
	return e, nil
}
