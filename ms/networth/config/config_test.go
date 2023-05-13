package config

import (
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/check"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

var (
	testShared   *sharedconfig.Shared
	testFlagsets = FlagSets{
		Serve: pflag.NewFlagSet("", 0),
	}
	want = &Config{
		Server: Server{
			BindAddr:    netx.NewAddr("localhost", 17002),
			BindAddrInt: netx.NewAddr("127.0.0.1", 17002),
			Cors: Cors{
				ExposeHeaders: []string{
					httplimit.HeaderRateLimitLimit,
					httplimit.HeaderRateLimitRemaining,
					httplimit.HeaderRateLimitReset,
					httplimit.HeaderRetryAfter,
				},
			},
		},
		SaltEdgeConfig: &SaltEdgeConfig{
			AppId:  "test",
			Secret: "test",
			PK:     "test",
		},
	}
	testOwn = own
)

func TestMain(m *testing.M) {
	os.Clearenv()
	// Shared env
	os.Setenv("HQ_NETWORTH_ADDR_HOST", "localhost")
	os.Setenv("HQ_NETWORTH_ADDR_HOST_INT", "127.0.0.1")
	os.Setenv("HQ_NETWORTH_ADDR_PORT", "17002")
	os.Setenv("HQ_POSTGRES_AUTH_PASS", "authpass")
	// Networth env
	os.Setenv("HQ_SALTEDGE_APP_ID", "test")
	os.Setenv("HQ_SALTEDGE_SECRET", "test")
	os.Setenv("HQ_SALTEDGE_PK", "test")

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

func Test(t *testing.T) {
	t.Run("env", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe()
		t.Nil(err)
		t.DeepEqual(c, want)

	})

	t.Run("flag", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe(
			"--networth.host=networthhost4",
			"--networth.host-int=networthhostint4",
			"--networth.port=4102",
		)
		t.Nil(err)

		want.Server.BindAddr = netx.NewAddr("networthhost4", 4102)
		want.Server.BindAddrInt = netx.NewAddr("networthhostint4", 4102)

		t.DeepEqual(c, want)
	})

	t.Run("cleanup", func(tt *testing.T) {
		t := check.T(tt)
		t.Panic(func() { GetServe() })
	})
}
