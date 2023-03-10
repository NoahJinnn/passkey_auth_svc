package server

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/app/crypto"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/labstack/echo/v4"
	"github.com/powerman/structlog"
)

type (
	// Log is a synonym for convenience.
	Log = *structlog.Logger

	HttpServer struct {
		app app.Appl
		cfg config.Config
	}
	CustomResponder func(http.ResponseWriter, runtime.Producer)
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, repo dal.Repo, cfg config.Config) (*echo.Echo, error) {
	srv := &HttpServer{
		app: appl,
		cfg: cfg,
	}
	e := echo.New()
	e.Validator = dto.NewCustomValidator()
	jwkManager, err := crypto.NewDefaultManager(cfg.Secrets.Keys, repo.GetIJwkRepo())
	if err != nil {
		panic(fmt.Errorf("failed to create jwk manager: %w", err))
	}
	sessionManager, err := session.NewManager(jwkManager, cfg.Session)
	if err != nil {
		panic(fmt.Errorf("failed to create session generator: %w", err))
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// log := structlog.New(structlog.KeyUnit, "swagger").SetDefaultKeyvals(structlog.KeyApp, config.ServiceName)
	// log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	// log.Info("Base path", "base", swaggerSpec.BasePath())

	e.Logger.Fatal(e.Start(cfg.BindAddr.String()))
	return e, nil
}
