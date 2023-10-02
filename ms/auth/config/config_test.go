package config

import (
	"log"
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
			BindAddr:    netx.NewAddr("0.0.0.0", 17000),
			BindAddrInt: netx.NewAddr("0.0.0.0", 17000),
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
				Id:          "127.0.0.1",
				DisplayName: "Authentication Service",
				Origins:     []string{"http://localhost:17000", "http://localhost:17001"},
			},
			Timeout: 60000,
		},
		Passcode: Passcode{
			Email: Email{
				FromAddress: "noah@hellohq.com",
				FromName:    "HelloHQ Pte. Ltd.",
			},
			OneSignalAppKey: "oneSignalAppKey",
			OneSignalAppID:  "oneSignalAppID",
			TTL:             300,
		},
		MaxEmailAddresses:        5,
		RequireEmailVerification: false,
	}
	testOwn = own
)

func TestMain(m *testing.M) {
	var err error
	testShared, err = sharedconfig.Get()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	check.TestMain(m)
}

func testGetServe(flags ...string) (*Config, error) {
	os.Setenv("HQ_AUTH_ONESIGNAL_APP_ID", "oneSignalAppID")
	os.Setenv("HQ_AUTH_ONESIGNAL_APP_KEY", "oneSignalAppKey")

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
