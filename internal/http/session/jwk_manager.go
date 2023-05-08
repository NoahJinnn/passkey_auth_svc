package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/sharedDal"
	"github.com/hellohq/hqservice/pkg/crypto/aes_gcm"
	hqJwk "github.com/hellohq/hqservice/pkg/crypto/jwk"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JwkManager interface {
	// GenerateKey is used to generate a jwk Key
	GenerateKey(ctx context.Context) (jwk.Key, error)
	// GetPublicKeys returns all Public keys that are persisted
	GetPublicKeys(ctx context.Context) (jwk.Set, error)
	// GetSigningKey returns the last added private key that is used for signing
	GetSigningKey(ctx context.Context) (jwk.Key, error)
}

type DefaultManager struct {
	encrypter *aes_gcm.AESGCM
	repo      sharedDal.IJwkRepo
	keys      []string
}

// Returns a DefaultManager that reads and persists the jwks to database and generates jwks if a new secret gets added to the config.
func NewDefaultManager(keys []string, repo sharedDal.IJwkRepo) (*DefaultManager, error) {
	encrypter, err := aes_gcm.NewAESGCM(keys)
	if err != nil {
		return nil, err
	}
	manager := &DefaultManager{
		encrypter: encrypter,
		repo:      repo,
		keys:      keys,
	}

	return manager, nil
}

func (m *DefaultManager) InitJwk() error {
	// for every key we should check if a jwk with index exists and create one if not.
	ctx := context.Background()
	for i := range m.keys {

		j, err := m.repo.GetJwk(ctx, uint(i))
		if j == nil && err == nil {
			fmt.Printf("jwk with index %d does not exist, creating one\n", i)
			_, err := m.GenerateKey(ctx)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}
	return nil
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
	err = m.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (m *DefaultManager) GetSigningKey(ctx context.Context) (jwk.Key, error) {
	sigModel, err := m.repo.GetLastJwk(ctx)
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
	modelList, err := m.repo.GetAllJwk(ctx)
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
