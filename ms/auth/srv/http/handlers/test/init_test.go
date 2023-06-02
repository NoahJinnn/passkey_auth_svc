package test

import (
	"context"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var (
	ctx        = context.Background()
	defaultCfg = config.Config{
		Webauthn: config.WebauthnSettings{
			RelyingParty: config.RelyingParty{
				Id:          "localhost",
				DisplayName: "Test Relying Party",
				Icon:        "",
				Origins:     []string{"http://localhost:8080"},
			},
			Timeout: 60000,
		},
	}
)

var sharedCfg = sharedconfig.Shared{
	Session: sharedconfig.Session{
		Lifespan: "1h",
		Cookie: sharedconfig.Cookie{
			HttpOnly: true,
			SameSite: "strict",
			Secure:   true,
		},
		EnableAuthTokenHeader: true,
	},
	Secrets: sharedconfig.Secrets{
		Keys: []string{"needsToBeAtLeast16Test"},
	},
}

type sessionManager struct{}

var userId = "ec4ef049-5b88-4321-a173-21b0eff06a04"

func (s sessionManager) GenerateJWT(uuid string) (string, error) {
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

var uId, _ = uuid.FromString(userId)

var emails = []*ent.Email{
	{
		ID:      uId,
		UserID:  &uId,
		Address: "john.doe@example.com",
	},
}

var users = []*ent.User{
	func() *ent.User {
		user := &ent.User{
			ID:        uId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		primE := &ent.PrimaryEmail{ID: uId, UserID: &uId}
		primE.Edges.Email = emails[0]
		user.Edges.PrimaryEmail = primE
		user.Edges.Emails = emails
		return user
	}(),
}
