package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
)

type IEmailRepo interface {
	GetByAddress(ctx Ctx, address string) (*ent.Email, error)
	GetById(ctx Ctx, id uuid.UUID) (*ent.Email, error)
}

type emailRepo struct {
	db *ent.Client
}

func NewEmailRepo(db *ent.Client) IEmailRepo {
	return &emailRepo{db: db}
}

func (r *emailRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.Email, error) {
	e, err := r.db.Email.
		Query().
		Where(email.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *emailRepo) GetByAddress(ctx Ctx, address string) (*ent.Email, error) {
	e, err := r.db.Email.
		Query().
		Where(email.Address(address)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return e, nil
}
