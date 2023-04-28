package config

import (
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/powerman/check"
	_ "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/pflag"
)

var (
	testShared   *sharedconfig.Shared
	testFlagsets = FlagSets{
		Serve: pflag.NewFlagSet("", 0),
	}
)

func TestMain(m *testing.M) {
	os.Clearenv()
	os.Setenv("HQ_AUTH_POSTGRES_ADDR_HOST", "postgres")
	os.Setenv("AUTH_ADDR_HOST", "localhost")
	os.Setenv("AUTH_ADDR_HOST_INT", "127.0.0.1")
	os.Setenv("AUTH_ADDR_PORT", "17000")
	testShared, _ = sharedconfig.Get()
	check.TestMain(m)
}

func testGetServe(flags ...string) (*Config, error) {
	err := Init(testShared, testFlagsets)
	if err != nil {
		return nil, err
	}
	if len(flags) > 0 {
		testFlagsets.Serve.Parse(flags)
	}
	return GetServe()
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
