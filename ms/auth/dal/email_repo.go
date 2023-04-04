package dal

import (
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
)

type IEmailRepo interface {
	GetByAddress(ctx Ctx, address string) (*ent.Email, error)
}

type emailRepo struct {
	db *ent.Client
}

func NewEmailRepo(db *ent.Client) IEmailRepo {
	return &emailRepo{db: db}
}

func (r *emailRepo) GetByAddress(ctx Ctx, address string) (*ent.Email, error) {
	e, err := r.db.Email.
		Query().
		Where(email.Address(address)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying email by address: %w", err)
	}

	return e, nil
}
