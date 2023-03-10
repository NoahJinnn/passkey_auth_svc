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
	"github.com/hellohq/hqservice/ms/auth/srv/http/handlers"
	"github.com/hellohq/hqservice/ms/auth/srv/http/server/middlewares"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/labstack/echo/v4"
	"github.com/powerman/structlog"
)

type (
	// Log is a synonym for convenience.
	Log = *structlog.Logger

	CustomResponder func(http.ResponseWriter, runtime.Producer)
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, repo dal.Repo, cfg config.Config) (*echo.Echo, error) {
	srv := &handlers.HttpDeps{
		App: appl,
		Cfg: cfg,
	}
	e := echo.New()
	e.HideBanner = true

	// TODO: Setup CORS by config
	// if cfg.Server.Public.Cors.Enabled {
	// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 		AllowOrigins:     cfg.Server.Public.Cors.AllowOrigins,
	// 		AllowMethods:     cfg.Server.Public.Cors.AllowMethods,
	// 		AllowHeaders:     cfg.Server.Public.Cors.AllowHeaders,
	// 		ExposeHeaders:    cfg.Server.Public.Cors.ExposeHeaders,
	// 		AllowCredentials: cfg.Server.Public.Cors.AllowCredentials,
	// 		MaxAge:           cfg.Server.Public.Cors.MaxAge,
	// 	}))
	// }

	e.Validator = dto.NewCustomValidator()
	jwkManager, err := crypto.NewDefaultManager(cfg.Secrets.Keys, repo.GetIJwkRepo())
	if err != nil {
		panic(fmt.Errorf("failed to create jwk manager: %w", err))
	}
	sessionManager, err := session.NewManager(jwkManager, cfg.Session)
	if err != nil {
		panic(fmt.Errorf("failed to create session generator: %w", err))
	}

	// TODO: Impl user handlers
	// user := e.Group("/users")
	// user.POST("", userHandler.Create)
	// user.GET("/:id", userHandler.Get, hankoMiddleware.Session(sessionManager))

	// e.POST("/user", userHandler.GetUserIdByEmail)
	// e.POST("/logout", userHandler.Logout, hankoMiddleware.Session(sessionManager))

	webauthnHandler := handlers.NewWebauthnHandler(srv)
	webauthn := e.Group("/webauthn")
	webauthnRegistration := webauthn.Group("/registration", middlewares.Session(sessionManager))
	webauthnRegistration.POST("/initialize", webauthnHandler.BeginRegistration)
	// webauthnRegistration.POST("/finalize", webauthnHandler.FinishRegistration)

	// log := structlog.New(structlog.KeyUnit, "swagger").SetDefaultKeyvals(structlog.KeyApp, config.ServiceName)
	// log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	// log.Info("Base path", "base", swaggerSpec.BasePath())

	e.Logger.Fatal(e.Start(cfg.BindAddr.String()))
	return e, nil
}
