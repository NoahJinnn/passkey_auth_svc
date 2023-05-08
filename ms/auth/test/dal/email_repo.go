package test

import (
	"github.com/gofrs/uuid"
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

func (r *emailRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.Email, error) {

	for _, m := range r.emails {
		if m.ID == id {
			return m, nil
		}
	}

	return nil, nil
}

func (r *emailRepo) ListByUser(ctx Ctx, userID uuid.UUID) ([]*ent.Email, error) {

	var emails []*ent.Email // TODO: mock this email list by user

	for _, m := range r.emails {
		if m.UserID == userID {
			emails = append(emails, m)
		}
	}

	return emails, nil
}
