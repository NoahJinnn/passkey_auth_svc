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
const EnvPrefix = "HQ_"

// Shared contains configurable values shared by microservices.
type Shared struct {
	AuthAddrHost    appcfg.NotEmptyString `env:"AUTH_ADDR_HOST"`
	AuthAddrHostInt appcfg.NotEmptyString `env:"AUTH_ADDR_HOST_INT"`
	AuthAddrPort    appcfg.Port           `env:"AUTH_ADDR_PORT"`
}

// Default ports.
const (
	AuthPort = 17000 + iota
)

var shared = &Shared{ //nolint:gochecknoglobals // Config is global anyway.
	AuthAddrHost:    appcfg.MustNotEmptyString(def.Hostname),
	AuthAddrHostInt: appcfg.MustNotEmptyString(def.Hostname),
	AuthAddrPort:    appcfg.MustPort(strconv.Itoa(AuthPort)),
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
