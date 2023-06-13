package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/httpx"
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
			BindAddr:    netx.NewAddr("localhost", 17000),
			BindAddrInt: netx.NewAddr("127.0.0.1", 17000),
			Cors: Cors{
				ExposeHeaders: []string{
					httplimit.HeaderRateLimitLimit,
					httplimit.HeaderRateLimitRemaining,
					httplimit.HeaderRateLimitReset,
					httplimit.HeaderRetryAfter,
				},
			},
		},
		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				Id:          "localhost",
				DisplayName: "Authentication Service",
				Origins:     []string{"http://localhost:17000", "http://localhost:17001"},
			},
			Timeout: 60000,
		},
		Passcode: Passcode{
			Email: Email{
				FromAddress: "test@gmail.com",
				FromName:    "Test Mail",
			},
			OneSignalAppKey: "oneSignalAppKey",
			OneSignalAppID:  "oneSignalAppID",
			TTL:             300,
		},
		MaxEmailAddresses:        5,
		RequireEmailVerification: true,
	}
	testOwn = own
)

func TestMain(m *testing.M) {
	loadDopplerEnvs()
	os.Clearenv()
	// Shared env
	os.Setenv("HQ_AUTH_ADDR_HOST", "localhost")
	os.Setenv("HQ_AUTH_ADDR_HOST_INT", "127.0.0.1")
	os.Setenv("HQ_AUTH_ADDR_PORT", "17000")
	os.Setenv("HQ_POSTGRES_AUTH_PASS", "authpass")
	os.Setenv("HQ_JWT_LIFESPAN", "1h")
	// Auth env
	os.Setenv("HQ_AUTH_RP_ORIGINS", "http://localhost:17000,http://localhost:17001")
	os.Setenv("HQ_ONESIGNAL_APP_ID", "oneSignalAppID")
	os.Setenv("HQ_ONESIGNAL_APP_KEY", "oneSignalAppKey")
	os.Setenv("HQ_MAIL_FROM_ADDRESS", "test@gmail.com")
	os.Setenv("HQ_MAIL_FROM_NAME", "Test Mail")
	os.Setenv("HQ_REQUIRE_EMAIL_VERIFICATION", "true")
	os.Setenv("HQ_AUTH_RP_ID", "localhost")

	var err error
	testShared, err = sharedconfig.Get()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	check.TestMain(m)
}

func loadDopplerEnvs() {
	token := fmt.Sprintf("Bearer %s", os.Getenv("DOPPLER_TOKEN"))
	fmt.Println(token)
	req := httpx.NewReq("https://api.doppler.com/v3/configs/config/secrets", map[string]string{
		"Content-Type":  "application/json",
		"accept":        "application/json",
		"accepts":       "application/json",
		"authorization": token,
	}, map[string]string{
		"project":                 "hqservice",
		"config":                  "dev",
		"include_dynamic_secrets": "false",
		"include_managed_secrets": "true",
	})

	resp, err := req.
		InitReq(context.Background(), "GET", "", nil).
		WithDefaultOpts().
		Send()
	fmt.Println(string(resp.Body()))
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
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
			"--auth.host=authhost4",
			"--auth.host-int=authhostint4",
			"--auth.port=4102",

			"--wa.id=flagrpid",
			"--wa.origins=localhost:8081,localhost:8082",
			"--from.mail=testflag@gmail.com",
			"--from.name=Test Mail Flag",
			"--onesignal.id=oneSignalIdFlag",
			"--onesignal.key=oneSignalKeyFlag",
		)
		t.Nil(err)

		want.Server.BindAddr = netx.NewAddr("authhost4", 4102)
		want.Server.BindAddrInt = netx.NewAddr("authhostint4", 4102)
		want.Webauthn.RelyingParty.Id = "flagrpid"
		want.Webauthn.RelyingParty.Origins = []string{"localhost:8081", "localhost:8082"}
		want.Passcode.Email.FromAddress = "testflag@gmail.com"
		want.Passcode.Email.FromName = "Test Mail Flag"
		want.Passcode.OneSignalAppID = "oneSignalIdFlag"
		want.Passcode.OneSignalAppKey = "oneSignalKeyFlag"
		t.DeepEqual(c, want)
	})

	t.Run("cleanup", func(tt *testing.T) {
		err := os.RemoveAll("static")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Directory removed successfully")
		}
	})
}
