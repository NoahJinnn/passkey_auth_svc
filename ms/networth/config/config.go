// Package config provides configurations for subcommands.
//
// It consists of both configuration values shared by own
// microservices and values specific to this microservice.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.

package config

import (
	"fmt"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	authCfg "github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/powerman/pqx"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet
}

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	ServiceName = "networth"
	fs          FlagSets
	shared      *sharedConfig.Shared
	own         = &struct {
		// Below envs is loaded by Doppler
		PostgresUser     appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_LOGIN"`
		PostgresPass     appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_PASS"`
		PostgresAddrHost appcfg.NotEmptyString `env:"AUTH_POSTGRES_ADDR_HOST"`
		PostgresAddrPort appcfg.Port           `env:"AUTH_POSTGRES_ADDR_PORT"`
		PostgresDBName   appcfg.NotEmptyString `env:"AUTH_POSTGRES_DB_NAME"`
		Secrets          appcfg.NotEmptyString `env:"AUTH_SECRETS"`
	}{
		PostgresUser:     appcfg.MustNotEmptyString(ServiceName),
		PostgresAddrPort: appcfg.MustPort("5432"),
		PostgresAddrHost: appcfg.MustNotEmptyString("localhost"),
		PostgresDBName:   appcfg.MustNotEmptyString("postgres"),
		Secrets:          appcfg.MustNotEmptyString("needsToBeAtLeast16"),
	}
)

type Config struct {
	Server      Server
	Session     authCfg.Session
	Secrets     authCfg.Secrets
	ServiceName string
	Postgres    *PostgresConfig
	Plaid       *PlaidConfig
}

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(sharedCfg *sharedConfig.Shared, flagsets FlagSets) error {
	shared, fs = sharedCfg, flagsets
	fromEnv := appcfg.NewFromEnv(sharedConfig.EnvPrefix)
	err := appcfg.ProvideStruct(own, fromEnv)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHost, "networth.host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHostInt, "networth.host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPort, "networth.port", "port to serve monolith introspection")
	appcfg.AddPFlag(fs.Serve, &own.PostgresAddrHost, "postgres.host", "host to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &own.PostgresAddrPort, "postgres.port", "port to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &own.PostgresDBName, "postgres.dbname", "PostgreSQL database name")
	appcfg.AddPFlag(fs.Serve, &own.PostgresUser, "postgres.user", "PostgreSQL username")
	appcfg.AddPFlag(fs.Serve, &own.PostgresPass, "postgres.pass", "PostgreSQL password")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()
	c = &Config{
		Server: Server{
			BindAddr:    netx.NewAddr(shared.AuthAddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
			BindAddrInt: netx.NewAddr(shared.AuthAddrHostInt.Value(&err), shared.AuthAddrPort.Value(&err)),
			Cors: Cors{
				ExposeHeaders: []string{
					httplimit.HeaderRateLimitLimit,
					httplimit.HeaderRateLimitRemaining,
					httplimit.HeaderRateLimitReset,
					httplimit.HeaderRetryAfter,
				},
			},
		},
		Postgres: NewPostgresConfig(pqx.Config{
			Host:   own.PostgresAddrHost.Value(&err),
			Port:   own.PostgresAddrPort.Value(&err),
			DBName: own.PostgresDBName.Value(&err),
			User:   own.PostgresUser.Value(&err),
			Pass:   own.PostgresPass.Value(&err),
		}),
		Session: authCfg.Session{
			Lifespan: "1h",
			Cookie: authCfg.Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: true,
		},
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own, shared)
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}

func (c *Config) Validate() error {
	err := c.Session.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate session settings: %w", err)
	}

	return nil
}
