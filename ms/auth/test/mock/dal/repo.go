package test

import (
	"context"

	"github.com/hellohq/hqservice/ent"
)

type Ctx = context.Context

func NewRepo(
	db *ent.Client,
	user []*ent.User,
	waCredential []*ent.WebauthnCredential,
	waSessionData []*ent.WebauthnSessionData,
	passcode []*ent.Passcode,
	email []*ent.Email,
	primaryEmail *ent.PrimaryEmail,
) *repoT {
	return &repoT{
		Db:                     db,
		userRepo:               NewUserRepo(user),
		webAuthnCredentialRepo: NewWebauthnCredentialRepo(waCredential),
		webAuthnSessionRepo:    NewWebauthnSessionRepo(waSessionData),
		emailRepo:              NewEmailRepo(email, primaryEmail),
		passcodeRepo:           NewPasscodeRepo(passcode),
	}
}

type repoT struct {
	Db                     *ent.Client
	userRepo               *userRepo
	webAuthnCredentialRepo *waCredentialRepo
	webAuthnSessionRepo    *waSessionRepo
	emailRepo              *emailRepo
	passcodeRepo           *passcodeRepo
}

func (r repoT) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	txForw := func(db *ent.Client) error {
		return exec(ctx, db)
	}
	return txForw(nil)
}

func (r repoT) GetUserRepo() *userRepo {
	return r.userRepo
}

func (r repoT) GetWebauthnCredentialRepo() *waCredentialRepo {
	return r.webAuthnCredentialRepo
}

func (r repoT) GetWebauthnSessionRepo() *waSessionRepo {
	return r.webAuthnSessionRepo
}

func (r repoT) GetEmailRepo() *emailRepo {
	return r.emailRepo
}

func (r repoT) GetPasscodeRepo() *passcodeRepo {
	return r.passcodeRepo
}
