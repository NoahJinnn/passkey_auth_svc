package test

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
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
) dal.IAuthRepo {
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
	userRepo               dal.IUserRepo
	webAuthnCredentialRepo dal.IWebauthnCredentialRepo
	webAuthnSessionRepo    dal.IWebauthnSessionRepo
	emailRepo              dal.IEmailRepo
	passcodeRepo           dal.IPasscodeRepo
}

func (r repoT) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	txForw := func(db *ent.Client) error {
		return exec(ctx, db)
	}
	return txForw(nil)
}

func (r repoT) GetUserRepo() dal.IUserRepo {
	return r.userRepo
}

func (r repoT) GetWebauthnCredentialRepo() dal.IWebauthnCredentialRepo {
	return r.webAuthnCredentialRepo
}

func (r repoT) GetWebauthnSessionRepo() dal.IWebauthnSessionRepo {
	return r.webAuthnSessionRepo
}

func (r repoT) GetEmailRepo() dal.IEmailRepo {
	return r.emailRepo
}

func (r repoT) GetPasscodeRepo() dal.IPasscodeRepo {
	return r.passcodeRepo
}
