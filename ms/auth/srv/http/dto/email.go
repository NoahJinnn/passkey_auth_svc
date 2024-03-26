package dto

import (
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/gofrs/uuid"
)

type EmailResponse struct {
	ID      uuid.UUID `json:"id"`
	Address string    `json:"address"`
}

// FromEmailModel Converts the DB model to a DTO object
func FromEmailModel(email *ent.Email) *EmailResponse {
	return &EmailResponse{
		ID:      email.ID,
		Address: email.Address,
	}
}

type EmailCreateBody struct {
	Address string `json:"address"`
}
