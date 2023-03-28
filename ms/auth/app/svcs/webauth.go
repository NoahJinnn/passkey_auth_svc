package svcs

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app/dom"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
)

type IWebauthnSvc interface {
	WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error)
	WebauthnFinishRegistration(ctx Ctx, request *protocol.ParsedCredentialCreationData, sessionUserId string) (credentialId string, userId string, err error)
}

type webauthnSvc struct {
	repo *dal.Repo
	wa   *webauthn.WebAuthn
}

var (
	WebauthnOperationRegistration   string = "registration"
	WebauthnOperationAuthentication string = "authentication"
)

func NewWebAuthn(cfg *config.Config, repo *dal.Repo, wa *webauthn.WebAuthn) IWebauthnSvc {
	return &webauthnSvc{
		repo: repo,
		wa:   wa,
	}
}

func (svc *webauthnSvc) WebauthnBeginRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error) {
	webauthnUser, _, err := svc.getWebauthnUser(ctx, userId)

	if webauthnUser == nil {
		// TODO: audit logger
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

	err = svc.repo.GetWebauthnSessionRepo().Create(ctx, *dom.WebauthnSessionDataToModel(sessionData, WebauthnOperationRegistration))
	if err != nil {
		return nil, fmt.Errorf("failed to store creation options session data: %w", err)
	}

	return options, nil
}

func (svc *webauthnSvc) WebauthnFinishRegistration(ctx Ctx, request *protocol.ParsedCredentialCreationData, sessionUserId string) (credentialId string, userId string, err error) {
	if err := svc.repo.WithTx(ctx, func(tx *ent.Tx) error {
		exec := (func(ctx Ctx, client *ent.Client) error {
			sessionDataRepo := svc.repo.GetWebauthnSessionRepo()
			sessionData, err := sessionDataRepo.GetByChallenge(ctx, request.Response.CollectedClientData.Challenge)
			if err != nil {
				return fmt.Errorf("failed to get webauthn registration session data: %w", err)
			}

			if sessionData != nil && sessionData.Operation != WebauthnOperationRegistration {
				sessionData = nil
			}

			if sessionData == nil {
				// TODO: audit logger
				return dto.NewHTTPError(http.StatusBadRequest, "Stored challenge and received challenge do not match").SetInternal(errors.New("sessionData not found"))
			}

			if sessionUserId != sessionData.UserID.String() {
				// TODO: audit logger
				return dto.NewHTTPError(http.StatusBadRequest, "Stored challenge and received challenge do not match").SetInternal(errors.New("userId in webauthn.sessionData does not match user session"))
			}

			webauthnUser, _, err := svc.getWebauthnUser(ctx, sessionData.UserID)
			if err != nil {
				return fmt.Errorf("failed to get user: %w", err)
			}

			if webauthnUser == nil {
				// TODO: audit logger
				return dto.NewHTTPError(http.StatusBadRequest).SetInternal(errors.New("user not found"))
			}

			credential, err := svc.wa.CreateCredential(webauthnUser, *dom.WebauthnSessionDataFromModel(sessionData), request)
			if err != nil {
				errorMessage := "failed to validate attestation"
				errorStatus := http.StatusBadRequest
				// Safari currently (v. 16.2) does not provide a UI in case of a (registration) ceremony
				// being performed with an authenticator NOT protected by e.g. a PIN. While Chromium based browsers do offer
				// a UI guiding through the setup of a PIN, Safari simply performs the ceremony without then setting the UV
				// flag even if it is required. In order to provide an appropriate error message to the frontend/user, we
				// need to return an error response distinguishable from other error cases. We use a dedicated/separate HTTP
				// status code because it seemed a bit more robust than forcing the frontend to check on a matching
				// (sub-)string in the error message in order to properly display the error.
				if err, ok := err.(*protocol.Error); ok && err.Type == protocol.ErrVerification.Type && strings.Contains(err.DevInfo, "User verification") {
					errorMessage = fmt.Sprintf("%s: %s: %s", errorMessage, err.Details, err.DevInfo)
					errorStatus = http.StatusUnprocessableEntity
				}
				// TODO: audit logger

				return dto.NewHTTPError(errorStatus, errorMessage).SetInternal(err)
			}

			backupEligible := request.Response.AttestationObject.AuthData.Flags.HasBackupEligible()
			backupState := request.Response.AttestationObject.AuthData.Flags.HasBackupState()
			model := dom.WebauthnCredentialToModel(credential, sessionData.UserID, backupEligible, backupState)
			err = svc.repo.GetWebauthnCredentialRepo().Create(ctx, *model)
			if err != nil {
				return fmt.Errorf("failed to store webauthn credential: %w", err)
			}

			err = sessionDataRepo.Delete(ctx, *sessionData)
			if err != nil {
				return fmt.Errorf("failed to delete attestation session data: %w", err)
			}

			// TODO: audit logger
			return nil
		})
		return exec(ctx, tx.Client())
	}); err != nil {
		return credentialId, userId, err
	}
	return credentialId, userId, nil
}

func (svc *webauthnSvc) getWebauthnUser(ctx Ctx, userId uuid.UUID) (*dom.WebauthnUser, *ent.User, error) {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, nil, nil
	}

	credentials, err := svc.repo.GetWebauthnCredentialRepo().GetFromUser(ctx, user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get webauthn credentials: %w", err)
	}

	webauthnUser, err := dom.NewWebauthnUser(ctx, *user, credentials)
	if err != nil {
		return nil, nil, err
	}

	return webauthnUser, user, nil
}
