package config

import (
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/check"
	_ "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/pflag"
)

var (
	testShared   *sharedconfig.Shared
	testOwn      = own
	testFlagsets = FlagSets{
		Serve: pflag.NewFlagSet("", 0),
	}
)

func TestMain(m *testing.M) {
	def.Init()
	os.Clearenv()
	os.Setenv("MONO_X_POSTGRES_ADDR_HOST", "postgres")
	testShared, _ = sharedconfig.Get()
	check.TestMain(m)
}

func testGetServe(flags ...string) (*Config, error) {
	own = testOwn
	err := Init(testShared, testFlagsets)
	if err != nil {
		return nil, err
	}
	if len(flags) > 0 {
		testFlagsets.Serve.Parse(flags)
	}
	return GetServe()
}

// Require helps testing for missing env var (required to set
// configuration value which don't have default value).
func require(t *check.C, field string) {
	t.Helper()
	c, err := testGetServe()
	t.Match(err, `^`+field+` .* required`)
	t.Nil(c)
}

// Constraint helps testing for invalid env var value.
func constraint(t *check.C, name, val, match string) { //nolint:unparam // val always receives "".
	t.Helper()
	old, ok := os.LookupEnv(name)

	t.Nil(os.Setenv(name, val))
	c, err := testGetServe()
	t.Match(err, match)
	t.Nil(c)

	if ok {
		os.Setenv(name, old)
	} else {
		os.Unsetenv(name)
	}
}
