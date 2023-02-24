package dal_test

import (
	"testing"

	"github.com/hellohq/hqservice/ms/hq/app"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/ms/hq/dal"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	"github.com/prometheus/client_golang/prometheus"
)

func TestMain(m *testing.M) {
	def.Init()
	reg := prometheus.NewPedanticRegistry()
	app.InitMetrics(reg)
	dal.InitMetrics(reg)
	cfg = config.MustGetServeTest()
	check.TestMain(m)
}

// type tLogger check.C

// func (t tLogger) Print(args ...interface{}) { t.Log(args...) }

var (
	// ctx = def.NewContext(config.ServiceName)
	cfg *config.Config
)

func newTestRepo(t *check.C) *dal.Repo {
	t.Helper()
	// TODO: Implement enttest
	// r, err := dal.New(ctx, cfg.Postgres)

	// t.Must(t.Nil(err))
	// t.Cleanup(r.Close)

	return &dal.Repo{}
}

// func matchErr(t *check.C, err, wantErr error) {
// 	t.Helper()
// 	if pqErr := new(*pq.Error); errors.As(err, pqErr) && wantErr != nil {
// 		t.Match(err, wantErr.Error())
// 	} else {
// 		t.Err(err, wantErr)
// 	}
// }
