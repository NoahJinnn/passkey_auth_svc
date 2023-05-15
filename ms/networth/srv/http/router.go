package http

import (
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/internal/http/health"
	"github.com/hellohq/hqservice/internal/http/hqlog"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/validator"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
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
func NewServer(appl app.Appl, sessionManager session.Manager, sharedCfg *sharedconfig.Shared, cfg *config.Config) error {
	srv := &handlers.HttpDeps{
		Appl:      appl,
		Cfg:       cfg,
		SharedCfg: sharedCfg,
	}
	e := echo.New()
	e.HideBanner = true

	// TODO: Turn Debug to "false" in production
	e.HTTPErrorHandler = errorhandler.NewHTTPErrorHandler(errorhandler.HTTPErrorHandlerConfig{Debug: true, Logger: e.Logger})
	e.Use(middleware.RequestID())
	e.Use(hqlog.GetLoggerMiddleware())

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

	e.Validator = validator.NewCustomValidator()

	healthHandler := health.NewHealthHandler()
	e.GET("/ready", healthHandler.Ready)
	e.GET("/alive", healthHandler.Alive)

	nw := e.Group(
		"/nw",
		// session.Session(sessionManager), TODO: Enable back when finish nw service
	)
	nwHandler := handlers.NewSeHandler(srv)
	se := nw.Group("/se")
	se.POST("/customers", nwHandler.CreateCustomer)
	se.POST("/connect_session", nwHandler.CreateConnectSession)
	se.GET("/connects", nwHandler.GetConnectionByCustomerId)
	e.Logger.Fatal(e.Start(cfg.Server.BindAddr.String()))
	return nil
}
