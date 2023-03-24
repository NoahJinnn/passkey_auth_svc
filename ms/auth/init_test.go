package hq

import (
	"testing"

	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/middlewares"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	def.Init()
	dal.InitMetrics(reg, "test")
	app.InitMetrics(reg)
	middlewares.InitMetrics(reg, "test")
	initMetrics(reg, "test")
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	ctx = def.NewContext((&Service{}).Name())
)
