package dal

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
)

type IWebauthnCredentialRepo interface {
	GetFromUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error)
	Create(ctx Ctx, credential ent.WebauthnCredential) error
}

type webauthnRepo struct {
	db *ent.Client
}

func NewWebauthnCredentialRepo(db *ent.Client) IWebauthnCredentialRepo {
	return &webauthnRepo{db: db}
}

func (r *webauthnRepo) GetFromUser(ctx Ctx, userId uuid.UUID) (credentials []*ent.WebauthnCredential, err error) {
	// Query all ent.WebauthnCredential by ent.User id and sort by created at return them
	credentials, err = r.db.WebauthnCredential.
		Query().
		Where(webauthncredential.HasUserWith(user.ID(userId))).
		Order(ent.Asc(webauthncredential.FieldCreatedAt)).
		All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return credentials, nil
}

func (r *webauthnRepo) Create(ctx Ctx, credential ent.WebauthnCredential) error {
	_, err := r.db.WebauthnCredential.Create().
		SetUserID(credential.UserID).
		SetPublicKey(credential.PublicKey).
		SetAttestationType(credential.AttestationType).
		SetAaguid(credential.Aaguid).
		SetSignCount(credential.SignCount).
		SetName(credential.Name).
		SetBackupEligible(credential.BackupEligible).
		SetBackupState(credential.BackupState).
		SetLastUsedAt(credential.LastUsedAt).
		AddWebauthnCredentialTransports(credential.Edges.WebauthnCredentialTransports...).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to store credential: %w", err)
	}
	return nil
}
