package wa

import (
	"encoding/base64"
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
)

func WebauthnCredentialToModel(credential *webauthn.Credential, userId uuid.UUID, backupEligible bool, backupState bool) *ent.WebauthnCredential {
	now := time.Now().UTC()
	aaguid, _ := uuid.FromBytes(credential.Authenticator.AAGUID)
	credentialID := base64.RawURLEncoding.EncodeToString(credential.ID)

	c := &ent.WebauthnCredential{
		ID:              credentialID,
		UserID:          userId,
		PublicKey:       base64.RawURLEncoding.EncodeToString(credential.PublicKey),
		AttestationType: credential.AttestationType,
		Aaguid:          aaguid,
		SignCount:       int32(credential.Authenticator.SignCount),
		LastUsedAt:      now,
		CreatedAt:       now,
		UpdatedAt:       now,
		BackupEligible:  backupEligible,
		BackupState:     backupState,
	}

	return c
}

func WebauthnCredentialFromModel(credential *ent.WebauthnCredential) *webauthn.Credential {
	credId, _ := base64.RawURLEncoding.DecodeString(credential.ID)
	pKey, _ := base64.RawURLEncoding.DecodeString(credential.PublicKey)
	// Make sure credential transports is loaded into the model
	transports := credential.Edges.WebauthnCredentialTransports
	transport := make([]protocol.AuthenticatorTransport, len(transports))

	for i, t := range transports {
		transport[i] = protocol.AuthenticatorTransport(t.Name)
	}

	return &webauthn.Credential{
		ID:              credId,
		PublicKey:       pKey,
		AttestationType: credential.AttestationType,
		Authenticator: webauthn.Authenticator{
			AAGUID:    credential.Aaguid.Bytes(),
			SignCount: uint32(credential.SignCount),
		},
		Transport: transport,
	}
}
