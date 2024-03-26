package test

import (
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/gofrs/uuid"
)

type passcodeRepo struct {
	passcodes []*ent.Passcode
}

func NewPasscodeRepo(init []*ent.Passcode) dal.IPasscodeRepo {
	return &passcodeRepo{append([]*ent.Passcode{}, init...)}
}

func (r *passcodeRepo) Create(ctx Ctx, passcode *ent.Passcode) (*ent.Passcode, error) {
	r.passcodes = append(r.passcodes, passcode)
	return passcode, nil
}

func (r *passcodeRepo) Update(ctx Ctx, passcode *ent.Passcode) error {
	for i, v := range r.passcodes {
		if v.ID == passcode.ID {
			r.passcodes[i] = passcode
			return nil
		}
	}
	return nil
}

func (r *passcodeRepo) Delete(ctx Ctx, passcode *ent.Passcode) error {
	for i, v := range r.passcodes {
		if v.ID == passcode.ID {
			r.passcodes = append(r.passcodes[:i], r.passcodes[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *passcodeRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.Passcode, error) {
	for _, v := range r.passcodes {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, nil
}
