package http

import (
	"fmt"

	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/handlers"
	hqMiddlewares "github.com/hellohq/hqservice/ms/auth/srv/http/middlewares"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
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
	e.File("/.well-known/apple-app-site-association", "static/apple-app-site-association")
	e.File("/.well-known/assetlinks.json", "static/assetlinks.json")
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

	// TODO: Impl user handlers
	user := e.Group("/users")
	userHandler := handlers.NewUserHandler(srv, sessionManager)
	user.POST("", userHandler.Create)
	user.GET("/:id", userHandler.Get, hqMiddlewares.Session(sessionManager))
	e.POST("/logout", userHandler.Logout, hqMiddlewares.Session(sessionManager))
	// e.POST("/user", userHandler.GetUserIdByEmail)

	webauthnHandler := handlers.NewWebauthnHandler(srv, sessionManager)
	webauthn := e.Group("/webauthn")
	webauthnRegistration := webauthn.Group("/registration", hqMiddlewares.Session(sessionManager))
	webauthnRegistration.POST("/initialize", webauthnHandler.BeginRegistration)
	webauthnRegistration.POST("/finalize", webauthnHandler.FinishRegistration)

	webauthnLogin := webauthn.Group("/login")
	webauthnLogin.POST("/initialize", webauthnHandler.BeginLogin)
	webauthnLogin.POST("/finalize", webauthnHandler.FinishLogin)

	webauthnCredentials := webauthn.Group("/credentials", hqMiddlewares.Session(sessionManager))
	webauthnCredentials.GET("", webauthnHandler.ListCredentials)
	webauthnCredentials.PATCH("/:id", webauthnHandler.UpdateCredential)
	webauthnCredentials.DELETE("/:id", webauthnHandler.DeleteCredential)

	e.Logger.Fatal(e.Start(cfg.Server.BindAddr.String()))
	return e, nil
}
