package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/passcode"
)

type IPasscodeRepo interface {
	GetById(ctx Ctx, id uuid.UUID) (*ent.Passcode, error)
	Create(ctx Ctx, pc *ent.Passcode) error
	Update(ctx Ctx, pc *ent.Passcode) error
	Delete(ctx Ctx, pc *ent.Passcode) error
}

type passcodeRepo struct {
	db *ent.Client
}

func NewPasscodeRepo(db *ent.Client) IPasscodeRepo {
	return &passcodeRepo{db: db}
}

func (r *passcodeRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.Passcode, error) {
	pc, err := r.db.Passcode.
		Query().
		Where(passcode.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return pc, nil
}

func (r *passcodeRepo) Create(ctx Ctx, pc *ent.Passcode) error {
	_, err := r.db.Passcode.
		Create().
		SetUserID(pc.UserID).
		SetEmailID(pc.EmailID).
		SetCode(pc.Code).
		SetTTL(pc.TTL).
		SetTryCount(0).
		SetCreatedAt(pc.CreatedAt).
		SetUpdatedAt(pc.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *passcodeRepo) Update(ctx Ctx, pc *ent.Passcode) error {
	_, err := r.db.Passcode.
		UpdateOne(pc).
		SetCode(pc.Code).
		SetTTL(pc.TTL).
		SetUpdatedAt(pc.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *passcodeRepo) Delete(ctx Ctx, pc *ent.Passcode) error {
	err := r.db.Passcode.
		DeleteOne(pc).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
