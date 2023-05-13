package session

import (
	"net/http"

	"github.com/hellohq/hqservice/internal/http/errorhandler"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Session is a convenience function to create a middleware.JWT with custom JWT verification
func Session(generator Manager) echo.MiddlewareFunc {
	c := echojwt.Config{
		ContextKey:     "session",
		TokenLookup:    "header:Authorization:Bearer,cookie:hqservice",
		ParseTokenFunc: parseToken(generator),
		ErrorHandler: func(c echo.Context, err error) error {
			return errorhandler.NewHTTPError(http.StatusUnauthorized).SetInternal(err)
		},
	}
	return echojwt.WithConfig(c)
}

type ParseTokenFunc = func(c echo.Context, auth string) (interface{}, error)

func parseToken(generator Manager) ParseTokenFunc {
	return func(c echo.Context, auth string) (interface{}, error) {
		return generator.Verify(auth)
	}
}
