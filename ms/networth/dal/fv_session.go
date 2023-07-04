package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/fvsession"
)

type IFvSessionRepo interface {
	GetByUserId(ctx Ctx, id uuid.UUID) (*ent.FvSession, error)
	Create(ctx Ctx, s *ent.FvSession) (*ent.FvSession, error)
}

type fvSessionRepo struct {
	pgsql *ent.Client
}

func NewFvSessionRepo(pgsql *ent.Client) IFvSessionRepo {
	return &fvSessionRepo{pgsql: pgsql}
}

func (r *fvSessionRepo) GetByUserId(ctx Ctx, id uuid.UUID) (*ent.FvSession, error) {
	s, err := r.pgsql.FvSession.
		Query().
		Where(fvsession.UserID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *fvSessionRepo) Create(ctx Ctx, s *ent.FvSession) (*ent.FvSession, error) {
	news, err := r.pgsql.FvSession.
		Create().
		SetUserID(*s.UserID).
		SetAccessToken(s.AccessToken).
		SetExpiresIn(s.ExpiresIn).
		SetIssuedAt(s.IssuedAt).
		SetTokenType(s.TokenType).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return news, nil
}
