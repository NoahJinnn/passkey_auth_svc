package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/passcode"
)

type IPasscodeRepo interface {
	GetById(ctx Ctx, id uuid.UUID) (*ent.Passcode, error)
	Create(ctx Ctx, pc *ent.Passcode) (*ent.Passcode, error)
	Update(ctx Ctx, pc *ent.Passcode) error
	Delete(ctx Ctx, pc *ent.Passcode) error
}

type passcodeRepo struct {
	pgsql *ent.Client
}

func NewPasscodeRepo(pgsql *ent.Client) *passcodeRepo {
	return &passcodeRepo{pgsql: pgsql}
}

func (r *passcodeRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.Passcode, error) {
	pc, err := r.pgsql.Passcode.
		Query().
		Where(passcode.ID(id)).
		WithEmail().
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return pc, nil
}

func (r *passcodeRepo) Create(ctx Ctx, pc *ent.Passcode) (*ent.Passcode, error) {
	pc, err := r.pgsql.Passcode.
		Create().
		SetUserID(pc.UserID).
		SetEmailID(pc.EmailID).
		SetCode(pc.Code).
		SetTTL(pc.TTL).
		SetTryCount(0).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return pc, nil
}

func (r *passcodeRepo) Update(ctx Ctx, pc *ent.Passcode) error {
	_, err := r.pgsql.Passcode.
		UpdateOne(pc).
		SetCode(pc.Code).
		SetTTL(pc.TTL).
		SetTryCount(pc.TryCount).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *passcodeRepo) Delete(ctx Ctx, pc *ent.Passcode) error {
	err := r.pgsql.Passcode.
		DeleteOne(pc).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
