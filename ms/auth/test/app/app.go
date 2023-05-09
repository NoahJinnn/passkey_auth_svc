package test

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

// App implements interface Appl.
type app struct {
	cfg  *config.Config
	repo dal.IAuthRepo
	wa   *webauthn.WebAuthn
}

// New creates and returns new App.
func NewApp(cfg *config.Config, repo dal.IAuthRepo) app {
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
	return app{
		cfg:  cfg,
		repo: repo,
		wa:   wa,
	}
}

func (a app) GetWebauthnSvc() svcs.IWebauthnSvc {
	return svcs.NewWebAuthn(a.cfg, a.repo, a.wa)
}

func (a app) GetUserSvc() svcs.IUserSvc {
	return svcs.NewUserSvc(a.cfg, a.repo)
}

func (a app) GetPasscodeSvc() svcs.IPasscodeSvc {
	return svcs.NewPasscodeSvc(a.cfg, a.repo)
}

func (a app) GetEmailSvc() svcs.IEmailSvc {
	return svcs.NewEmailSvc(a.repo)
}
