package dal

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
)

type IWebauthnCredentialRepo interface {
	GetFromUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error)
	GetById(ctx Ctx, id string) (*ent.WebauthnCredential, error)
	Create(ctx Ctx, credential ent.WebauthnCredential, transports []protocol.AuthenticatorTransport) error
	Update(ctx Ctx, credential ent.WebauthnCredential) error
	Delete(ctx Ctx, credential ent.WebauthnCredential) error
}

type webauthnRepo struct {
	db *ent.Client
}

func NewWebauthnCredentialRepo(db *ent.Client) IWebauthnCredentialRepo {
	return &webauthnRepo{db: db}
}

func (r *webauthnRepo) GetById(ctx Ctx, id string) (credential *ent.WebauthnCredential, err error) {
	credential, err = r.db.WebauthnCredential.
		Query().
		Where(webauthncredential.ID(id)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return credential, nil
}

func (r *webauthnRepo) GetFromUser(ctx Ctx, userId uuid.UUID) (credentials []*ent.WebauthnCredential, err error) {
	// Query all ent.WebauthnCredential by ent.User id and sort by created at return them
	credentials, err = r.db.WebauthnCredential.
		Query().
		Where(webauthncredential.HasUserWith(user.ID(userId))).
		WithWebauthnCredentialTransports().
		Order(ent.Asc(webauthncredential.FieldCreatedAt)).
		All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return credentials, nil
}

func (r *webauthnRepo) Create(ctx Ctx, credential ent.WebauthnCredential, transports []protocol.AuthenticatorTransport) error {
	bulk := make([]*ent.WebauthnCredentialTransportCreate, len(transports))
	for i, transport := range transports {
		bulk[i] = r.db.WebauthnCredentialTransport.Create().SetName(string(transport))
	}

	createdTransports, err := r.db.WebauthnCredentialTransport.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	_, err = r.db.WebauthnCredential.Create().
		SetID(credential.ID).
		SetUserID(credential.UserID).
		SetPublicKey(credential.PublicKey).
		SetAttestationType(credential.AttestationType).
		SetAaguid(credential.Aaguid).
		SetSignCount(credential.SignCount).
		SetName(credential.Name).
		SetBackupEligible(credential.BackupEligible).
		SetBackupState(credential.BackupState).
		SetLastUsedAt(credential.LastUsedAt).
		AddWebauthnCredentialTransports(createdTransports...).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (r *webauthnRepo) Update(ctx Ctx, credential ent.WebauthnCredential) error {
	_, err := r.db.WebauthnCredential.
		UpdateOneID(credential.ID).
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
		return err
	}
	return nil
}

func (r *webauthnRepo) Delete(ctx Ctx, credential ent.WebauthnCredential) error {
	err := r.db.WebauthnCredential.DeleteOneID(credential.ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
