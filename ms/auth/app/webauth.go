package app

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app/dom"
)

func NewWebAuthn(cfg config) *webauthn.WebAuthn {
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
	return w
}

func (app *App) WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	user, err := app.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, nil, nil
	}

	// TODO: Impl DAL for WebauthnCredential
	credentials := []ent.WebauthnCredential{}
	// credentials, err := h.persister.GetWebauthnCredentialPersisterWithConnection(connection).GetFromUser(user.ID)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to get webauthn credentials: %w", err)
	// }

	webauthnUser, err := dom.NewWebauthnUser(*user, credentials)
	if err != nil {
		return nil, nil, err
	}

	t := true
	options, sessionData, err := app.wAuthn.BeginRegistration(
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
		return nil, nil, fmt.Errorf("failed to create webauthn creation options: %w", err)
	}

	return options, sessionData, nil
}

// func (app *App) WebauthnFinishRegistration(postBody []byte) *webauthn.Credential {

// 	return credential
// }
