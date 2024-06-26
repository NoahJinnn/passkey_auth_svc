package http

import (
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/errorhandler"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/health"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/hqlog"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/session"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/validator"
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http/handlers"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http/ws"
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
	e.File("/.well-known/assetlinks.json", "static/assetlinks.json")
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

	changeset := handlers.NewChangesetHandler(srv)

	user := e.Group("/users")
	userHandler := handlers.NewUserHandler(srv, sessionManager)
	user.POST("", userHandler.Create)
	user.GET("/:id", userHandler.Get, session.Session(sessionManager))
	e.POST("/user", userHandler.GetUserIdByEmail)
	e.POST("/logout", userHandler.Logout, session.Session(sessionManager))

	e.GET("/firstlaunch", changeset.FirstLaunch, session.Session(sessionManager))
	e.DELETE("/changeset", changeset.Delete, session.Session(sessionManager))

	webauthnHandler := handlers.NewWebauthnHandler(srv, sessionManager)
	webauthn := e.Group("/webauthn")
	webauthnRegistration := webauthn.Group("/registration", session.Session(sessionManager))
	webauthnRegistration.POST("/initialize", webauthnHandler.InitRegistration)
	webauthnRegistration.POST("/finalize", webauthnHandler.FinishRegistration)

	webauthnLogin := webauthn.Group("/login")
	webauthnLogin.POST("/initialize", webauthnHandler.InitLogin)
	webauthnLogin.POST("/finalize", webauthnHandler.FinishLogin)

	webauthnCredentialsHandler := handlers.NewWebauthnCredentialHandler(srv)
	webauthnCredentials := webauthn.Group("/credentials", session.Session(sessionManager))
	webauthnCredentials.GET("", webauthnCredentialsHandler.ListByUser)
	webauthnCredentials.PATCH("/:id", webauthnCredentialsHandler.UpdateCredential)
	webauthnCredentials.DELETE("/:id", webauthnCredentialsHandler.DeleteCredential)

	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager)
	passcode := e.Group("/passcode")
	passcodeLogin := passcode.Group("/login")
	passcodeLogin.POST("/initialize", passcodeHandler.Init)
	passcodeLogin.POST("/finalize", passcodeHandler.Finish)

	emailHandler := handlers.NewEmailHandler(srv)
	email := e.Group("/emails", session.Session(sessionManager))
	email.GET("", emailHandler.ListByUser)
	email.POST("/:id/set_primary", emailHandler.SetPrimaryEmail)
	email.POST("", emailHandler.Create)
	email.DELETE("/:id", emailHandler.Delete)

	ws := ws.NewManager(srv)
	e.GET("/sync", ws.SyncBetweenUserDevices, session.Session(sessionManager))

	return e, nil
}
