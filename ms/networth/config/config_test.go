package config

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/nikoksr/doppler-go"
	"github.com/nikoksr/doppler-go/secret"
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
		SaltEdge: &SaltEdge{
			AppId:  "se_app_id",
			Secret: "se_secret",
			PK:     "se_pk",
		},
		Finverse: &Finverse{
			AppId:       "fv_app_id",
			ClientID:    "fv_client_id",
			Secret:      "fv_secret",
			RedirectURI: "fv_redirect_uri",
		},
	}
	testOwn = own
)

func TestMain(m *testing.M) {
	loadDopplerEnvs()
	os.Setenv("HQ_SALTEDGE_APP_ID", "se_app_id")
	os.Setenv("HQ_SALTEDGE_SECRET", "se_secret")
	os.Setenv("HQ_SALTEDGE_PK", "se_pk")

	os.Setenv("HQ_FINVERSE_APP_ID", "fv_app_id")
	os.Setenv("HQ_FINVERSE_CLIENT_ID", "fv_client_id")
	os.Setenv("HQ_FINVERSE_SECRET", "fv_secret")
	os.Setenv("HQ_FINVERSE_REDIRECT_URI", "fv_redirect_uri")
	var err error
	testShared, err = sharedconfig.Get()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	check.TestMain(m)
}

func loadDopplerEnvs() {
	// Set your API key
	doppler.Key = os.Getenv("DOPPLER_TOKEN")

	// List all your secrets
	secrets, _, err := secret.List(context.Background(), &doppler.SecretListOptions{
		Project: "hqservice",
		Config:  "dev",
	})
	if err != nil {
		log.Fatalf("failed to list secrets: %v", err)
	}
	os.Clearenv()
	for name, value := range secrets {
		os.Setenv(name, *value.Raw)
	}
	os.Setenv("HQ_ONESIGNAL_APP_ID", "oneSignalAppID")
	os.Setenv("HQ_ONESIGNAL_APP_KEY", "oneSignalAppKey")
}

func testGetServe(flags ...string) (*Config, error) {
	own = testOwn
	err := Init(testShared, testFlagsets)
	if err != nil {
		return nil, err
	}
	if len(flags) > 0 {
		err = testFlagsets.Serve.Parse(flags)
		if err != nil {
			return nil, err
		}
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
}
