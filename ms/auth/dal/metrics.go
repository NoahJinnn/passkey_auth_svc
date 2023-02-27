package dal

import (
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/pkg/repo"
	"github.com/prometheus/client_golang/prometheus"
)

// var metric repo.Metrics //nolint:gochecknoglobals // Metrics are global anyway.

// InitMetrics must be called once before using this package.
// It registers and initializes metrics used by this package.
func InitMetrics(reg *prometheus.Registry) {
	const subsystem = "dal_mysql"

	repo.NewMetrics(reg, config.ServiceName, subsystem, new(app.Repo))
}
