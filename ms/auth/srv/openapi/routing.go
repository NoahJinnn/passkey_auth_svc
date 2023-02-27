package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/sebest/xff"
)

func bindOAIHandlers(api *op.HqServiceAPI, srv *httpServer) {
	api.HealthCheckHandler = op.HealthCheckHandlerFunc(srv.HealthCheck)
	// Plaid API
	api.LinkTokenCreateHandler = op.LinkTokenCreateHandlerFunc(srv.LinkTokenCreate)
	api.GetAuthAccountHandler = op.GetAuthAccountHandlerFunc(srv.GetAuthAccount)
	api.GetAccessTokenHandler = op.GetAccessTokenHandlerFunc(srv.GetAccessToken)
	api.GetTransactionsHandler = op.GetTransactionsHandlerFunc(srv.GetTransactions)
	api.GetIdentityHandler = op.GetIdentityHandlerFunc(srv.GetIdentity)
	api.GetBalanceHandler = op.GetBalanceHandlerFunc(srv.GetBalance)
	api.GetAccountsHandler = op.GetAccountsHandlerFunc(srv.GetAccounts)
	// TODO: Only for testing, remove this route on production
	api.GetInfoHandler = op.GetInfoHandlerFunc(srv.GetInfo)
	api.GetSandboxAccessTokenHandler = op.GetSandboxAccessTokenHandlerFunc(srv.GetSandboxAccessToken)

	// User API
	api.GetUsersHandler = op.GetUsersHandlerFunc(srv.GetUsers)
	api.GetUserByIDHandler = op.GetUserByIDHandlerFunc(srv.GetUserById)
	api.CreateUserHandler = op.CreateUserHandlerFunc(srv.CreateUser)
	api.UpdateUserHandler = op.UpdateUserHandlerFunc(srv.UpdateUser)
}

func bindMiddlewares(api *op.HqServiceAPI, server *restapi.Server, basePath string) {
	// The middleware executes before anything.
	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(basePath)
		accesslog := makeAccessLog(basePath)
		return noCache(
			xffmw.Handler(
				logger(
					recovery(
						accesslog( //FIXME: middleware log error
							middleware.Spec(basePath, restapi.FlatSwaggerJSON, cors(handler)),
						),
					),
				),
			),
		)
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))
}
