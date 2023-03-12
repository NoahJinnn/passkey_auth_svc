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
	"strconv"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/powerman/pqx"
	"github.com/spf13/pflag"
)

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet
}

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	ServiceName = "auth"
	fs          FlagSets
	shared      *sharedconfig.Shared
	own         = &struct {
		MetricsAddrPort appcfg.Port           `env:"AUTH_METRICS_ADDR_PORT"`
		PostgresUser    appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_LOGIN"`
		PostgresPass    appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_PASS"`
	}{
		MetricsAddrPort: appcfg.MustPort(strconv.Itoa(sharedconfig.MetricsPort)),
		PostgresUser:    appcfg.MustNotEmptyString(ServiceName),
	}
)

type Config struct {
	BindAddr        netx.Addr
	BindAddrInt     netx.Addr
	BindMetricsAddr netx.Addr
	Webauthn        WebauthnSettings
	Session         Session
	Secrets         Secrets
	Postgres        *PostgresConfig
	Plaid           *PlaidConfig
}

// // Ref: https://github.com/plaid/quickstart/blob/master/.env.example

type PlaidConfig struct {
	// 	// See https://dashboard.plaid.com/account/keys
	ClientId appcfg.String `env:"PLAID_CLIENT_ID"`
	Secret   appcfg.String `env:"PLAID_SECRET"`
	// See sandbox, development, product
	Env appcfg.String `env:"PLAID_ENV"`
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-products
	Products appcfg.String `env:"PLAID_PRODUCTS"`
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-country-codes
	CountryCodes appcfg.String `env:"PLAID_COUNTRY_CODES"`
	// See https://dashboard.plaid.com/team/api
	RedirectUri appcfg.String `env:"PLAID_REDIRECT_URI"`
}

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(sharedCfg *sharedconfig.Shared, flagsets FlagSets) error {
	shared, fs = sharedCfg, flagsets
	fromEnv := appcfg.NewFromEnv(sharedconfig.EnvPrefix)
	err := appcfg.ProvideStruct(own, fromEnv)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHost, "host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPort, "host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPortInt, "port", "port to serve monolith introspection")
	appcfg.AddPFlag(fs.Serve, &shared.XPostgresAddrHost, "postgres.host", "host to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &shared.XPostgresAddrPort, "postgres.port", "port to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &shared.XPostgresDBName, "postgres.dbname", "PostgreSQL database name")
	appcfg.AddPFlag(fs.Serve, &own.PostgresUser, "postgres.user", "PostgreSQL username")
	appcfg.AddPFlag(fs.Serve, &own.PostgresPass, "postgres.pass", "PostgreSQL password")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()

	c = &Config{
		BindAddr:        netx.NewAddr(shared.AuthAddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
		BindMetricsAddr: netx.NewAddr(shared.AuthAddrHostInt.Value(&err), own.MetricsAddrPort.Value(&err)),
		Postgres: NewPostgresConfig(pqx.Config{
			Host:   shared.XPostgresAddrHost.Value(&err),
			Port:   shared.XPostgresAddrPort.Value(&err),
			DBName: shared.XPostgresDBName.Value(&err),
			User:   own.PostgresUser.Value(&err),
			Pass:   own.PostgresPass.Value(&err),
		}),
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

	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own, shared)
	}
	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}
