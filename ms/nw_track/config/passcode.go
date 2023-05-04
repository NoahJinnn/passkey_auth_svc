package config

import (
	"errors"
	"fmt"
	"strings"
)

type Email struct {
	FromAddress string `split_words:"true"`
	FromName    string `split_words:"true"`
}

func (e *Email) Validate() error {
	if len(strings.TrimSpace(e.FromAddress)) == 0 {
		return errors.New("from_address must not be empty")
	}
	return nil
}

type Passcode struct {
	OneSignalAppKey string
	OneSignalAppID  string
	Email           Email
	TTL             int32
}

func (p *Passcode) Validate() error {
	err := p.Email.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate email settings: %w", err)
	}
	return nil
}
