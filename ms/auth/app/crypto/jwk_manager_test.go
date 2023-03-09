package jwk

import (
	"testing"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/test"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockJwkPersister struct {
	jwks []ent.Jwk
}

func (m *mockJwkPersister) Get(i int) (*ent.Jwk, error) {
	for _, v := range m.jwks {
		if v.ID == uint(i) {
			return &v, nil
		}
	}
	return nil, nil
}

func (m *mockJwkPersister) GetAll() ([]ent.Jwk, error) {
	return m.jwks, nil
}

func (m *mockJwkPersister) GetLast() (*ent.Jwk, error) {
	index := len(m.jwks)
	return &m.jwks[index-1], nil
}

func (m *mockJwkPersister) Create(jwk ent.Jwk) error {
	//increment id
	index := len(m.jwks)
	jwk.ID = uint(index)

	m.jwks = append(m.jwks, jwk)
	return nil
}

func TestDefaultManager(t *testing.T) {
	keys := []string{"asfnoadnfoaegnq3094intoaegjnoadjgnoadng", "apdisfoaiegnoaiegnbouaebgn982"}

	persister := test.NewJwkPersister(nil)

	dm, err := NewDefaultManager(keys, persister)
	require.NoError(t, err)
	all, err := persister.GetAll()

	require.NoError(t, err)
	assert.Equal(t, 2, len(all))

	js, err := dm.GetPublicKeys()
	require.NoError(t, err)
	assert.Equal(t, 2, js.Len())

	sk, err := dm.GetSigningKey()
	require.NoError(t, err)

	token := jwt.New()
	token.Set("Payload", "isJustFine")
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, sk))
	require.NoError(t, err)

	// Get Public Key of signing key
	pk, err := sk.PublicKey()
	require.NoError(t, err)

	// Parse and Verify
	tokenParsed, err := jwt.Parse(signed, jwt.WithKey(jwa.RS256, pk))
	assert.NoError(t, err)
	assert.Equal(t, token, tokenParsed)
}
