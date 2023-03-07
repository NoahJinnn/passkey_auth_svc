package dom

import (
	"encoding/base64"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
)

func WebauthnCredentialToModel(credential *webauthn.Credential, userId uuid.UUID, backupEligible bool, backupState bool) *models.WebauthnCredential {
	now := time.Now().UTC()
	aaguid, _ := uuid.FromBytes(credential.Authenticator.AAGUID)
	credentialID := base64.RawURLEncoding.EncodeToString(credential.ID)

	c := &ent.WebauthnCredential{
		ID:              credentialID,
		UserId:          userId,
		PublicKey:       base64.RawURLEncoding.EncodeToString(credential.PublicKey),
		AttestationType: credential.AttestationType,
		AAGUID:          aaguid,
		SignCount:       int(credential.Authenticator.SignCount),
		LastUsedAt:      &now,
		CreatedAt:       now,
		UpdatedAt:       now,
		BackupEligible:  backupEligible,
		BackupState:     backupState,
	}

	for _, name := range credential.Transport {
		if string(name) != "" {
			id, _ := uuid.NewV4()
			t := ent.WebauthnCredentialTransport{
				ID:                   id,
				Name:                 string(name),
				WebauthnCredentialID: credentialID,
			}
			c.Transports = append(c.Transports, t)
		}
	}

	return c
}

func WebauthnCredentialFromModel(credential *ent.WebauthnCredential) *webauthn.Credential {
	credId, _ := base64.RawURLEncoding.DecodeString(credential.ID)
	pKey, _ := base64.RawURLEncoding.DecodeString(credential.PublicKey)
	transport := make([]protocol.AuthenticatorTransport, len(credential.Transports))

	for i, t := range credential.Transports {
		transport[i] = protocol.AuthenticatorTransport(t.Name)
	}

	return &webauthn.Credential{
		ID:              credId,
		PublicKey:       pKey,
		AttestationType: credential.AttestationType,
		Authenticator: webauthn.Authenticator{
			AAGUID:    credential.AAGUID.Bytes(),
			SignCount: uint32(credential.SignCount),
		},
		Transport: transport,
	}
}
