package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metric contains general metrics for OpenAPI methods.
var metric struct { //nolint:gochecknoglobals // Metrics are global anyway.
	reqInFlight prometheus.Gauge
	reqTotal    *prometheus.CounterVec
	reqDuration *prometheus.HistogramVec
}

const (
	resourceLabel = "resource"
	methodLabel   = "method"
	codeLabel     = "code"
	failedLabel   = "failed"
)

var (
	// Initialized with codes returned by go-swagger and middlewares
	// after metrics middleware (accessLog).
	CodeLabels = []int{400, 401, 403, 422}
)

// InitMetrics must be called once before using this package.
// It registers and initializes metrics used by this package.
func InitMetrics(reg *prometheus.Registry, namespace string) {
	const subsystem = "openapi"

	metric.reqInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_requests_in_flight",
			Help:      "Amount of currently processing API requests.",
		},
	)
	reg.MustRegister(metric.reqInFlight)
	metric.reqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_requests_total",
			Help:      "Amount of processed API requests.",
		},
		[]string{resourceLabel, methodLabel, codeLabel},
	)
	reg.MustRegister(metric.reqTotal)
	metric.reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_request_duration_seconds",
			Help:      "API request latency distributions.",
		},
		[]string{resourceLabel, methodLabel, failedLabel},
	)
	reg.MustRegister(metric.reqDuration)

}
