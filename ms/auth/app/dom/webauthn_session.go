package dom

import (
	"encoding/base64"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
)

func WebauthnSessionDataFromModel(data *ent.WebauthnSessionData) *webauthn.SessionData {
	var allowedCredentials [][]byte
	for _, credential := range data.AllowedCredentials {
		credentialId, err := base64.RawURLEncoding.DecodeString(credential.CredentialId)
		if err != nil {
			continue
		}
		allowedCredentials = append(allowedCredentials, credentialId)
	}
	var userId []byte = nil
	if !data.UserId.IsNil() {
		userId = data.UserId.Bytes()
	}
	return &webauthn.SessionData{
		Challenge:            data.Challenge,
		UserID:               userId,
		AllowedCredentialIDs: allowedCredentials,
		UserVerification:     protocol.UserVerificationRequirement(data.UserVerification),
	}
}

func WebauthnSessionDataToModel(data *webauthn.SessionData, operation ent.Operation) *ent.WebauthnSessionData {
	id, _ := uuid.NewV4()
	userId, _ := uuid.FromBytes(data.UserID)
	now := time.Now()

	var allowedCredentials []ent.WebauthnSessionDataAllowedCredential
	for _, credentialID := range data.AllowedCredentialIDs {
		aId, _ := uuid.NewV4()
		allowedCredential := ent.WebauthnSessionDataAllowedCredential{
			ID:                    aId,
			CredentialId:          base64.RawURLEncoding.EncodeToString(credentialID),
			WebauthnSessionDataID: id,
			CreatedAt:             now,
			UpdatedAt:             now,
		}

		allowedCredentials = append(allowedCredentials, allowedCredential)
	}

	return &ent.WebauthnSessionData{
		ID:                 id,
		Challenge:          data.Challenge,
		UserId:             userId,
		UserVerification:   string(data.UserVerification),
		CreatedAt:          now,
		UpdatedAt:          now,
		Operation:          operation,
		AllowedCredentials: allowedCredentials,
	}
}
