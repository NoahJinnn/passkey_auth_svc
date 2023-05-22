//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

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

// Appl provides application features (use cases) service.
type Appl interface {
	GetWebauthnSvc() wa.IWebauthnSvc
	GetUserSvc() user.IUserSvc
	GetPasscodeSvc() passcode.IPasscodeSvc
	GetEmailSvc() email.IEmailSvc
}

// App implements interface Appl.
type App struct {
	waSvc       wa.IWebauthnSvc
	userSvc     user.IUserSvc
	passcodeSvc passcode.IPasscodeSvc
	emailSvc    email.IEmailSvc
}

// New creates and returns new App.
func New(mailer mail.IMailer, renderer *mail.Renderer, cfg *config.Config, repo dal.IAuthRepo) App {
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
	return App{
		waSvc:       waSvc,
		userSvc:     userSvc,
		passcodeSvc: passcodeSvc,
		emailSvc:    emailSvc,
	}
}

func (a App) GetWebauthnSvc() wa.IWebauthnSvc {
	return a.waSvc
}

func (a App) GetUserSvc() user.IUserSvc {
	return a.userSvc
}

func (a App) GetPasscodeSvc() passcode.IPasscodeSvc {
	return a.passcodeSvc
}

func (a App) GetEmailSvc() email.IEmailSvc {
	return a.emailSvc
}
