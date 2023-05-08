package http

import (
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/hellohq/hqservice/internal/http/sharedHandlers"
	"github.com/hellohq/hqservice/internal/http/sharedMiddlewares"
	"github.com/hellohq/hqservice/internal/sharedConfig"
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
func NewServer(appl app.Appl, sessionManager session.Manager, sharedCfg *sharedConfig.Shared, cfg *config.Config) error {
	srv := &handlers.HttpDeps{
		Appl:      appl,
		Cfg:       cfg,
		SharedCfg: sharedCfg,
	}
	e := echo.New()
	e.File("/.well-known/apple-app-site-association", "static/apple-app-site-association")
	e.File("/.well-known/assetlinks.json", "static/assetlinks.json")
	e.HideBanner = true

	// TODO: Turn Debug to "false" in production
	e.HTTPErrorHandler = sharedDto.NewHTTPErrorHandler(sharedDto.HTTPErrorHandlerConfig{Debug: true, Logger: e.Logger})
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
	// jwkManager, err := session.NewDefaultManager(cfg.Secrets.Keys, repo.GetJwkRepo())
	// if err != nil {
	// 	panic(fmt.Errorf("failed to create jwk manager: %w", err))
	// }
	// err = jwkManager.InitJwk()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to create jwks: %w", err))
	// }
	// sessionManager, err := session.NewManager(jwkManager, cfg.Session)
	// if err != nil {
	// 	panic(fmt.Errorf("failed to create session generator: %w", err))
	// }

	healthHandler := sharedHandlers.NewHealthHandler()
	e.GET("/ready", healthHandler.Ready)
	e.GET("/alive", healthHandler.Alive)

	user := e.Group("/users")
	userHandler := handlers.NewUserHandler(srv, sessionManager)
	user.POST("", userHandler.Create)
	user.GET("/:id", userHandler.Get, sharedMiddlewares.Session(sessionManager))
	e.POST("/logout", userHandler.Logout, sharedMiddlewares.Session(sessionManager))

	webauthnHandler := handlers.NewWebauthnHandler(srv, sessionManager)
	webauthn := e.Group("/webauthn")
	webauthnRegistration := webauthn.Group("/registration", sharedMiddlewares.Session(sessionManager))
	webauthnRegistration.POST("/initialize", webauthnHandler.BeginRegistration)
	webauthnRegistration.POST("/finalize", webauthnHandler.FinishRegistration)

	webauthnLogin := webauthn.Group("/login")
	webauthnLogin.POST("/initialize", webauthnHandler.BeginLogin)
	webauthnLogin.POST("/finalize", webauthnHandler.FinishLogin)

	webauthnCredentials := webauthn.Group("/credentials", sharedMiddlewares.Session(sessionManager))
	webauthnCredentials.GET("", webauthnHandler.ListCredentials)
	webauthnCredentials.PATCH("/:id", webauthnHandler.UpdateCredential)
	webauthnCredentials.DELETE("/:id", webauthnHandler.DeleteCredential)

	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager)
	passcode := e.Group("/passcode")
	passcodeLogin := passcode.Group("/login")
	passcodeLogin.POST("/initialize", passcodeHandler.Init)
	passcodeLogin.POST("/finalize", passcodeHandler.Finish)

	e.Logger.Fatal(e.Start(cfg.Server.BindAddr.String()))
	return nil
}
