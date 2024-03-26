// Package config provides configuration shared by microservices.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package sharedconfig

import (
	"fmt"
	"strconv"

	"github.com/NoahJinnn/passkey_auth_svc/pkg/def"
	"github.com/pkg/errors"
	"github.com/powerman/appcfg"
	"github.com/powerman/pqx"
)

const (
	// EnvPrefix defines common prefix for environment variables.
	EnvPrefix = "HQ_"
	// Default ports
	AuthPort = 17000 + iota*2
)

// Shared contains configurable values shared by microservices.
type Shared struct {
	AuthAddrHost    appcfg.NotEmptyString
	AuthAddrHostInt appcfg.NotEmptyString
	AuthAddrPort    appcfg.Port

	Postgres *PostgresConfig
	Session  Session
	Secrets  Secrets
}

var shared = &struct {
	AuthAddrHost    appcfg.NotEmptyString `env:"AUTH_ADDR_HOST"`
	AuthAddrHostInt appcfg.NotEmptyString `env:"AUTH_ADDR_HOST_INT"`
	AuthAddrPort    appcfg.Port           `env:"AUTH_ADDR_PORT"`

	PostgresUser   appcfg.NotEmptyString `env:"X_POSTGRES_LOGIN"`
	PostgresPass   appcfg.NotEmptyString `env:"X_POSTGRES_PASS"`
	PostgresHost   appcfg.NotEmptyString `env:"X_POSTGRES_HOST"`
	PostgresPort   appcfg.Port           `env:"X_POSTGRES_PORT"`
	PostgresDBName appcfg.NotEmptyString `env:"X_POSTGRES_NAME"`

	Secrets appcfg.NotEmptyString `env:"JWK_SECRETS"`

	SessionLifespan       appcfg.String      `env:"JWT_LIFESPAN"`
	Issuer                appcfg.String      `env:"JWT_ISSUER"`
	Audience              appcfg.StringSlice `env:"JWT_AUDIENCE"`
	EnableAuthTokenHeader appcfg.Bool        `env:"JWT_ENABLE_TOKEN_HEADER"`
}{ //nolint:gochecknoglobals // Config is global anyway.
	AuthAddrHost:    appcfg.MustNotEmptyString(def.Hostname),
	AuthAddrHostInt: appcfg.MustNotEmptyString(def.HostnameInt),
	AuthAddrPort:    appcfg.MustPort(strconv.Itoa(AuthPort)),

	PostgresUser:   appcfg.MustNotEmptyString("auth"),
	PostgresPort:   appcfg.MustPort("5432"),
	PostgresHost:   appcfg.MustNotEmptyString("localhost"),
	PostgresDBName: appcfg.MustNotEmptyString("postgres"),
	Secrets:        appcfg.MustNotEmptyString("needsToBeAtLeast16"),
}

// Get updates config defaults (from env) and returns shared config.
func Get() (*Shared, error) {
	defer cleanup()

	fromEnv := appcfg.NewFromEnv(EnvPrefix)
	err := appcfg.ProvideStruct(shared, fromEnv)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sharedCfg := &Shared{
		AuthAddrHost:    shared.AuthAddrHost,
		AuthAddrHostInt: shared.AuthAddrHostInt,
		AuthAddrPort:    shared.AuthAddrPort,

		Postgres: NewPostgresConfig(pqx.Config{
			Host:   shared.PostgresHost.Value(&err),
			Port:   shared.PostgresPort.Value(&err),
			DBName: shared.PostgresDBName.Value(&err),
			User:   shared.PostgresUser.Value(&err),
			Pass:   shared.PostgresPass.Value(&err),
		}),

		Session: Session{
			Lifespan: shared.SessionLifespan.Value(&err),
			Cookie: Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: shared.EnableAuthTokenHeader.Value(&err),
			Issuer:                shared.Issuer.Value(&err),
			Audience:              shared.Audience.Value(&err),
		},
		Secrets: Secrets{
			Keys: []string{shared.Secrets.Value(&err)},
		},
	}
	err = sharedCfg.Validate()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return sharedCfg, nil
}

func (c *Shared) Validate() error {
	err := c.Session.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate session settings: %w", err)
	}
	err = c.Secrets.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate secrets settings: %w", err)
	}

	return nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	shared = nil
}
