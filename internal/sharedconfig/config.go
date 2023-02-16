// Package config provides configuration shared by microservices.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package sharedconfig

import (
	"strconv"

	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
)

// EnvPrefix defines common prefix for environment variables.
const EnvPrefix = "MONO_"

// Shared contains configurable values shared by microservices.
type Shared struct {
	AddrHost          appcfg.NotEmptyString `env:"ADDR_HOST"`
	AddrHostInt       appcfg.NotEmptyString `env:"ADDR_HOST_INT"`
	MetricsAddrPort   appcfg.Port           `env:"METRICS_ADDR_PORT"`
	TLSCACert         appcfg.NotEmptyString `env:"TLS_CA_CERT"`
	XPostgresAddrHost appcfg.NotEmptyString `env:"X_POSTGRES_ADDR_HOST"`
	XPostgresAddrPort appcfg.Port           `env:"X_POSTGRES_ADDR_PORT"`
	XPostgresDBName   appcfg.NotEmptyString `env:"X_POSTGRES_DB_NAME"`
}

// Default ports.
const (
	MonoPort = 17000 + iota
	MetricsPort
)

var shared = &Shared{ //nolint:gochecknoglobals // Config is global anyway.
	AddrHost:          appcfg.MustNotEmptyString(def.Hostname),
	AddrHostInt:       appcfg.MustNotEmptyString(def.Hostname),
	MetricsAddrPort:   appcfg.MustPort(strconv.Itoa(MetricsPort)),
	XPostgresAddrPort: appcfg.MustPort("5432"),
	XPostgresDBName:   appcfg.MustNotEmptyString("postgres"),
}

// Get updates config defaults (from env) and returns shared config.
func Get() (*Shared, error) {
	defer cleanup()

	fromEnv := appcfg.NewFromEnv(EnvPrefix)
	err := appcfg.ProvideStruct(shared, fromEnv)
	if err != nil {
		return nil, err
	}
	return shared, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	shared = nil
}
