package app

import "github.com/hellohq/hqservice/ent"

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
}

func (u *User) FromEnt(eu *ent.User) *User {
	return &User{
		ID:          eu.ID,
		FirstName:   eu.FirstName,
		LastName:    eu.LastName,
		Email:       eu.Email,
		Password:    eu.Password,
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

type AssetInfo struct {
	ID              uint
	AccountInfo     struct{}
	InstitutionInfo struct{}
	AssetInfo       struct{}
	SensibleData    string
	Descriptions    string
}

type BankAccount struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Car struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Collectible struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type CryptoAccount struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Loan struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type PrivateShare struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}
