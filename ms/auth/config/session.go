package config

import (
	"errors"
	"fmt"
	"time"
)

func (c *Config) Validate() error {

	err := c.Webauthn.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate webauthn settings: %w", err)
	}

	err = c.Session.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate session settings: %w", err)
	}

	return nil
}

// WebauthnSettings defines the settings for the webauthn authentication mechanism
type WebauthnSettings struct {
	RelyingParty RelyingParty `yaml:"relying_party" json:"relying_party" koanf:"relying_party" split_words:"true"`
	Timeout      int          `yaml:"timeout" json:"timeout" koanf:"timeout"`
}

// Validate does not need to validate the config, because the library does this already
func (r *WebauthnSettings) Validate() error {
	return nil
}

// RelyingParty webauthn settings for your application.
type RelyingParty struct {
	Id          string `yaml:"id" json:"id" koanf:"id"`
	DisplayName string `yaml:"display_name" json:"display_name" koanf:"display_name" split_words:"true"`
	Icon        string `yaml:"icon" json:"icon" koanf:"icon"`
	// Deprecated: Use Origins instead
	Origin  string   `yaml:"origin" json:"origin" koanf:"origin"`
	Origins []string `yaml:"origins" json:"origins" koanf:"origins"`
}

type Session struct {
	EnableAuthTokenHeader bool   `yaml:"enable_auth_token_header" json:"enable_auth_token_header" koanf:"enable_auth_token_header" split_words:"true"`
	Lifespan              string `yaml:"lifespan" json:"lifespan" koanf:"lifespan"`
	Cookie                Cookie `yaml:"cookie" json:"cookie" koanf:"cookie"`
}

type Cookie struct {
	Domain   string `yaml:"domain" json:"domain" koanf:"domain"`
	HttpOnly bool   `yaml:"http_only" json:"http_only" koanf:"http_only" split_words:"true"`
	SameSite string `yaml:"same_site" json:"same_site" koanf:"same_site" split_words:"true"`
	Secure   bool   `yaml:"secure" json:"secure" koanf:"secure"`
}

func (s *Session) Validate() error {
	_, err := time.ParseDuration(s.Lifespan)
	if err != nil {
		return errors.New("failed to parse lifespan")
	}
	return nil
}

type Secrets struct {
	// Keys secrets are used to en- and decrypt the JWKs which get used to sign the JWTs.
	// For every key a JWK is generated, encrypted with the key and persisted in the database.
	//
	// You can use this list for key rotation: add a new key to the beginning of the list and the corresponding
	// JWK will then be used for signing JWTs. All tokens signed with the previous JWK(s) will still
	// be valid until they expire. Removing a key from the list does not remove the corresponding
	// database record. If you remove a key, you also have to remove the database record, otherwise
	// application startup will fail.
	//
	// Each key must be at least 16 characters long.
	Keys []string `yaml:"keys" json:"keys" koanf:"keys"`
}

func (s *Secrets) Validate() error {
	if len(s.Keys) == 0 {
		return errors.New("at least one key must be defined")
	}
	return nil
}
