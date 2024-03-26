package wa

import (
	"encoding/base64"
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
)

func WebauthnSessionDataFromModel(data *ent.WebauthnSessionData) *webauthn.SessionData {
	var allowedCredentials [][]byte
	for _, credential := range data.Edges.WebauthnSessionDataAllowedCredentials {
		credentialId, err := base64.RawURLEncoding.DecodeString(credential.CredentialID)
		if err != nil {
			continue
		}
		allowedCredentials = append(allowedCredentials, credentialId)
	}
	var userId []byte = nil
	if !data.UserID.IsNil() {
		userId = data.UserID.Bytes()
	}
	return &webauthn.SessionData{
		Challenge:            data.Challenge,
		UserID:               userId,
		AllowedCredentialIDs: allowedCredentials,
		UserVerification:     protocol.UserVerificationRequirement(data.UserVerification),
	}
}

func WebauthnSessionDataToModel(data *webauthn.SessionData, operation string) *ent.WebauthnSessionData {
	id, _ := uuid.NewV4()
	userId, _ := uuid.FromBytes(data.UserID)
	now := time.Now()

	var allowedCredentials []*ent.WebauthnSessionDataAllowedCredential
	for _, credentialID := range data.AllowedCredentialIDs {
		aId, _ := uuid.NewV4()
		allowedCredential := ent.WebauthnSessionDataAllowedCredential{
			ID:                    aId,
			CredentialID:          base64.RawURLEncoding.EncodeToString(credentialID),
			WebauthnSessionDataID: id,
			CreatedAt:             now,
			UpdatedAt:             now,
		}

		allowedCredentials = append(allowedCredentials, &allowedCredential)
	}

	sessionData := &ent.WebauthnSessionData{
		ID:               id,
		Challenge:        data.Challenge,
		UserID:           userId,
		UserVerification: string(data.UserVerification),
		CreatedAt:        now,
		UpdatedAt:        now,
		Operation:        operation,
	}
	sessionData.Edges.WebauthnSessionDataAllowedCredentials = allowedCredentials
	return sessionData
}
