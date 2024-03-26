// Package app provides business logic.
package app

import (
	"fmt"

	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app/email"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app/passcode"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app/user"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app/wa"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app/ws"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/mail"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetWebauthnSvc() *wa.WebauthnSvc
	GetUserSvc() *user.UserSvc
	GetPasscodeSvc() *passcode.PasscodeSvc
	GetEmailSvc() *email.EmailSvc
	GetChangesetSvc() *ws.ChangesetSvc
}

// App implements interface Appl.
type App struct {
	waSvc       *wa.WebauthnSvc
	userSvc     *user.UserSvc
	passcodeSvc *passcode.PasscodeSvc
	emailSvc    *email.EmailSvc
	csSvc       *ws.ChangesetSvc
}

// New creates and returns new App.
func New(mailer mail.IMailer, renderer *mail.Renderer, cfg *config.Config, repo dal.IAuthRepo) *App {
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
	csSvc := ws.NewChangesetSvc(repo)
	if err != nil {
		panic(fmt.Errorf("failed to create webauthn instance: %w", err))
	}
	return &App{
		waSvc:       waSvc,
		userSvc:     userSvc,
		passcodeSvc: passcodeSvc,
		emailSvc:    emailSvc,
		csSvc:       csSvc,
	}
}

func (a App) GetWebauthnSvc() *wa.WebauthnSvc {
	return a.waSvc
}

func (a App) GetUserSvc() *user.UserSvc {
	return a.userSvc
}

func (a App) GetPasscodeSvc() *passcode.PasscodeSvc {
	return a.passcodeSvc
}

func (a App) GetEmailSvc() *email.EmailSvc {
	return a.emailSvc
}

func (a App) GetChangesetSvc() *ws.ChangesetSvc {
	return a.csSvc
}
