// Package config provides configurations for subcommands.
//
// It consists of both configuration values shared by all
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
	"github.com/spf13/pflag"
)

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	ServiceName = "hq"
	fs          *pflag.FlagSet
	shared      *sharedconfig.Shared
	all         = &struct {
		Port            appcfg.Port `env:"HQ_ADDR_PORT"`
		MetricsAddrPort appcfg.Port `env:"HQ_METRICS_ADDR_PORT"`
	}{
		Port:            appcfg.MustPort(strconv.Itoa(sharedconfig.MonoPort)),
		MetricsAddrPort: appcfg.MustPort(strconv.Itoa(sharedconfig.MetricsPort)),
	}
)

type Config struct {
	BindAddr        netx.Addr
	BindMetricsAddr netx.Addr
}

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(svcName string, flagsets *pflag.FlagSet, sharedCfg *sharedconfig.Shared) error {
	fs, shared = flagsets, sharedCfg
	fromEnv := appcfg.NewFromEnv(sharedconfig.EnvPrefix)
	err := appcfg.ProvideStruct(all, fromEnv)
	if err != nil {
		return err
	}

	pfx := svcName + "."
	appcfg.AddPFlag(fs, &shared.AddrHostInt, "host-int", "internal host to serve")
	appcfg.AddPFlag(fs, &all.Port, pfx+"port", "port to serve monolith introspection")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()

	c = &Config{
		BindAddr:        netx.NewAddr(shared.AddrHostInt.Value(&err), all.Port.Value(&err)),
		BindMetricsAddr: netx.NewAddr(shared.AddrHostInt.Value(&err), all.MetricsAddrPort.Value(&err)),
	}

	if err != nil {
		return nil, appcfg.WrapPErr(err, fs, all)
	}
	return c, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	all = nil
}
