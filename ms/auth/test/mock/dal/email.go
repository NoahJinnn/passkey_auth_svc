package test

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

func NewEmailRepo(init []*ent.Email, primaryEmail *ent.PrimaryEmail) dal.IEmailRepo {
	return &emailRepo{
		emails:       append([]*ent.Email{}, init...),
		primaryEmail: primaryEmail,
	}
}

type emailRepo struct {
	emails       []*ent.Email
	primaryEmail *ent.PrimaryEmail
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

func (r *emailRepo) GetPrimary(ctx Ctx, emailId uuid.UUID) (*ent.PrimaryEmail, error) {
	return r.primaryEmail, nil
}

func (r *emailRepo) ListByUser(ctx Ctx, userID uuid.UUID) ([]*ent.Email, error) {
	var emails []*ent.Email

	for _, m := range r.emails {
		if *m.UserID == userID {
			emails = append(emails, m)
		}
	}

	return emails, nil
}

func (r *emailRepo) CountByUserId(ctx Ctx, userID uuid.UUID) (int, error) {
	var count int

	for _, m := range r.emails {
		if *m.UserID == userID {
			count++
		}
	}

	return count, nil
}

func (r *emailRepo) Update(ctx Ctx, email *ent.Email) error {
	for i, m := range r.emails {
		if email.ID == m.ID {
			r.emails[i] = email
		}
	}
	return nil
}

func (r *emailRepo) UpdatePrimary(ctx Ctx, primary ent.PrimaryEmail) error {
	r.primaryEmail = &primary
	return nil
}

func (r *emailRepo) Delete(ctx Ctx, email *ent.Email) error {
	for i, m := range r.emails {
		if email.ID == m.ID {
			r.emails = append(r.emails[:i], r.emails[i+1:]...)
		}
	}
	return nil
}
