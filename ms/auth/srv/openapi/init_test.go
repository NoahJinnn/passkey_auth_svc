package openapi_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hellohq/hqservice/api/openapi/client"
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

func testNewServer(t *check.C, cfg openapi.Config) (c *client.HqService, url string, mockAppl *app.MockAppl, logc <-chan string) {
	cfg.Addr = netx.NewAddr("localhost", 0)

	t.Helper()
	ctrl := gomock.NewController(t)

	mockAppl = app.NewMockAppl(ctrl)

	server, err := openapi.NewServer(mockAppl, cfg)
	t.Must(t.Nil(err, "NewServer"))

	piper, pipew := io.Pipe()
	server.SetHandler(interceptLog(pipew, server.GetHandler()))
	logch := make(chan string, 64) // Keep some unread log messages.
	go func() {
		scanner := bufio.NewScanner(piper)
		for scanner.Scan() {
			select {
			default: // Do not hang test because of some unread log messages.
			case logch <- scanner.Text():
			}
		}
		close(logch)
	}()

	t.Must(t.Nil(server.Listen(), "server.Listen"))
	errc := make(chan error, 1)
	go func() { errc <- server.Serve() }()

	t.Cleanup(func() {
		t.Helper()
		t.Nil(server.Shutdown(), "server.Shutdown")
		t.Nil(<-errc, "server.Serve")
		pipew.Close()
	})

	ln, err := server.HTTPListener()
	t.Must(t.Nil(err, "server.HTTPListener"))
	c = client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     ln.Addr().String(),
		BasePath: client.DefaultBasePath,
	})
	url = fmt.Sprintf("http://%s", ln.Addr().String())

	// Avoid race between server.Serve() and server.Shutdown().
	ctx, cancel := context.WithTimeout(context.Background(), def.TestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	t.Must(t.Nil(err))
	_, err = (&http.Client{}).Do(req)
	t.Must(t.Nil(err, "connect to service"))
	<-logch

	return c, url, mockAppl, logch
}

func interceptLog(out io.Writer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := structlog.FromContext(r.Context(), nil)
		log.SetOutput(out)
		r = r.WithContext(structlog.NewContext(r.Context(), log))
		next.ServeHTTP(w, r)
	})
}
