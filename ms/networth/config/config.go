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
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet
}

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	fs     FlagSets
	shared *sharedconfig.Shared
	own    = &struct {
		// Below envs is loaded by Doppler
		// 	// See https://dashboard.plaid.com/account/keys
		PlaidClientId appcfg.String `env:"PLAID_CLIENT_ID"`
		PlaidSecret   appcfg.String `env:"PLAID_SECRET"`
		// See sandbox, development, product
		PlaidEnv appcfg.String `env:"PLAID_ENV"`
		// See https://plaid.com/docs/api/tokens/#link-token-create-request-products
		PlaidProducts appcfg.String `env:"PLAID_PRODUCTS"`
		// See https://plaid.com/docs/api/tokens/#link-token-create-request-country-codes
		PlaidCountryCodes appcfg.String `env:"PLAID_COUNTRY_CODES"`
		// See https://dashboard.plaid.com/team/api
		PlaidRedirectUri appcfg.String `env:"PLAID_REDIRECT_URI"`

		SeAppID  appcfg.String `env:"SALTEDGE_APP_ID"`
		SeSecret appcfg.String `env:"SALTEDGE_SECRET"`
		SePK     appcfg.String `env:"SALTEDGE_PK"`

		FvAppID       appcfg.String `env:"FINVERSE_APP_ID"`
		FvClientID    appcfg.String `env:"FINVERSE_CLIENT_ID"`
		FvSecret      appcfg.String `env:"FINVERSE_SECRET"`
		FvRedirectURI appcfg.String `env:"FINVERSE_REDIRECT_URI"`
	}{}
)

type Config struct {
	Server   Server
	Plaid    *PlaidConfig // TODO: Need to finalize Plaid integration
	SaltEdge *SaltEdge
	Finverse *Finverse
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

	appcfg.AddPFlag(fs.Serve, &shared.NetworthAddrHost, "networth.host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.NetworthAddrHostInt, "networth.host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.NetworthAddrPort, "networth.port", "port to serve monolith introspection")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()
	c = &Config{
		Server: Server{
			BindAddr:    netx.NewAddr(shared.NetworthAddrHost.Value(&err), shared.NetworthAddrPort.Value(&err)),
			BindAddrInt: netx.NewAddr(shared.NetworthAddrHostInt.Value(&err), shared.NetworthAddrPort.Value(&err)),
			Cors: Cors{
				ExposeHeaders: []string{
					httplimit.HeaderRateLimitLimit,
					httplimit.HeaderRateLimitRemaining,
					httplimit.HeaderRateLimitReset,
					httplimit.HeaderRetryAfter,
				},
			},
		},
		SaltEdge: &SaltEdge{
			AppId:  own.SeAppID.Value(&err),
			Secret: own.SeSecret.Value(&err),
			PK:     own.SePK.Value(&err),
		},
		Finverse: &Finverse{
			AppId:       own.FvAppID.Value(&err),
			ClientID:    own.FvClientID.Value(&err),
			Secret:      own.FvSecret.Value(&err),
			RedirectURI: own.FvRedirectURI.Value(&err),
		},
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own)
	}

	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}
