package openapi_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/srv/openapi"
	"github.com/hellohq/hqservice/ms/auth/srv/openapi/middlewares"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/check"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	//nolint:errcheck
	reg := prometheus.NewPedanticRegistry()
	def.Init()
	app.InitMetrics(reg)
	middlewares.InitMetrics(reg, "test")
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	// nolint:unused
	apiError500 = openapi.APIError(500, "internal error")
)

func testNewServer(t *check.C, cfg openapi.Config) {
	cfg.Addr = netx.NewAddr("localhost", 0)

	t.Helper()
	// ctrl := gomock.NewController(t)

	// mockAppl = app.NewMockAppl(ctrl)

	// server, err := openapi.NewServer(mockAppl, cfg)
	// t.Must(t.Nil(err, "NewServer"))

}

func interceptLog(out io.Writer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := structlog.FromContext(r.Context(), nil)
		log.SetOutput(out)
		r = r.WithContext(structlog.NewContext(r.Context(), log))
		next.ServeHTTP(w, r)
	})
}
