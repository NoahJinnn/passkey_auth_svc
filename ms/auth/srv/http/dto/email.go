package dto

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
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
