package session

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	hqJwt "github.com/hellohq/hqservice/pkg/crypto/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Manager interface {
	GenerateJWT(string) (string, error)
	Verify(string) (jwt.Token, error)
	GenerateCookie(token string) (*http.Cookie, error)
	DeleteCookie() (*http.Cookie, error)
}

// Manager is used to create and verify session JWTs
type manager struct {
	jwtGenerator  hqJwt.Generator
	sessionLength time.Duration
	cookieConfig  cookieConfig
}

type cookieConfig struct {
	Domain   string
	HttpOnly bool
	SameSite http.SameSite
	Secure   bool
}

// NewManager returns a new Manager which will be used to create and verify sessions JWTs
func NewManager(jwkManager JwkManager, config sharedConfig.Session) (Manager, error) {
	ctx := context.Background()
	signatureKey, err := jwkManager.GetSigningKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create session generator: %w", err)
	}
	verificationKeys, err := jwkManager.GetPublicKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create session generator: %w", err)
	}
	g, err := hqJwt.NewGenerator(signatureKey, verificationKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to create session generator: %w", err)
	}

	duration, _ := time.ParseDuration(config.Lifespan) // error can be ignored, value is checked in config validation
	sameSite := http.SameSite(0)
	switch config.Cookie.SameSite {
	case "lax":
		sameSite = http.SameSiteLaxMode
	case "strict":
		sameSite = http.SameSiteStrictMode
	case "none":
		sameSite = http.SameSiteNoneMode
	default:
		sameSite = http.SameSiteDefaultMode
	}
	return &manager{
		jwtGenerator:  g,
		sessionLength: duration,
		cookieConfig: cookieConfig{
			Domain:   config.Cookie.Domain,
			HttpOnly: config.Cookie.HttpOnly,
			SameSite: sameSite,
			Secure:   config.Cookie.Secure,
		},
	}, nil
}

// GenerateJWT creates a new session JWT for the given user
func (g *manager) GenerateJWT(userId string) (string, error) {
	issuedAt := time.Now()
	expiration := issuedAt.Add(g.sessionLength)

	token := jwt.New()
	_ = token.Set(jwt.SubjectKey, userId)
	_ = token.Set(jwt.IssuedAtKey, issuedAt)
	_ = token.Set(jwt.ExpirationKey, expiration)
	//_ = token.Set(jwt.AudienceKey, []string{"http://localhost"})

	signed, err := g.jwtGenerator.Sign(token)
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

// Verify verifies the given JWT and returns a parsed one if verification was successful
func (g *manager) Verify(token string) (jwt.Token, error) {
	parsedToken, err := g.jwtGenerator.Verify([]byte(token))
	if err != nil {
		return nil, fmt.Errorf("failed to verify session token: %w", err)
	}

	return parsedToken, nil
}

// GenerateCookie creates a new session cookie for the given user
func (g *manager) GenerateCookie(token string) (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hqservice",
		Value:    token,
		Domain:   g.cookieConfig.Domain,
		Path:     "/",
		Secure:   g.cookieConfig.Secure,
		HttpOnly: g.cookieConfig.HttpOnly,
		SameSite: g.cookieConfig.SameSite,
	}, nil
}

// DeleteCookie returns a cookie that will expire the cookie on the frontend
func (g *manager) DeleteCookie() (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hqservice",
		Value:    "",
		Domain:   g.cookieConfig.Domain,
		Path:     "/",
		Secure:   g.cookieConfig.Secure,
		HttpOnly: g.cookieConfig.HttpOnly,
		SameSite: g.cookieConfig.SameSite,
		MaxAge:   -1,
	}, nil
}
