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
	"github.com/hellohq/hqservice/ms/networth/srv/http/ws"
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
func NewServer(appl app.Appl, sessionManager session.IManager, sharedCfg *sharedconfig.Shared, cfg *config.Config) (*echo.Echo, error) {
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
		"/networth",
		session.Session(sessionManager),
	)
	se := nw.Group("/se")
	seAccountInfo := handlers.NewSeAccountInfoHandler(srv)
	se.GET("/customers/:customer_id", seAccountInfo.Customer)
	se.POST("/customers", seAccountInfo.CreateCustomer)
	se.DELETE("/customers/:customer_id", seAccountInfo.DeleteCustomer)
	se.POST("/connect_session", seAccountInfo.CreateConnectSession)
	se.GET("/connections", seAccountInfo.GetConnectionByCustomerId)
	se.GET("/accounts", seAccountInfo.GetAccountByConnectionId)
	se.GET("/transactions", seAccountInfo.GetTxByConnectionIdAndAccountId)

	fv := nw.Group("/fv")
	fvAuth := handlers.NewFvAuthHandler(srv)
	fv.POST("/customer/token", fvAuth.CreateCustomerToken)
	fv.POST("/link/token", fvAuth.CreateLinkToken)
	fv.POST("/auth/token", fvAuth.ExchangeAccessToken)

	fvData := handlers.NewFvDataHandler(srv)
	fv.GET("/institutions/all", fvData.AllInstitution)
	fv.GET("/accounts/all", fvData.AllAccount)
	fv.GET("/transactions/all", fvData.AllTransaction)
	fv.GET("/income", fvData.Income)
	// TODO: Leave the below routes for FE testing purposes only
	fv.GET("/transactions", fvData.PagingTransaction)
	fv.GET("/balance_history/:accountId", fvData.GetBalanceHistoryByAccountId)

	itemTable := nw.Group("/item_tables")
	itHandler := handlers.NewItemTableHandler(srv)
	itemTable.GET("", itHandler.ListByUser)
	itemTable.POST("/item_table", itHandler.Create)
	itemTable.PUT("/item_table", itHandler.Update)
	itemTable.DELETE("/:itemTableId", itHandler.Delete)

	ws := ws.NewManager()
	e.GET("/sync", ws.Sync, session.Session(sessionManager))

	return e, nil
}
