package svcs

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app/dom"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type IWebauthnSvc interface {
	WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error)
}

type webauthnSvc struct {
	repo dal.Repo
	wa   *webauthn.WebAuthn
}

func NewWebAuthn(cfg *config.Config, repo dal.Repo) IWebauthnSvc {
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
	return &webauthnSvc{
		repo: repo,
		wa:   wa,
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
	options, sessionData, err := svc.wa.BeginRegistration(
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
