package http

import (
	"fmt"

	"github.com/hellohq/hqservice/ms/nw_track/app"
	"github.com/hellohq/hqservice/ms/nw_track/config"
	"github.com/hellohq/hqservice/ms/nw_track/dal"
	"github.com/hellohq/hqservice/ms/nw_track/srv/http/dto"
	"github.com/hellohq/hqservice/ms/nw_track/srv/http/handlers"
	hqMiddlewares "github.com/hellohq/hqservice/ms/nw_track/srv/http/middlewares"
	"github.com/hellohq/hqservice/ms/nw_track/srv/http/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/powerman/structlog"
)

type (
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// NewServer returns Echo server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, repo dal.Repo, cfg *config.Config) (*echo.Echo, error) {
	srv := &handlers.HttpDeps{
		Appl: appl,
		Cfg:  cfg,
	}
	e := echo.New()
	e.HideBanner = true

	// TODO: Turn Debug to "false" in production
	e.HTTPErrorHandler = dto.NewHTTPErrorHandler(dto.HTTPErrorHandlerConfig{Debug: true, Logger: e.Logger})
	e.Use(middleware.RequestID())
	e.Use(hqMiddlewares.GetLoggerMiddleware())

	if cfg.Server.Cors.Enabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     cfg.Server.Cors.AllowOrigins,
			AllowMethods:     cfg.Server.Cors.AllowMethods,
			AllowHeaders:     cfg.Server.Cors.AllowHeaders,
			ExposeHeaders:    cfg.Server.Cors.ExposeHeaders,
			AllowCredentials: cfg.Server.Cors.AllowCredentials,
			MaxAge:           cfg.Server.Cors.MaxAge,
		}))
	}

	e.Validator = dto.NewCustomValidator()
	jwkManager, err := session.NewDefaultManager(cfg.Secrets.Keys, repo.GetJwkRepo())
	if err != nil {
		panic(fmt.Errorf("failed to create jwk manager: %w", err))
	}
	sessionManager, err := session.NewManager(jwkManager, cfg.Session)
	if err != nil {
		panic(fmt.Errorf("failed to create session generator: %w", err))
	}
	e.Logger.Fatal(e.Start(cfg.Server.BindAddr.String()))
	return e, nil
}
