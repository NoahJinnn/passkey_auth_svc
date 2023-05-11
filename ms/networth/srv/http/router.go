package http

import (
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/hellohq/hqservice/internal/http/sharedHandlers"
	"github.com/hellohq/hqservice/internal/http/sharedMiddlewares"
	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/networth/srv/http/handlers"
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
func NewServer(appl app.Appl, sessionManager session.Manager, sharedCfg *sharedConfig.Shared, cfg *config.Config) error {
	srv := &handlers.HttpDeps{
		Appl:      appl,
		Cfg:       cfg,
		SharedCfg: sharedCfg,
	}
	e := echo.New()
	e.HideBanner = true

	// TODO: Turn Debug to "false" in production
	e.HTTPErrorHandler = dto.NewHTTPErrorHandler(dto.HTTPErrorHandlerConfig{Debug: true, Logger: e.Logger})
	e.Use(middleware.RequestID())
	e.Use(sharedMiddlewares.GetLoggerMiddleware())

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

	e.Validator = sharedDto.NewCustomValidator()

	healthHandler := sharedHandlers.NewHealthHandler()
	e.GET("/ready", healthHandler.Ready)
	e.GET("/alive", healthHandler.Alive)

	nw := e.Group(
		"/nw",
		// sharedMiddlewares.Session(sessionManager), TODO: Enable back when finish nw service
	)
	nwHandler := handlers.NewSeHandler(srv)
	se := nw.Group("/se")
	se.POST("/customers", nwHandler.CreateCustomer)
	se.POST("/connect_session", nwHandler.CreateConnectSession)

	e.Logger.Fatal(e.Start(cfg.Server.BindAddr.String()))
	return nil
}
