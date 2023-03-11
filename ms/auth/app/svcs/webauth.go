package svcs

import (
	"context"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app/dom"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type Ctx = context.Context

type IWebauthnSvc interface {
	WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error)
}

type webauthnSvc struct {
	repo dal.Repo
	w    *webauthn.WebAuthn
}

func NewWebAuthn(repo dal.Repo) IWebauthnSvc {
	var w *webauthn.WebAuthn
	var err error
	wconfig := &webauthn.Config{
		RPDisplayName: "Go Webauthn",                               // Display Name for your site
		RPID:          "go-webauthn.local",                         // Generally the FQDN for your site
		RPOrigins:     []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	}

	if w, err = webauthn.New(wconfig); err != nil {
		fmt.Println(err)
	}
	return &webauthnSvc{
		repo: repo,
		w:    w,
	}
}

func (svc *webauthnSvc) WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error) {
	user, err := svc.repo.GetUserRepo().GetUserById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	// TODO: Impl DAL for WebauthnCredential
	credentials := []ent.WebauthnCredential{}
	// credentials, err := h.persister.GetWebauthnCredentialPersisterWithConnection(connection).GetFromUser(user.ID)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to get webauthn credentials: %w", err)
	// }

	webauthnUser, err := dom.NewWebauthnUser(*user, credentials)
	if err != nil {
		return nil, err
	}

	if webauthnUser == nil {
		return nil, fmt.Errorf("failed to get webauthnuser: %w", err)
	}

	t := true
	options, sessionData, err := svc.w.BeginRegistration(
		webauthnUser,
		webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
			RequireResidentKey: &t,
			ResidentKey:        protocol.ResidentKeyRequirementRequired,
			UserVerification:   protocol.VerificationRequired,
		}),
		webauthn.WithConveyancePreference(protocol.PreferNoAttestation),
		// don't set the excludeCredentials list, so an already registered device can be re-registered
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create webauthn creation options: %w", err)
	}

	// TODO: Impl DAL
	// err = h.persister.GetWebauthnSessionDataPersister().Create(*intern.WebauthnSessionDataToModel(sessionData, models.WebauthnOperationRegistration))
	// if err != nil {
	// 	return fmt.Errorf("failed to store creation options session data: %w", err)
	// }

	return options, nil
}

// func (app *App) WebauthnFinishRegistration(postBody []byte) *webauthn.Credential {

// 	return credential
// }
