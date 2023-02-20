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

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	ServiceName = "hq"
	fs          *pflag.FlagSet
	shared      *sharedconfig.Shared
	own         = &struct {
		Port            appcfg.Port           `env:"HQ_ADDR_PORT"`
		MetricsAddrPort appcfg.Port           `env:"HQ_METRICS_ADDR_PORT"`
		PostgresUser    appcfg.NotEmptyString `env:"HQ_POSTGRES_AUTH_LOGIN"`
		PostgresPass    appcfg.NotEmptyString `env:"HQ_POSTGRES_AUTH_PASS"`
	}{
		Port:            appcfg.MustPort(strconv.Itoa(sharedconfig.MonoPort)),
		MetricsAddrPort: appcfg.MustPort(strconv.Itoa(sharedconfig.MetricsPort)),
		PostgresUser:    appcfg.MustNotEmptyString(ServiceName),
	}
)

type Config struct {
	BindAddr        netx.Addr
	BindMetricsAddr netx.Addr
	Postgres        *PostgresConfig
}

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(svcName string, flagsets *pflag.FlagSet, sharedCfg *sharedconfig.Shared) error {
	fs, shared = flagsets, sharedCfg
	fromEnv := appcfg.NewFromEnv(sharedconfig.EnvPrefix)
	err := appcfg.ProvideStruct(own, fromEnv)
	if err != nil {
		return err
	}

	pfx := svcName + "."
	appcfg.AddPFlag(fs, &shared.AddrHostInt, "host-int", "internal host to serve")
	appcfg.AddPFlag(fs, &own.Port, pfx+"port", "port to serve monolith introspection")
	appcfg.AddPFlag(fs, &shared.XPostgresAddrHost, "postgres.host", "host to connect to PostgreSQL")
	appcfg.AddPFlag(fs, &shared.XPostgresAddrPort, "postgres.port", "port to connect to PostgreSQL")
	appcfg.AddPFlag(fs, &shared.XPostgresDBName, "postgres.dbname", "PostgreSQL database name")
	appcfg.AddPFlag(fs, &own.PostgresUser, pfx+"postgres.user", "PostgreSQL username")
	appcfg.AddPFlag(fs, &own.PostgresPass, pfx+"postgres.pass", "PostgreSQL password")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()

	c = &Config{
		BindAddr:        netx.NewAddr(shared.AddrHostInt.Value(&err), own.Port.Value(&err)),
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
		return nil, appcfg.WrapPErr(err, fs, own)
	}
	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
}
