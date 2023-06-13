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
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/spf13/pflag"
)

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet // "hq serve" flags
}

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	fs     FlagSets
	shared *sharedconfig.Shared
	own    = &struct {
		// Below envs is loaded by Doppler
		RpId                     appcfg.String         `env:"AUTH_RP_ID"`
		RpOrigins                appcfg.StringSlice    `env:"AUTH_RP_ORIGINS"`
		IosAssociationSite       appcfg.String         `env:"IOS_SITE_ASSOCIATION"`
		AndroidAssetLinks        appcfg.String         `env:"ANDROID_ASSET_LINKS"`
		OneSignalAppID           appcfg.String         `env:"ONESIGNAL_APP_ID"`
		OneSignalAppKey          appcfg.String         `env:"ONESIGNAL_APP_KEY"`
		FromAddress              appcfg.NotEmptyString `env:"MAIL_FROM_ADDRESS"`
		FromName                 appcfg.String         `env:"MAIL_FROM_NAME"`
		TTL                      appcfg.Int            `env:"PASSCODE_TTL"`
		RequireEmailVerification appcfg.Bool           `env:"REQUIRE_EMAIL_VERIFICATION"`
	}{
		TTL:         appcfg.MustInt("300"),
		FromAddress: appcfg.MustNotEmptyString("noah@hellohq.com"),
	}
)

type Config struct {
	Server                   Server
	Webauthn                 WebauthnSettings
	Passcode                 Passcode
	ServiceName              string
	MaxEmailAddresses        int
	RequireEmailVerification bool
}

func saveStaticFileConfig(fileNameContent map[string]string) error {
	for filename, content := range fileNameContent {
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
	}
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

	// Save required setting file of fido2 for mobile platforms
	filenameContent := map[string]string{
		"apple-app-site-association": own.IosAssociationSite.Value(&err),
		"assetlinks.json":            own.AndroidAssetLinks.Value(&err),
	}
	err = saveStaticFileConfig(filenameContent)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHost, "auth.host", "host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrHostInt, "auth.host-int", "internal host to serve")
	appcfg.AddPFlag(fs.Serve, &shared.AuthAddrPort, "auth.port", "port to serve monolith introspection")

	appcfg.AddPFlag(fs.Serve, &own.RpId, "wa.id", "Webauthn id")
	appcfg.AddPFlag(fs.Serve, &own.RpOrigins, "wa.origins", "Webauthn origin")
	appcfg.AddPFlag(fs.Serve, &own.FromAddress, "from.mail", "sender email address")
	appcfg.AddPFlag(fs.Serve, &own.FromName, "from.name", "sender email name")
	appcfg.AddPFlag(fs.Serve, &own.OneSignalAppID, "onesignal.id", "onesignal app id")
	appcfg.AddPFlag(fs.Serve, &own.OneSignalAppKey, "onesignal.key", "onesignal app key")

	return nil
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *Config, err error) {
	defer cleanup()
	c = &Config{
		Server: Server{
			BindAddr:    netx.NewAddr(shared.AuthAddrHost.Value(&err), shared.AuthAddrPort.Value(&err)),
			BindAddrInt: netx.NewAddr(shared.AuthAddrHostInt.Value(&err), shared.AuthAddrPort.Value(&err)),
			Cors: Cors{
				ExposeHeaders: []string{
					httplimit.HeaderRateLimitLimit,
					httplimit.HeaderRateLimitRemaining,
					httplimit.HeaderRateLimitReset,
					httplimit.HeaderRetryAfter,
				},
			},
		},
		// TODO: Add env vars for below config fields
		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				Id:          own.RpId.Value(&err),
				DisplayName: "Authentication Service",
				Origins:     own.RpOrigins.Value(&err),
			},
			Timeout: 60000,
		},
		Passcode: Passcode{
			Email: Email{
				FromAddress: own.FromAddress.Value(&err),
				FromName:    own.FromName.Value(&err),
			},
			OneSignalAppKey: own.OneSignalAppKey.Value(&err),
			OneSignalAppID:  own.OneSignalAppID.Value(&err),
			TTL:             int32(own.TTL.Value(&err)),
		},
		MaxEmailAddresses:        5,
		RequireEmailVerification: own.RequireEmailVerification.Value(&err),
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, own)
	}

	return c, nil
}

// Cleanup must be called by own Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}
