package app_test

import (
	"context"
	"testing"

	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	"github.com/prometheus/client_golang/prometheus"
)

func TestMain(m *testing.M) {
	def.Init()
	reg := prometheus.NewPedanticRegistry()
	app.InitMetrics(reg)
	check.TestMain(m)
}

type Ctx = context.Context

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	ctx = def.NewContext(config.ServiceName)
)

// func testNew(t *check.C) (*app.App, *app.MockRepo) {
// 	ctrl := gomock.NewController(t)

// 	mockRepo := app.NewMockRepo(ctrl)
// 	a, err := app.New(mockRepo)
// 	if err != nil {
// 		panic("Init app test failed!")
// 	}
// 	return a, mockRepo
// }
