package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ms/auth/app/crypto"
	hqJwt "github.com/hellohq/hqservice/pkg/crypto/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Config struct {
	Webauthn WebauthnSettings
	Session  Session `yaml:"session" json:"session" koanf:"session"`
}

func DefaultConfig() *Config {
	return &Config{

		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				Id:          "localhost",
				DisplayName: "Hanko Authentication Service",
				Origins:     []string{"http://localhost"},
			},
			Timeout: 60000,
		},

		Session: Session{
			Lifespan: "1h",
			Cookie: Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
		},
	}
}

func (c *Config) Validate() error {

	err := c.Webauthn.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate webauthn settings: %w", err)
	}

	err = c.Session.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate session settings: %w", err)
	}

	return nil
}

// WebauthnSettings defines the settings for the webauthn authentication mechanism
type WebauthnSettings struct {
	RelyingParty RelyingParty `yaml:"relying_party" json:"relying_party" koanf:"relying_party" split_words:"true"`
	Timeout      int          `yaml:"timeout" json:"timeout" koanf:"timeout"`
}

// Validate does not need to validate the config, because the library does this already
func (r *WebauthnSettings) Validate() error {
	return nil
}

// RelyingParty webauthn settings for your application using hanko.
type RelyingParty struct {
	Id          string `yaml:"id" json:"id" koanf:"id"`
	DisplayName string `yaml:"display_name" json:"display_name" koanf:"display_name" split_words:"true"`
	Icon        string `yaml:"icon" json:"icon" koanf:"icon"`
	// Deprecated: Use Origins instead
	Origin  string   `yaml:"origin" json:"origin" koanf:"origin"`
	Origins []string `yaml:"origins" json:"origins" koanf:"origins"`
}

type Session struct {
	EnableAuthTokenHeader bool   `yaml:"enable_auth_token_header" json:"enable_auth_token_header" koanf:"enable_auth_token_header" split_words:"true"`
	Lifespan              string `yaml:"lifespan" json:"lifespan" koanf:"lifespan"`
	Cookie                Cookie `yaml:"cookie" json:"cookie" koanf:"cookie"`
}

type Cookie struct {
	Domain   string `yaml:"domain" json:"domain" koanf:"domain"`
	HttpOnly bool   `yaml:"http_only" json:"http_only" koanf:"http_only" split_words:"true"`
	SameSite string `yaml:"same_site" json:"same_site" koanf:"same_site" split_words:"true"`
	Secure   bool   `yaml:"secure" json:"secure" koanf:"secure"`
}

func (s *Session) Validate() error {
	_, err := time.ParseDuration(s.Lifespan)
	if err != nil {
		return errors.New("failed to parse lifespan")
	}
	return nil
}

type Manager interface {
	GenerateJWT(uuid.UUID) (string, error)
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
func NewManager(jwkManager crypto.JwkManager, config Session) (Manager, error) {
	signatureKey, err := jwkManager.GetSigningKey()
	if err != nil {
		return nil, fmt.Errorf("failed to create session generator: %w", err)
	}
	verificationKeys, err := jwkManager.GetPublicKeys()
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
func (g *manager) GenerateJWT(userId uuid.UUID) (string, error) {
	issuedAt := time.Now()
	expiration := issuedAt.Add(g.sessionLength)

	token := jwt.New()
	_ = token.Set(jwt.SubjectKey, userId.String())
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
