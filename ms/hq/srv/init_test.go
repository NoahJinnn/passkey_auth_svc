package srv

import (
	"testing"

	"github.com/hellohq/hqservice/ms/hq/app"
	"github.com/hellohq/hqservice/ms/hq/srv/openapi"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	//nolint:errcheck
	def.Init()
	initMetrics(reg, "test")
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, "test")
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	// nolint:unused
	ctx = def.NewContext((&Service{}).Name())
)
