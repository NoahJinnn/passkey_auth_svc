package test

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/ms/auth/app/email"
	"github.com/hellohq/hqservice/ms/auth/app/passcode"
	"github.com/hellohq/hqservice/ms/auth/app/user"
	"github.com/hellohq/hqservice/ms/auth/app/wa"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

// App implements interface Appl.
type appT struct {
	cfg  *config.Config
	repo dal.IAuthRepo
	wa   *webauthn.WebAuthn
}

// New creates and returns new App.
func NewApp(cfg *config.Config, repo dal.IAuthRepo) appT {
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
	return appT{
		cfg:  cfg,
		repo: repo,
		wa:   wa,
	}
}

func (a appT) GetWebauthnSvc() wa.IWebauthnSvc {
	return wa.NewWebAuthn(a.cfg, a.repo, a.wa)
}

func (a appT) GetUserSvc() user.IUserSvc {
	return user.NewUserSvc(a.cfg, a.repo)
}

func (a appT) GetPasscodeSvc() passcode.IPasscodeSvc {
	return passcode.NewPasscodeSvc(a.cfg, a.repo)
}

func (a appT) GetEmailSvc() email.IEmailSvc {
	return email.NewEmailSvc(a.repo)
}
