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
	"github.com/hellohq/hqservice/ms/auth/srv/mail"
)

// App implements interface Appl.
type appT struct {
	waSvc       wa.IWebauthnSvc
	userSvc     user.IUserSvc
	passcodeSvc passcode.IPasscodeSvc
	emailSvc    email.IEmailSvc
}

// New creates and returns new App.
func NewApp(mailer mail.IMailer, renderer *mail.Renderer, cfg *config.Config, repo dal.IAuthRepo) appT {
	f := false
	waClient, err := webauthn.New(&webauthn.Config{
		RPDisplayName:         cfg.Webauthn.RelyingParty.DisplayName,
		RPID:                  cfg.Webauthn.RelyingParty.Id,
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

	waSvc := wa.NewWebAuthn(cfg, repo, waClient)
	userSvc := user.NewUserSvc(cfg, repo)
	passcodeSvc := passcode.NewPasscodeSvc(mailer, renderer, cfg, repo)
	emailSvc := email.NewEmailSvc(cfg, repo)

	if err != nil {
		panic(fmt.Errorf("failed to create webauthn instance: %w", err))
	}
	return appT{
		waSvc:       waSvc,
		userSvc:     userSvc,
		passcodeSvc: passcodeSvc,
		emailSvc:    emailSvc,
	}
}

func (a appT) GetWebauthnSvc() wa.IWebauthnSvc {
	return a.waSvc
}

func (a appT) GetUserSvc() user.IUserSvc {
	return a.userSvc
}

func (a appT) GetPasscodeSvc() passcode.IPasscodeSvc {
	return a.passcodeSvc
}

func (a appT) GetEmailSvc() email.IEmailSvc {
	return a.emailSvc
}
