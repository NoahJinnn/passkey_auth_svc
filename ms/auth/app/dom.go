package app

import (
	"encoding/binary"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/api/openapi/model"
	"github.com/hellohq/hqservice/ent"
)

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
}

type UserSession struct {
	Username string
	Webauthn *webauthn.SessionData
	// U2F      *u2f.Challenge
}

func (u *User) FromEnt(eu *ent.User) *User {
	return &User{
		ID:          eu.ID,
		FirstName:   eu.FirstName,
		LastName:    eu.LastName,
		Email:       eu.Email,
		PhoneNumber: eu.PhoneNumber,
		Address:     eu.Address,
	}
}

func UserListFromEnt(eus []*ent.User) []*User {
	us := make([]*User, len(eus))
	for _, eu := range eus {
		u := &User{}
		us = append(us, u.FromEnt(eu))
	}
	return us
}

func (u *User) ToOAIResp() *model.User {
	respId := int64(u.ID)
	return &model.User{
		ID:          &respId,
		FirstName:   &u.FirstName,
		LastName:    &u.LastName,
		Email:       &u.Email,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
	}
}

func (u *User) FromOAIReq(oaiU *model.User) *User {
	return &User{
		ID:          uint(*oaiU.ID),
		FirstName:   *oaiU.FirstName,
		LastName:    *oaiU.LastName,
		Email:       *oaiU.Email,
		PhoneNumber: oaiU.PhoneNumber,
		Address:     oaiU.Address,
	}
}

func (u *User) WebAuthnID() []byte {
	a := make([]byte, 4)
	h := uint64(u.ID)
	binary.LittleEndian.PutUint64(a, h)
	return a
}

func (u *User) WebAuthnName() string {
	return "newUser"
}

func (u *User) WebAuthnDisplayName() string {
	return "New User"
}

func (u *User) WebAuthnIcon() string {
	return "https://pics.com/avatar.png"
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{}
}
