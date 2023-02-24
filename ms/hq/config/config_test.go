package config

import (
	"os"
	"testing"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/check"
	"github.com/powerman/go-monolith-example/pkg/def"
	"github.com/powerman/pqx"
)

func Test(t *testing.T) {
	want := &Config{
		Postgres: NewPostgresConfig(pqx.Config{
			Host:        "postgres",
			Port:        5432,
			DBName:      "postgres",
			User:        "auth",
			Pass:        "authpass",
			SSLRootCert: "ca.crt",
		}),
		BindAddr:        netx.NewAddr(def.Hostname, sharedconfig.MonoPort),
		BindMetricsAddr: netx.NewAddr(def.Hostname, sharedconfig.MonoPort),
	}

	t.Run("required", func(tt *testing.T) {
		t := check.T(tt)
		require(t, "PostgresPass")
		os.Setenv("MONO_HQ_POSTGRES_AUTH_PASS", "authpass")
	})
	t.Run("default", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe()
		t.Nil(err)
		t.DeepEqual(c, want)
	})
	t.Run("constraint", func(tt *testing.T) {
		t := check.T(tt)
		constraint(t, "MONO_HQ_POSTGRES_AUTH_LOGIN", "", `^PostgresUser .* empty`)
		constraint(t, "MONO_HQ_POSTGRES_AUTH_PASS", "", `^PostgresPass .* empty`)

	})
	t.Run("env", func(tt *testing.T) {
		t := check.T(tt)
		os.Setenv("MONO_HQ_POSTGRES_AUTH_LOGIN", "auth3")
		os.Setenv("MONO_HQ_POSTGRES_AUTH_PASS", "authpass3")

		c, err := testGetServe()
		t.Nil(err)
		want.Postgres.User = "auth3"
		want.Postgres.Pass = "authpass3"

		t.DeepEqual(c, want)
	})
	t.Run("flag", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe(
			"--postgres.host=localhost4",
			"--postgres.port=4200",
			"--postgres.dbname=postgres4",
			"--hq.postgres.user=auth4",
			"--hq.postgres.pass=authpass4",
			"--host=host4",
			"--host-int=hostint4",
			"--hq.host=authhost4",
			"--hq.port=4102",
			"--hq.port-int=4103",
			"--hq.metrics.port=4101",
		)
		t.Nil(err)
		want.Postgres.Host = "localhost4"
		want.Postgres.Port = 4200
		want.Postgres.DBName = "postgres4"
		want.Postgres.User = "auth4"
		want.Postgres.Pass = "authpass4"
		want.BindAddr = netx.NewAddr("host4", 4102)
		want.BindMetricsAddr = netx.NewAddr("hostint4", 4101)
		t.DeepEqual(c, want)
	})
	t.Run("cleanup", func(tt *testing.T) {
		t := check.T(tt)
		t.Panic(func() { GetServe() })
	})
}
