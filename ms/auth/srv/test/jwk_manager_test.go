package test

import (
	"context"
	"testing"

	"github.com/hellohq/hqservice/ms/auth/dal/test"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultManager(t *testing.T) {
	keys := []string{"asfnoadnfoaegnq3094intoaegjnoadjgnoadng", "apdisfoaiegnoaiegnbouaebgn982"}
	ctx := context.Background()
	repo := test.NewJwkRepo(nil)

	dm, err := session.NewDefaultManager(keys, repo)
	require.NoError(t, err)
	all, err := repo.GetAllJwk(ctx)

	require.NoError(t, err)
	assert.Equal(t, 2, len(all))

	js, err := dm.GetPublicKeys(ctx)
	require.NoError(t, err)
	assert.Equal(t, 2, js.Len())

	sk, err := dm.GetSigningKey(ctx)
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
