package dom_test

import (
	"testing"

	"github.com/powerman/check"
	"github.com/powerman/go-monolith-example/pkg/def"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	//nolint:errcheck
	def.Init()
	check.TestMain(m)
}
