package middlewares

import (
	"net"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/felixge/httpsnoop"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
)

type middlewareFunc func(http.Handler) http.Handler

// Provide a logger configured using request's context.
//
// Usually it should be one of the first (but after xff, if used) middleware.
func MakeLogger(basePath string) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := structlog.FromContext(r.Context(), nil)
			remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
			log.SetDefaultKeyvals(
				structlog.KeyApp, config.ServiceName,
				def.LogRemoteIP, remoteIP,
				def.LogHTTPStatus, "",
				def.LogHTTPMethod, r.Method,
				def.LogFunc, path.Join("/", strings.TrimPrefix(r.URL.Path, basePath)),
			)
			r = r.WithContext(structlog.NewContext(r.Context(), log))

			next.ServeHTTP(w, r)
		})
	}
}

func MakeAccessLog(basePath string) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			metric.reqInFlight.Inc()
			defer metric.reqInFlight.Dec()

			m := httpsnoop.CaptureMetrics(next, w, r)

			l := prometheus.Labels{
				resourceLabel: strings.TrimPrefix(r.URL.Path, basePath),
				methodLabel:   r.Method,
				codeLabel:     strconv.Itoa(m.Code),
			}
			metric.reqTotal.With(l).Inc()
			l = prometheus.Labels{
				resourceLabel: strings.TrimPrefix(r.URL.Path, basePath),
				methodLabel:   r.Method,
				failedLabel:   strconv.FormatBool(m.Code >= http.StatusInternalServerError),
			}
			metric.reqDuration.With(l).Observe(m.Duration.Seconds())

			log := structlog.FromContext(r.Context(), nil)
			if m.Code < http.StatusInternalServerError {
				log.Info("handled", def.LogHTTPStatus, m.Code)
			} else {
				log.PrintErr("failed to handle", def.LogHTTPStatus, m.Code)
			}
		})
	}
}
