package http

import (
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/internal/http/health"
	"github.com/hellohq/hqservice/internal/http/hqlog"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/validator"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/srv/http/handlers"
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
func NewServer(appl app.Appl, sessionManager *session.Manager, sharedCfg *sharedconfig.Shared, cfg *config.Config) (*echo.Echo, error) {
	srv := &handlers.HttpDeps{
		Appl:      appl,
		Cfg:       cfg,
		SharedCfg: sharedCfg,
	}
	e := echo.New()
	e.File("/.well-known/apple-app-site-association", "static/apple-app-site-association")
	e.File("/.well-known/assetlinks.jsons", "static/assetlinks.json")
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

	user := e.Group("/users")
	userHandler := handlers.NewUserHandler(srv, sessionManager)
	user.POST("", userHandler.Create)
	user.GET("/:id", userHandler.Get, session.Session(sessionManager))
	e.POST("/user", userHandler.GetUserIdByEmail)
	e.POST("/logout", userHandler.Logout, session.Session(sessionManager))

	webauthnHandler := handlers.NewWebauthnHandler(srv, sessionManager)
	webauthn := e.Group("/webauthn")
	webauthnRegistration := webauthn.Group("/registration", session.Session(sessionManager))
	webauthnRegistration.POST("/initialize", webauthnHandler.InitRegistration)
	webauthnRegistration.POST("/finalize", webauthnHandler.FinishRegistration)

	webauthnLogin := webauthn.Group("/login")
	webauthnLogin.POST("/initialize", webauthnHandler.InitLogin)
	webauthnLogin.POST("/finalize", webauthnHandler.FinishLogin)

	webauthnCredentials := webauthn.Group("/credentials", session.Session(sessionManager))
	webauthnCredentials.GET("", webauthnHandler.ListCredentials)
	webauthnCredentials.PATCH("/:id", webauthnHandler.UpdateCredential)
	webauthnCredentials.DELETE("/:id", webauthnHandler.DeleteCredential)

	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager)
	passcode := e.Group("/passcode")
	passcodeLogin := passcode.Group("/login")
	passcodeLogin.POST("/initialize", passcodeHandler.Init)
	passcodeLogin.POST("/finalize", passcodeHandler.Finish)

	emailHandler := handlers.NewEmailHandler(srv, sessionManager)
	email := e.Group("/emails", session.Session(sessionManager))
	email.GET("", emailHandler.ListByUser)
	email.POST("/:id/set_primary", emailHandler.SetPrimaryEmail)
	email.POST("", emailHandler.Create)
	email.DELETE("/:id", emailHandler.Delete)
	return e, nil
}
