package handlers

import (
	"context"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var ctx = context.Background()
var defaultConfig = config.Config{}

type sessionManager struct {
}

var userId = "ec4ef049-5b88-4321-a173-21b0eff06a04"

func (s sessionManager) GenerateJWT(uuid uuid.UUID) (string, error) {
	return userId, nil
}

func (s sessionManager) GenerateCookie(token string) (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hanko",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}, nil
}

func (s sessionManager) DeleteCookie() (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hanko",
		Value:    "",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}, nil
}

func (s sessionManager) Verify(token string) (jwt.Token, error) {
	return nil, nil
}
