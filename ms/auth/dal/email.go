package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ent/primaryemail"
)

type IEmailRepo interface {
	GetByAddress(ctx Ctx, address string) (*ent.Email, error)
	GetById(ctx Ctx, id uuid.UUID) (*ent.Email, error)
	GetPrimary(ctx Ctx, emailId uuid.UUID) (*ent.PrimaryEmail, error)
	ListByUser(ctx Ctx, userID uuid.UUID) ([]*ent.Email, error)
	CountByUserId(ctx Ctx, userID uuid.UUID) (int, error)
	Update(ctx Ctx, email *ent.Email) error
	UpdatePrimary(ctx Ctx, primary ent.PrimaryEmail) error
	Delete(ctx Ctx, email *ent.Email) error
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

func (r *emailRepo) GetPrimary(ctx Ctx, emailId uuid.UUID) (*ent.PrimaryEmail, error) {
	e, err := r.db.PrimaryEmail.
		Query().
		Where(primaryemail.EmailID(emailId)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return e, nil
}

func (r *emailRepo) UpdatePrimary(ctx Ctx, primary ent.PrimaryEmail) error {
	_, err := r.db.PrimaryEmail.
		UpdateOneID(primary.ID).
		SetUserID(*primary.UserID).
		SetEmailID(primary.EmailID).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *emailRepo) ListByUser(ctx Ctx, userID uuid.UUID) ([]*ent.Email, error) {
	emails, err := r.db.Email.
		Query().
		Where(email.UserID(userID)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (r *emailRepo) CountByUserId(ctx Ctx, userID uuid.UUID) (int, error) {
	cnt, err := r.db.Email.
		Query().
		Where(email.UserID(userID)).
		Count(ctx)

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

func (r *emailRepo) Update(ctx Ctx, email *ent.Email) error {
	_, err := r.db.Email.
		UpdateOneID(email.ID).
		SetAddress(email.Address).
		SetVerified(email.Verified).
		SetUserID(*email.UserID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *emailRepo) Delete(ctx Ctx, email *ent.Email) error {
	return r.db.Email.DeleteOne(email).Exec(ctx)
}
