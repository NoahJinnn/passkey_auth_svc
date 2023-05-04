// Package def provides default values for both commands and tests.
package def

import (
	"time"

	"github.com/powerman/getenv"
	"github.com/powerman/must"
	"github.com/powerman/sensitive"
	"github.com/prometheus/client_golang/prometheus"
)

func init() { //nolint:gochecknoinits // Ensure time.Now() assigned to global vars uses UTC.
	// Make time.Now()==time.Now().UTC() https://github.com/golang/go/issues/19486
	time.Local = nil
}

// Init must be called once before using this package.
// It provides common initialization for both commands and tests.
func Init() error {
	// Make sure no one occasionally uses global objects.
	prometheus.DefaultRegisterer = nil
	prometheus.DefaultGatherer = nil
	must.AbortIf = must.PanicIf
	sensitive.Redact()
	setupLog()
	return getenv.LastErr()
}
