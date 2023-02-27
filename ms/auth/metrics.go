package hq

import (
	"regexp"
	"runtime"

	"github.com/powerman/go-monolith-example/pkg/def"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func initMetrics(reg *prometheus.Registry, namespace string) {

	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	version := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "build_info",
			Help:      "A metric with a constant '1' value labeled by build-time details.",
		},
		[]string{"version", "goversion"},
	)
	reg.MustRegister(version)

	version.With(prometheus.Labels{
		"version":   def.Version(),
		"goversion": runtime.Version(),
	}).Set(1)
}
