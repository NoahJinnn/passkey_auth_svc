package test

import (
	"context"
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

func NewRepo(db *ent.Client) dal.IRepo {
	return &repo{
		Db: db,
	}
}

type repo struct {
	Db *ent.Client
}

func (r repo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.Db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (r repo) GetJwkRepo() dal.IJwkRepo {
	return NewJwkRepo([]*ent.Jwk{})
}

func (r repo) GetUserRepo() dal.IUserRepo {
	return NewUserRepo([]*ent.User{})
}

func (r repo) GetWebauthnCredentialRepo() dal.IWebauthnCredentialRepo {
	return NewWebauthnCredentialRepo([]*ent.WebauthnCredential{})
}

func (r repo) GetWebauthnSessionRepo() dal.IWebauthnSessionRepo {
	return NewWebauthnSessionRepo([]*ent.WebauthnSessionData{})
}

func (r repo) GetEmailRepo() dal.IEmailRepo {
	return NewEmailRepo([]*ent.Email{})
}
