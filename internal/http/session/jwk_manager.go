package session

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/pkg/crypto/aes_gcm"
	hqJwk "github.com/NoahJinnn/passkey_auth_svc/pkg/crypto/jwk"
	"github.com/gofrs/uuid"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type IJwkManager interface {
	// GenerateKey is used to generate a jwk Key
	GenerateKey(ctx context.Context) (jwk.Key, error)
	// GetPublicKeys returns all Public keys that are persisted
	GetPublicKeys(ctx context.Context) (jwk.Set, error)
	// GetSigningKey returns the last added private key that is used for signing
	GetSigningKey(ctx context.Context) (jwk.Key, error)
}

type DefaultManager struct {
	encrypter *aes_gcm.AESGCM
	jwkRepo   IJwkRepo
}

// Returns a DefaultManager that reads and persists the jwks to database and generates jwks if a new secret gets added to the config.
func NewDefaultManager(keys []string, jwkRepo IJwkRepo) (*DefaultManager, error) {
	encrypter, err := aes_gcm.NewAESGCM(keys)
	if err != nil {
		return nil, err
	}
	manager := &DefaultManager{
		encrypter: encrypter,
		jwkRepo:   jwkRepo,
	}
	// for every key we should check if a jwk with index exists and create one if not.
	ctx := context.Background()
	for i := range keys {

		j, err := jwkRepo.Jwk(ctx, uint(i))
		if j == nil && err == nil {
			_, err := manager.GenerateKey(ctx)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
	}

	return manager, nil
}

func (m *DefaultManager) GenerateKey(ctx context.Context) (jwk.Key, error) {
	rsa := &hqJwk.RSAKeyGenerator{}
	id, _ := uuid.NewV4()
	key, err := rsa.Generate(id.String())
	if err != nil {
		return nil, err
	}
	marshalled, err := json.Marshal(key)
	if err != nil {
		return nil, err
	}
	encryptedKey, err := m.encrypter.Encrypt(marshalled)
	if err != nil {
		return nil, err
	}
	model := ent.Jwk{
		KeyData:   encryptedKey,
		CreatedAt: time.Now(),
	}
	err = m.jwkRepo.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (m *DefaultManager) GetSigningKey(ctx context.Context) (jwk.Key, error) {
	sigModel, err := m.jwkRepo.Last(ctx)
	if err != nil {
		return nil, err
	}
	k, err := m.encrypter.Decrypt(sigModel.KeyData)
	if err != nil {
		return nil, err
	}

	key, err := jwk.ParseKey(k)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (m *DefaultManager) GetPublicKeys(ctx context.Context) (jwk.Set, error) {
	modelList, err := m.jwkRepo.All(ctx)
	if err != nil {
		return nil, err
	}

	publicKeys := jwk.NewSet()
	for _, model := range modelList {
		k, err := m.encrypter.Decrypt(model.KeyData)
		if err != nil {
			return nil, err
		}

		key, err := jwk.ParseKey(k)
		if err != nil {
			return nil, err
		}

		publicKey, err := jwk.PublicKeyOf(key)
		if err != nil {
			return nil, err
		}
		err = publicKeys.AddKey(publicKey)
		if err != nil {
			return nil, err
		}
	}

	return publicKeys, nil
}
