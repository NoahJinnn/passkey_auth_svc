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
		Port            appcfg.Port           `env:"AUTH_ADDR_PORT"`
		MetricsAddrPort appcfg.Port           `env:"AUTH_METRICS_ADDR_PORT"`
		PostgresUser    appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_LOGIN"`
		PostgresPass    appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_PASS"`
	}{
		Port:            appcfg.MustPort(strconv.Itoa(sharedconfig.MonoPort)),
		MetricsAddrPort: appcfg.MustPort(strconv.Itoa(sharedconfig.MetricsPort)),
		PostgresUser:    appcfg.MustNotEmptyString(ServiceName),
	}
)

type Config struct {
	AuthAddr        netx.Addr
	BindAddr        netx.Addr
	BindAddrInt     netx.Addr
	BindMetricsAddr netx.Addr
	Postgres        *PostgresConfig
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

	appcfg.AddPFlag(fs.Serve, &shared.AddrHost, "host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AddrHostInt, "host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHost, "auth.host", "ms/auth API host")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPort, "auth.port", "ms/auth API port")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPortInt, "auth.port-int", "ms/auth internal API port")
	appcfg.AddPFlag(fs.Serve, &own.Port, "port", "port to serve monolith introspection")
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
		AuthAddr:        netx.NewAddr(shared.AuthAddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
		BindAddr:        netx.NewAddr(shared.AddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
		BindAddrInt:     netx.NewAddr(shared.AddrHostInt.Value(&err), shared.AuthAddrPortInt.Value(&err)),
		BindMetricsAddr: netx.NewAddr(shared.AddrHostInt.Value(&err), own.MetricsAddrPort.Value(&err)),
		Postgres: NewPostgresConfig(pqx.Config{
			Host:   shared.XPostgresAddrHost.Value(&err),
			Port:   shared.XPostgresAddrPort.Value(&err),
			DBName: shared.XPostgresDBName.Value(&err),
			User:   own.PostgresUser.Value(&err),
			Pass:   own.PostgresPass.Value(&err),
		}),
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
