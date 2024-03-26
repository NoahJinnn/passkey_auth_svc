package dto

import (
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/gofrs/uuid"
)

type WebauthnCredentialUpdateBody struct {
	Name *string `json:"name"`
}

type WebauthnCredentialResponse struct {
	ID              string     `json:"id"`
	Name            *string    `json:"name,omitempty"`
	PublicKey       string     `json:"public_key"`
	AttestationType string     `json:"attestation_type"`
	AAGUID          uuid.UUID  `json:"aaguid"`
	LastUsedAt      *time.Time `json:"last_used_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	Transports      []string   `json:"transports"`
	BackupEligible  bool       `json:"backup_eligible"`
	BackupState     bool       `json:"backup_state"`
}

// FromWebauthnCredentialModel Converts the DB model to a DTO object
func FromWebauthnCredentialModel(c *ent.WebauthnCredential) *WebauthnCredentialResponse {
	transports := make([]string, len(c.Edges.WebauthnCredentialTransports))
	for i := range transports {
		transports = append(transports, c.Edges.WebauthnCredentialTransports[i].Name)
	}
	return &WebauthnCredentialResponse{
		ID:              c.ID,
		Name:            &c.Name,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		AAGUID:          c.Aaguid,
		LastUsedAt:      &c.LastUsedAt,
		CreatedAt:       c.CreatedAt,
		Transports:      transports,
		BackupEligible:  c.BackupEligible,
		BackupState:     c.BackupState,
	}
}
