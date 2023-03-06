// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/hellohq/hqservice/api/openapi/restapi/op/web_authn"
)


func configureFlags(api *op.HqServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *op.HqServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// op.GetAccessTokenMaxParseMemory = 32 << 20

	if api.CreateUserHandler == nil {
		api.CreateUserHandler = op.CreateUserHandlerFunc(func(params op.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation op.CreateUser has not yet been implemented")
		})
	}
	if api.GetAccessTokenHandler == nil {
		api.GetAccessTokenHandler = op.GetAccessTokenHandlerFunc(func(params op.GetAccessTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetAccessToken has not yet been implemented")
		})
	}
	if api.GetAccountsHandler == nil {
		api.GetAccountsHandler = op.GetAccountsHandlerFunc(func(params op.GetAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetAccounts has not yet been implemented")
		})
	}
	if api.GetAuthAccountHandler == nil {
		api.GetAuthAccountHandler = op.GetAuthAccountHandlerFunc(func(params op.GetAuthAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetAuthAccount has not yet been implemented")
		})
	}
	if api.GetBalanceHandler == nil {
		api.GetBalanceHandler = op.GetBalanceHandlerFunc(func(params op.GetBalanceParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetBalance has not yet been implemented")
		})
	}
	if api.GetIdentityHandler == nil {
		api.GetIdentityHandler = op.GetIdentityHandlerFunc(func(params op.GetIdentityParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetIdentity has not yet been implemented")
		})
	}
	if api.GetInfoHandler == nil {
		api.GetInfoHandler = op.GetInfoHandlerFunc(func(params op.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetInfo has not yet been implemented")
		})
	}
	if api.GetSandboxAccessTokenHandler == nil {
		api.GetSandboxAccessTokenHandler = op.GetSandboxAccessTokenHandlerFunc(func(params op.GetSandboxAccessTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetSandboxAccessToken has not yet been implemented")
		})
	}
	if api.GetTransactionsHandler == nil {
		api.GetTransactionsHandler = op.GetTransactionsHandlerFunc(func(params op.GetTransactionsParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetTransactions has not yet been implemented")
		})
	}
	if api.GetUserByIDHandler == nil {
		api.GetUserByIDHandler = op.GetUserByIDHandlerFunc(func(params op.GetUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetUserByID has not yet been implemented")
		})
	}
	if api.GetUsersHandler == nil {
		api.GetUsersHandler = op.GetUsersHandlerFunc(func(params op.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation op.GetUsers has not yet been implemented")
		})
	}
	if api.LinkTokenCreateHandler == nil {
		api.LinkTokenCreateHandler = op.LinkTokenCreateHandlerFunc(func(params op.LinkTokenCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation op.LinkTokenCreate has not yet been implemented")
		})
	}
	if api.UpdateUserHandler == nil {
		api.UpdateUserHandler = op.UpdateUserHandlerFunc(func(params op.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation op.UpdateUser has not yet been implemented")
		})
	}
	if api.HealthCheckHandler == nil {
		api.HealthCheckHandler = op.HealthCheckHandlerFunc(func(params op.HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation op.HealthCheck has not yet been implemented")
		})
	}
	if api.WebAuthnWebauthnLoginFinalHandler == nil {
		api.WebAuthnWebauthnLoginFinalHandler = web_authn.WebauthnLoginFinalHandlerFunc(func(params web_authn.WebauthnLoginFinalParams) middleware.Responder {
			return middleware.NotImplemented("operation web_authn.WebauthnLoginFinal has not yet been implemented")
		})
	}
	if api.WebauthnLoginInitHandler == nil {
		api.WebauthnLoginInitHandler = op.WebauthnLoginInitHandlerFunc(func(params op.WebauthnLoginInitParams) middleware.Responder {
			return middleware.NotImplemented("operation op.WebauthnLoginInit has not yet been implemented")
		})
	}
	if api.WebauthnRegFinalHandler == nil {
		api.WebauthnRegFinalHandler = op.WebauthnRegFinalHandlerFunc(func(params op.WebauthnRegFinalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation op.WebauthnRegFinal has not yet been implemented")
		})
	}
	if api.WebauthnRegInitHandler == nil {
		api.WebauthnRegInitHandler = op.WebauthnRegInitHandlerFunc(func(params op.WebauthnRegInitParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation op.WebauthnRegInit has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
