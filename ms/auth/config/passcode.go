package config

import (
	"errors"
	"fmt"
	"strings"
)

type Email struct {
	FromAddress string `yaml:"from_address" json:"from_address" koanf:"from_address" split_words:"true"`
	FromName    string `yaml:"from_name" json:"from_name" koanf:"from_name" split_words:"true"`
}

func (e *Email) Validate() error {
	if len(strings.TrimSpace(e.FromAddress)) == 0 {
		return errors.New("from_address must not be empty")
	}
	return nil
}

// SMTP Server Settings for sending passcodes
type SMTP struct {
	OneSignalAppKey string
	OneSignalAppID  string
}

type Passcode struct {
	Email Email `yaml:"email" json:"email" koanf:"email"`
	Smtp  SMTP  `yaml:"smtp" json:"smtp" koanf:"smtp"`
	TTL   int32 `yaml:"ttl" json:"ttl" koanf:"ttl"`
}

func (p *Passcode) Validate() error {
	err := p.Email.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate email settings: %w", err)
	}
	return nil
}
