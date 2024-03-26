package wa

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/errorhandler"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
)

type Ctx = context.Context

type WebauthnSvc struct {
	repo     dal.IAuthRepo
	waClient *webauthn.WebAuthn
}

var (
	WebauthnOperationRegistration   string = "registration"
	WebauthnOperationAuthentication string = "authentication"
)

func NewWebAuthn(cfg *config.Config, repo dal.IAuthRepo, waClient *webauthn.WebAuthn) *WebauthnSvc {
	return &WebauthnSvc{
		repo:     repo,
		waClient: waClient,
	}
}

func (svc *WebauthnSvc) InitRegistration(ctx Ctx, userId uuid.UUID) (*protocol.CredentialCreation, error) {
	webauthnUser, _, err := svc.getWebauthnUser(ctx, userId)
	if webauthnUser == nil {
		// TODO: audit logger
		return nil, fmt.Errorf("failed to get webauthnuser: %w", err)
	}

	t := true
	options, sessionData, err := svc.waClient.BeginRegistration(
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

	err = svc.repo.GetWebauthnSessionRepo().Create(ctx, *WebauthnSessionDataToModel(sessionData, WebauthnOperationRegistration))
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return options, nil
}

func (svc *WebauthnSvc) FinishRegistration(ctx Ctx, request *protocol.ParsedCredentialCreationData, sessionUserId string) (credentialId string, userId string, err error) {
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		sessionDataRepo := svc.repo.GetWebauthnSessionRepo()
		sessionData, err := sessionDataRepo.GetByChallenge(ctx, request.Response.CollectedClientData.Challenge)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if sessionData != nil && sessionData.Operation != WebauthnOperationRegistration {
			sessionData = nil
		}

		if sessionData == nil {
			// TODO: audit logger
			return errorhandler.NewHTTPError(http.StatusBadRequest, "Stored challenge and received challenge do not match").SetInternal(errors.New("sessionData not found"))
		}

		if sessionUserId != sessionData.UserID.String() {
			// TODO: audit logger
			return errorhandler.NewHTTPError(http.StatusBadRequest, "Stored challenge and received challenge do not match").SetInternal(errors.New("userId in webauthn.sessionData does not match user session"))
		}

		webauthnUser, _, err := svc.getWebauthnUser(ctx, sessionData.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		if webauthnUser == nil {
			// TODO: audit logger
			return errorhandler.NewHTTPError(http.StatusBadRequest).SetInternal(errors.New("user not found"))
		}

		credential, err := svc.waClient.CreateCredential(webauthnUser, *WebauthnSessionDataFromModel(sessionData), request)
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

			return errorhandler.NewHTTPError(errorStatus, errorMessage).SetInternal(err)
		}

		backupEligible := request.Response.AttestationObject.AuthData.Flags.HasBackupEligible()
		backupState := request.Response.AttestationObject.AuthData.Flags.HasBackupState()
		model := WebauthnCredentialToModel(credential, sessionData.UserID, backupEligible, backupState)
		err = svc.repo.GetWebauthnCredentialRepo().Create(ctx, *model, credential.Transport)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		err = sessionDataRepo.Delete(ctx, *sessionData)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// TODO: audit logger
		userId = webauthnUser.UserId.String()
		credentialId = model.ID
		return nil
	}); err != nil {
		return credentialId, userId, err
	}
	return credentialId, userId, nil
}

func (svc *WebauthnSvc) InitLogin(ctx Ctx, reqUserId *string) (*protocol.CredentialAssertion, error) {
	var options *protocol.CredentialAssertion
	var sessionData *webauthn.SessionData
	if reqUserId != nil {
		// non discoverable login initialization
		userId, err := uuid.FromString(*reqUserId)
		if err != nil {
			// TODO: audit logger
			return nil, errorhandler.NewHTTPError(http.StatusBadRequest, "failed to parse UserID as uuid").SetInternal(err)
		}
		var webauthnUser *WebauthnUser
		webauthnUser, _, err = svc.getWebauthnUser(ctx, userId)
		if err != nil {
			return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get user: %w", err))
		}
		if webauthnUser == nil {
			// TODO: audit logger
			return nil, errorhandler.NewHTTPError(http.StatusBadRequest, "user not found")
		}

		if len(webauthnUser.WebAuthnCredentials()) > 0 {
			options, sessionData, err = svc.waClient.BeginLogin(webauthnUser, webauthn.WithUserVerification(protocol.VerificationRequired))
			if err != nil {
				return nil, fmt.Errorf("failed to create webauthn assertion options: %w", err)
			}
		}
	}

	if options == nil && sessionData == nil {
		var err error
		options, sessionData, err = svc.waClient.BeginDiscoverableLogin(webauthn.WithUserVerification(protocol.VerificationRequired))
		if err != nil {
			return nil, fmt.Errorf("failed to create webauthn assertion options for discoverable login: %w", err)
		}
	}

	err := svc.repo.GetWebauthnSessionRepo().Create(ctx, *WebauthnSessionDataToModel(sessionData, WebauthnOperationAuthentication))
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Remove all transports, because of a bug in android and windows where the internal authenticator gets triggered,
	// when the transports array contains the type 'internal' although the credential is not available on the device.
	for i := range options.Response.AllowedCredentials {
		options.Response.AllowedCredentials[i].Transport = nil
	}

	return options, nil
}

func (svc *WebauthnSvc) FinishLogin(ctx Ctx, request *protocol.ParsedCredentialAssertionData) (credentialId string, userId string, err error) {
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		sessionDataRepo := svc.repo.GetWebauthnSessionRepo()
		sessionData, err := sessionDataRepo.GetByChallenge(ctx, request.Response.CollectedClientData.Challenge)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if sessionData != nil && sessionData.Operation != WebauthnOperationAuthentication {
			sessionData = nil
		}

		if sessionData == nil {
			// TODO: audit logger
			return errorhandler.NewHTTPError(http.StatusUnauthorized, "Stored challenge and received challenge do not match").SetInternal(errors.New("sessionData not found"))
		}

		credential, webauthnUser, err := svc.getCredentialFromLoginSession(ctx, request, sessionData)
		if err != nil {
			return err
		}

		var dbCred *ent.WebauthnCredential
		for i := range webauthnUser.WebauthnCredentials {
			if webauthnUser.WebauthnCredentials[i].ID == base64.RawURLEncoding.EncodeToString(credential.ID) {
				dbCred = webauthnUser.WebauthnCredentials[i]
				break
			}
		}
		if dbCred != nil {
			if dbCred.BackupEligible != request.Response.AuthenticatorData.Flags.HasBackupEligible() || dbCred.BackupState != request.Response.AuthenticatorData.Flags.HasBackupState() {
				dbCred.BackupState = request.Response.AuthenticatorData.Flags.HasBackupState()
				dbCred.BackupEligible = request.Response.AuthenticatorData.Flags.HasBackupEligible()
			}

			now := time.Now().UTC()
			dbCred.LastUsedAt = now

			err = svc.repo.GetWebauthnCredentialRepo().Update(ctx, *dbCred)
			if err != nil {
				return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		err = sessionDataRepo.Delete(ctx, *sessionData)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		userId = webauthnUser.UserId.String()
		credentialId = base64.RawURLEncoding.EncodeToString(credential.ID)
		return nil
	}); err != nil {
		return credentialId, userId, err
	}
	return credentialId, userId, nil
}

func (svc *WebauthnSvc) ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error) {
	return svc.repo.GetWebauthnCredentialRepo().ListByUser(ctx, userId)
}

func (svc *WebauthnSvc) UpdateCredential(ctx Ctx, userId uuid.UUID, id string, name *string) error {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	credential, err := svc.repo.GetWebauthnCredentialRepo().GetById(ctx, id)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if credential == nil || credential.UserID.String() != user.ID.String() {
		return errorhandler.NewHTTPError(http.StatusNotFound).SetInternal(errors.New("the user does not have a webauthn credential with the specified credentialId"))
	}
	if name != nil {
		credential.Name = *name
	}

	return svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		return svc.repo.GetWebauthnCredentialRepo().Update(ctx, *credential)
	})
}

func (svc *WebauthnSvc) DeleteCredential(ctx Ctx, userId uuid.UUID, id string) error {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	credential, err := svc.repo.GetWebauthnCredentialRepo().GetById(ctx, id)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if credential == nil || credential.UserID.String() != user.ID.String() {
		return errorhandler.NewHTTPError(http.StatusNotFound).SetInternal(errors.New("the user does not have a webauthn credential with the specified credentialId"))
	}

	return svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		return svc.repo.GetWebauthnCredentialRepo().Delete(ctx, *credential)
	})
}

func (svc *WebauthnSvc) getWebauthnUser(ctx Ctx, userId uuid.UUID) (*WebauthnUser, *ent.User, error) {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return nil, nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return nil, nil, nil
	}

	credentials, err := svc.repo.GetWebauthnCredentialRepo().ListByUser(ctx, user.ID)
	if err != nil {
		return nil, nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webauthnUser, err := NewWebauthnUser(ctx, *user, credentials)
	if err != nil {
		return nil, nil, err
	}

	return webauthnUser, user, nil
}

func (svc *WebauthnSvc) getCredentialFromLoginSession(ctx Ctx, request *protocol.ParsedCredentialAssertionData, sessionData *ent.WebauthnSessionData) (credential *webauthn.Credential, webauthnUser *WebauthnUser, err error) {
	model := WebauthnSessionDataFromModel(sessionData)
	if sessionData.UserID.IsNil() {
		// Discoverable Login
		userId, err := uuid.FromBytes(request.Response.UserHandle)
		if err != nil {
			return nil, nil, errorhandler.NewHTTPError(http.StatusBadRequest, "failed to parse userHandle as uuid").SetInternal(err)
		}
		webauthnUser, _, err = svc.getWebauthnUser(ctx, userId)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get webauthn user: %w", err)
		}

		if webauthnUser == nil {
			// TODO: audit logger
			return nil, nil, errorhandler.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.New("user not found"))
		}

		credential, err = svc.waClient.ValidateDiscoverableLogin(func(rawID, userHandle []byte) (user webauthn.User, err error) {
			return webauthnUser, nil
		}, *model, request)
		if err != nil {
			// TODO: audit logger
			return nil, nil, errorhandler.NewHTTPError(http.StatusUnauthorized, "failed to validate assertion").SetInternal(err)
		}
	} else {
		// non discoverable Login
		webauthnUser, _, err = svc.getWebauthnUser(ctx, sessionData.UserID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get webauthn user: %w", err)
		}
		if webauthnUser == nil {
			// TODO: audit logger
			return nil, nil, errorhandler.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.New("user not found"))
		}
		credential, err = svc.waClient.ValidateLogin(webauthnUser, *model, request)
		if err != nil {
			// TODO: audit logger
			return nil, nil, errorhandler.NewHTTPError(http.StatusUnauthorized, "failed to validate assertion").SetInternal(err)
		}
	}
	return credential, webauthnUser, nil
}
