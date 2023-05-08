// Package config provides configuration shared by microservices.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package sharedConfig

import (
	"fmt"
	"strconv"

	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
	"github.com/powerman/pqx"
)

// EnvPrefix defines common prefix for environment variables.
const EnvPrefix = "HQ_"

// Shared contains configurable values shared by microservices.
type Shared struct {
	AuthAddrHost    appcfg.NotEmptyString
	AuthAddrHostInt appcfg.NotEmptyString
	AuthAddrPort    appcfg.Port

	NetworthAddrHost    appcfg.NotEmptyString
	NetworthAddrHostInt appcfg.NotEmptyString
	NetworthAddrPort    appcfg.Port

	Postgres *PostgresConfig
	Session  Session
	Secrets  Secrets
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

// Default ports.
const (
	AuthPort = 17000 + iota
	NetworthPort
)

var shared = &struct {
	AuthAddrHost    appcfg.NotEmptyString `env:"AUTH_ADDR_HOST"`
	AuthAddrHostInt appcfg.NotEmptyString `env:"AUTH_ADDR_HOST_INT"`
	AuthAddrPort    appcfg.Port           `env:"AUTH_ADDR_PORT"`

	NetworthAddrHost    appcfg.NotEmptyString `env:"NETWORTH_ADDR_HOST"`
	NetworthAddrHostInt appcfg.NotEmptyString `env:"NETWORTH_ADDR_HOST_INT"`
	NetworthAddrPort    appcfg.Port           `env:"NETWORTH_ADDR_PORT"`

	PostgresUser     appcfg.NotEmptyString `env:"POSTGRES_AUTH_LOGIN"`
	PostgresPass     appcfg.NotEmptyString `env:"POSTGRES_AUTH_PASS"`
	PostgresAddrHost appcfg.NotEmptyString `env:"POSTGRES_ADDR_HOST"`
	PostgresAddrPort appcfg.Port           `env:"POSTGRES_ADDR_PORT"`
	PostgresDBName   appcfg.NotEmptyString `env:"POSTGRES_DB_NAME"`

	Secrets appcfg.NotEmptyString `env:"AUTH_SECRETS"`
}{ //nolint:gochecknoglobals // Config is global anyway.
	PostgresUser:     appcfg.MustNotEmptyString("auth"),
	PostgresAddrPort: appcfg.MustPort("5432"),
	PostgresAddrHost: appcfg.MustNotEmptyString("localhost"),
	PostgresDBName:   appcfg.MustNotEmptyString("postgres"),

	Secrets: appcfg.MustNotEmptyString("needsToBeAtLeast16"),
}

// Get updates config defaults (from env) and returns shared config.
func Get() (*Shared, error) {
	defer cleanup()

	fromEnv := appcfg.NewFromEnv(EnvPrefix)
	err := appcfg.ProvideStruct(shared, fromEnv)
	if err != nil {
		return nil, err
	}

	sharedCfg := &Shared{
		AuthAddrHost:    appcfg.MustNotEmptyString(def.Hostname),
		AuthAddrHostInt: appcfg.MustNotEmptyString(def.HostnameInt),
		AuthAddrPort:    appcfg.MustPort(strconv.Itoa(AuthPort)),

		NetworthAddrHost:    appcfg.MustNotEmptyString(def.Hostname),
		NetworthAddrHostInt: appcfg.MustNotEmptyString(def.Hostname),
		NetworthAddrPort:    appcfg.MustPort(strconv.Itoa(NetworthPort)),
		Postgres: NewPostgresConfig(pqx.Config{
			Host:   shared.PostgresAddrHost.Value(&err),
			Port:   shared.PostgresAddrPort.Value(&err),
			DBName: shared.PostgresDBName.Value(&err),
			User:   shared.PostgresUser.Value(&err),
			Pass:   shared.PostgresPass.Value(&err),
		}),

		Session: Session{
			Lifespan: "1h",
			Cookie: Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: true,
		},
		Secrets: Secrets{
			Keys: []string{shared.Secrets.Value(&err)},
		},
	}
	err = sharedCfg.Validate()
	if err != nil {
		return nil, err
	}

	return sharedCfg, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	shared = nil
}
