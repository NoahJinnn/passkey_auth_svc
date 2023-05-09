//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetWebauthnSvc() svcs.IWebauthnSvc
	GetUserSvc() svcs.IUserSvc
	GetPasscodeSvc() svcs.IPasscodeSvc
	GetEmailSvc() svcs.IEmailSvc
}

// App implements interface Appl.
type App struct {
	cfg  *config.Config
	repo *dal.AuthRepo
	wa   *webauthn.WebAuthn
}

// New creates and returns new App.
func New(cfg *config.Config, repo *dal.AuthRepo) App {
	f := false
	wa, err := webauthn.New(&webauthn.Config{
		RPDisplayName:         cfg.Webauthn.RelyingParty.DisplayName,
		RPID:                  cfg.Webauthn.RelyingParty.Id,
		RPOrigin:              cfg.Webauthn.RelyingParty.Origin,
		RPOrigins:             cfg.Webauthn.RelyingParty.Origins,
		AttestationPreference: protocol.PreferNoAttestation,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			RequireResidentKey: &f,
			ResidentKey:        protocol.ResidentKeyRequirementDiscouraged,
			UserVerification:   protocol.VerificationRequired,
		},
		Timeout: cfg.Webauthn.Timeout,
		Debug:   false,
	})

	if err != nil {
		panic(fmt.Errorf("failed to create webauthn instance: %w", err))
	}
	return App{
		cfg:  cfg,
		repo: repo,
		wa:   wa,
	}
}

func (a App) GetWebauthnSvc() svcs.IWebauthnSvc {
	return svcs.NewWebAuthn(a.cfg, a.repo, a.wa)
}

func (a App) GetUserSvc() svcs.IUserSvc {
	return svcs.NewUserSvc(a.cfg, a.repo)
}

func (a App) GetPasscodeSvc() svcs.IPasscodeSvc {
	return svcs.NewPasscodeSvc(a.cfg, a.repo)
}

func (a App) GetEmailSvc() svcs.IEmailSvc {
	return svcs.NewEmailSvc(a.repo)
}
