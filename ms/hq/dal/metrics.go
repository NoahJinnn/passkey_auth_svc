package dal

import (
	"github.com/hellohq/hqservice/ms/hq/app"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/ms/hq/dal/repo"
	"github.com/prometheus/client_golang/prometheus"
)

var metric repo.Metrics //nolint:gochecknoglobals // Metrics are global anyway.

// InitMetrics must be called once before using this package.
// It registers and initializes metrics used by this package.
func InitMetrics(reg *prometheus.Registry) {
	const subsystem = "dal_postgres"

	metric = repo.NewMetrics(reg, config.ServiceName, subsystem, new(app.Repo))
}
