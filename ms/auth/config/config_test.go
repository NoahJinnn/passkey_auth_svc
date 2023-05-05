package config

import (
	"log"
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/check"
	"github.com/powerman/pqx"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

var (
	testShared   *sharedConfig.Shared
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
		Postgres: NewPostgresConfig(pqx.Config{
			Host:   "localhost",
			Port:   5432,
			DBName: "postgres",
			User:   "auth",
			Pass:   "authpass",
		}),
		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				Id:          "localhost",
				DisplayName: "Authentication Service",
				Origins:     []string{"http://localhost:17000", "http://localhost:17001"},
			},
			Timeout: 60000,
		},
		Session: Session{
			Lifespan: "1h",
			Cookie: Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: true,
		},
		Secrets: Secrets{
			Keys: []string{"needsToBeAtLeast16"},
		},
		ServiceName: ServiceName,
		Passcode: Passcode{
			Email: Email{
				FromAddress: "test@gmail.com",
				FromName:    "Test Mail",
			},
			OneSignalAppKey: "oneSignalAppKey",
			OneSignalAppID:  "oneSignalAppID",
			TTL:             300,
		},
	}
	testOwn = own
)

func TestMain(m *testing.M) {
	os.Clearenv()
	// Shared env
	os.Setenv("HQ_AUTH_ADDR_HOST", "localhost")
	os.Setenv("HQ_AUTH_ADDR_HOST_INT", "127.0.0.1")
	os.Setenv("HQ_AUTH_ADDR_PORT", "17000")
	// Auth env
	os.Setenv("HQ_AUTH_POSTGRES_AUTH_PASS", "authpass")
	os.Setenv("HQ_AUTH_RP_ORIGINS", "http://localhost:17000,http://localhost:17001")
	os.Setenv("HQ_ONESIGNAL_APP_ID", "oneSignalAppID")
	os.Setenv("HQ_ONESIGNAL_APP_KEY", "oneSignalAppKey")
	os.Setenv("HQ_MAIL_FROM_ADDRESS", "test@gmail.com")
	os.Setenv("HQ_MAIL_FROM_NAME", "Test Mail")

	testShared, _ = sharedConfig.Get()
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
			"--postgres.host=localhost4",
			"--postgres.port=4200",
			"--postgres.dbname=postgres4",
			"--postgres.user=auth4",
			"--postgres.pass=authpass4",
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

		want.Postgres.Host = "localhost4"
		want.Postgres.Port = 4200
		want.Postgres.DBName = "postgres4"
		want.Postgres.User = "auth4"
		want.Postgres.Pass = "authpass4"
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
		t := check.T(tt)
		t.Panic(func() { GetServe() })
		err := os.RemoveAll("static")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Directory removed successfully")
		}
	})
}
