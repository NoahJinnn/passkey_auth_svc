package test

import (
	"context"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/session"
)

func NewJwkRepo(init []*ent.Jwk) session.IJwkRepo {
	if init == nil {
		return &jwkPersister{[]*ent.Jwk{}}
	}
	return &jwkPersister{append([]*ent.Jwk{}, init...)}
}

type jwkPersister struct {
	keys []*ent.Jwk
}

func (j *jwkPersister) Jwk(ctx context.Context, id uint) (*ent.Jwk, error) {
	var found *ent.Jwk
	for _, data := range j.keys {
		if data.ID == uint(id) {
			d := data
			found = d
		}
	}
	return found, nil
}

func (j *jwkPersister) All(ctx context.Context) ([]*ent.Jwk, error) {
	return j.keys, nil
}

func (j *jwkPersister) Last(ctx context.Context) (*ent.Jwk, error) {
	l := len(j.keys)
	if l == 0 {
		return nil, nil
	}
	return j.keys[l-1], nil
}

func (j *jwkPersister) Create(ctx context.Context, jwk ent.Jwk) error {
	var lastId uint = 0
	for _, key := range j.keys {
		if key.ID > uint(lastId) {
			lastId = key.ID
		}
	}
	jwk.ID = lastId
	j.keys = append(j.keys, &jwk)
	return nil
}
