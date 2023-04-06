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
	jwk []*ent.Jwk,
	waCredential []*ent.WebauthnCredential,
	waSessionData []*ent.WebauthnSessionData,
	email []*ent.Email,
) dal.IRepo {
	return &repo{
		Db:                     db,
		userRepo:               NewUserRepo(user),
		jwkRepo:                NewJwkRepo(jwk),
		webAuthnCredentialRepo: NewWebauthnCredentialRepo(waCredential),
		webAuthnSessionRepo:    NewWebauthnSessionRepo(waSessionData),
		emailRepo:              NewEmailRepo(email),
	}
}

type repo struct {
	Db                     *ent.Client
	userRepo               dal.IUserRepo
	jwkRepo                dal.IJwkRepo
	webAuthnCredentialRepo dal.IWebauthnCredentialRepo
	webAuthnSessionRepo    dal.IWebauthnSessionRepo
	emailRepo              dal.IEmailRepo
}

func (r repo) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {

	txForw := func(db *ent.Client) error {
		return exec(ctx, db)
	}
	return txForw(nil)
}

func (r repo) GetJwkRepo() dal.IJwkRepo {
	return r.jwkRepo
}

func (r repo) GetUserRepo() dal.IUserRepo {
	return r.userRepo
}

func (r repo) GetWebauthnCredentialRepo() dal.IWebauthnCredentialRepo {
	return r.webAuthnCredentialRepo
}

func (r repo) GetWebauthnSessionRepo() dal.IWebauthnSessionRepo {
	return r.webAuthnSessionRepo
}

func (r repo) GetEmailRepo() dal.IEmailRepo {
	return r.emailRepo
}