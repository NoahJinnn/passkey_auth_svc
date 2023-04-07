// Package config provides configurations for subcommands.
//
// It consists of both configuration values shared by own
// microservices and values specific to this microservice.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.

package config

import (
	"fmt"
	"os"

	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/powerman/pqx"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet
}

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	ServiceName = "auth"
	fs          FlagSets
	shared      *sharedconfig.Shared
	own         = &struct {
		// Below envs is loaded by Doppler
		PostgresUser       appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_LOGIN"`
		PostgresPass       appcfg.NotEmptyString `env:"AUTH_POSTGRES_AUTH_PASS"`
		PostgresAddrHost   appcfg.NotEmptyString `env:"AUTH_POSTGRES_ADDR_HOST"`
		PostgresAddrPort   appcfg.Port           `env:"AUTH_POSTGRES_ADDR_PORT"`
		PostgresDBName     appcfg.NotEmptyString `env:"AUTH_POSTGRES_DB_NAME"`
		Secrets            appcfg.NotEmptyString `env:"AUTH_SECRETS"`
		RpId               appcfg.NotEmptyString `env:"AUTH_RP_ID"`
		RpOrigin           appcfg.NotEmptyString `env:"AUTH_RP_ORIGIN"`
		RpOrigins          appcfg.StringSlice    `env:"AUTH_RP_ORIGINS"`
		IosAssociationSite appcfg.String         `env:"IOS_SITE_ASSOCIATION"`
		AndroidAssetLinks  appcfg.String         `env:"ANDROID_ASSET_LINKS"`
	}{
		PostgresUser:     appcfg.MustNotEmptyString(ServiceName),
		PostgresAddrPort: appcfg.MustPort("5432"),
		PostgresAddrHost: appcfg.MustNotEmptyString("localhost"),
		PostgresDBName:   appcfg.MustNotEmptyString("postgres"),
		Secrets:          appcfg.MustNotEmptyString("needstobeatleast16"),
		RpId:             appcfg.MustNotEmptyString("localhost"),
		RpOrigin:         appcfg.MustNotEmptyString("localhost:17000"),
	}
)

type Config struct {
	Server      Server
	Webauthn    WebauthnSettings
	Session     Session
	Secrets     Secrets
	Emails      Emails
	Passcode    Passcode
	ServiceName string
	Postgres    *PostgresConfig
	Plaid       *PlaidConfig
}

// Save apple association site file to static folder
func saveStaticFileConfig(content string, filename string) error {

	_, err := os.Stat("static")
	if err != nil {
		fmt.Println("Static dir does not exist", err)
		if err := os.Mkdir("static", os.ModePerm); err != nil {
			return fmt.Errorf("create static dir failed: %w", err)
		}
	}

	destination, err := os.Create(fmt.Sprintf("static/%s", filename))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer destination.Close()

	_, err = fmt.Fprintf(destination, "%s", content)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	fmt.Printf("File %s saved successfully\n", filename)
	return nil
}

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(sharedCfg *sharedconfig.Shared, flagsets FlagSets) error {

	shared, fs = sharedCfg, flagsets
	fromEnv := appcfg.NewFromEnv(sharedconfig.EnvPrefix)
	err := appcfg.ProvideStruct(own, fromEnv)
	if err != nil {
		return err
	}

	err = saveStaticFileConfig(own.IosAssociationSite.Value(&err), "apple-app-site-association")
	if err != nil {
		return err
	}

	err = saveStaticFileConfig(own.AndroidAssetLinks.Value(&err), "assetlinks.json")
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHost, "host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHostInt, "host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPort, "port", "port to serve monolith introspection")
	appcfg.AddPFlag(fs.Serve, &own.PostgresAddrHost, "postgres.host", "host to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &own.PostgresAddrPort, "postgres.port", "port to connect to PostgreSQL")
	appcfg.AddPFlag(fs.Serve, &own.PostgresDBName, "postgres.dbname", "PostgreSQL database name")
	appcfg.AddPFlag(fs.Serve, &own.PostgresUser, "postgres.user", "PostgreSQL username")
	appcfg.AddPFlag(fs.Serve, &own.PostgresPass, "postgres.pass", "PostgreSQL password")
	appcfg.AddPFlag(fs.Serve, &own.RpId, "wa.id", "Webauthn id")
	appcfg.AddPFlag(fs.Serve, &own.RpOrigin, "wa.origin", "Webauthn origin")
	appcfg.AddPFlag(fs.Serve, &own.RpOrigins, "wa.origins", "Webauthn origin")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()
	c = &Config{
		Server: Server{
			BindAddr: netx.NewAddr(shared.AuthAddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
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
			Host:   own.PostgresAddrHost.Value(&err),
			Port:   own.PostgresAddrPort.Value(&err),
			DBName: own.PostgresDBName.Value(&err),
			User:   own.PostgresUser.Value(&err),
			Pass:   own.PostgresPass.Value(&err),
		}),
		// TODO: Add env vars for below config fields
		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				Id:          own.RpId.Value(&err),
				DisplayName: "Authentication Service",
				Origins:     own.RpOrigins.Value(&err),
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
			Keys: []string{own.Secrets.Value(&err)},
		},
		Emails: Emails{
			RequireVerification: false,
			MaxNumOfAddresses:   50,
		},
		ServiceName: ServiceName,
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own, shared)
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}

func (c *Config) Validate() error {
	err := c.Webauthn.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate webauthn settings: %w", err)
	}
	err = c.Passcode.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate passcode settings: %w", err)
	}
	err = c.Session.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate session settings: %w", err)
	}
	err = c.Secrets.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate secrets settings: %w", err)
	}

	return nil
}
