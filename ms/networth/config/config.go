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
	shared *sharedConfig.Shared
	own    = &struct {
		// Below envs is loaded by Doppler
		Secrets appcfg.NotEmptyString `env:"AUTH_SECRETS"`
	}{

		Secrets: appcfg.MustNotEmptyString("needsToBeAtLeast16"),
	}
)

type Config struct {
	Server  Server
	Session authCfg.Session
	Secrets authCfg.Secrets
	Plaid   *PlaidConfig
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
		Session: authCfg.Session{
			Lifespan: "1h",
			Cookie: authCfg.Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: true,
		},
		Secrets: authCfg.Secrets{
			Keys: []string{own.Secrets.Value(&err)},
		},
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own)
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
