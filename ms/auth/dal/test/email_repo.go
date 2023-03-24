package test

import (
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

func NewEmailRepo(init []*ent.Email) dal.IEmailRepo {
	return &emailRepo{append([]*ent.Email{}, init...)}
}

type emailRepo struct {
	emails []*ent.Email
}

func (r *emailRepo) GetByAddress(ctx Ctx, address string) (*ent.Email, error) {

	for _, m := range r.emails {
		if m.Address == address {
			return m, nil
		}
	}

	return nil, nil
}
