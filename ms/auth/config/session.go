package config

import (
	"errors"
	"time"
)

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
