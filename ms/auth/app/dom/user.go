package dom

import (
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
