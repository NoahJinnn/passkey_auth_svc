package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/hellohq/hqservice/ms/auth/srv/openapi/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/sebest/xff"
)

func bindOAIHandlers(e *echo.Echo, srv *httpServer) {
	// Plaid API

	// TODO: Only for testing, remove this route on production

	// User API
	// api.GetUsersHandler = op.GetUsersHandlerFunc(srv.GetUsers)
	// api.GetUserByIDHandler = op.GetUserByIDHandlerFunc(srv.GetUserById)
	// api.CreateUserHandler = op.CreateUserHandlerFunc(srv.CreateUser)

	// Webauthn API

}

func bindMiddlewares(api *op.HqServiceAPI, server *restapi.Server, basePath string) {
	// The middleware executes before anything.
	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := middlewares.MakeLogger(basePath)
		accesslog := middlewares.MakeAccessLog(basePath)
		return middlewares.NoCache(
			xffmw.Handler(
				logger(
					middlewares.Recovery(
						accesslog( //FIXME: middleware log error
							middleware.Spec(basePath, restapi.FlatSwaggerJSON, middlewares.Cors(handler)),
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
