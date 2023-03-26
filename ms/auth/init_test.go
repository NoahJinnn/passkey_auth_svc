package hq

import (
	"testing"

	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	def.Init()
	check.TestMain(m)
}
