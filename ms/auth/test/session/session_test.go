package test

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGenerator(t *testing.T) {
	manager := JwkManager{}
	cfg := config.Session{}
	sessionGenerator, err := session.NewManager(&manager, cfg)
	assert.NoError(t, err)
	require.NotEmpty(t, sessionGenerator)
}

func TestGenerator_Generate(t *testing.T) {
	manager := JwkManager{}
	cfg := config.Session{}
	sessionGenerator, err := session.NewManager(&manager, cfg)
	assert.NoError(t, err)
	require.NotEmpty(t, sessionGenerator)

	userId, err := uuid.NewV4()
	assert.NoError(t, err)

	session, err := sessionGenerator.GenerateJWT(userId.String())
	assert.NoError(t, err)
	require.NotEmpty(t, session)
}

func TestGenerator_Verify(t *testing.T) {
	sessionLifespan := "5m"
	manager := JwkManager{}
	cfg := config.Session{Lifespan: sessionLifespan}
	sessionGenerator, err := session.NewManager(&manager, cfg)
	assert.NoError(t, err)
	require.NotEmpty(t, sessionGenerator)

	userId, err := uuid.NewV4()
	assert.NoError(t, err)

	session, err := sessionGenerator.GenerateJWT(userId.String())
	assert.NoError(t, err)
	require.NotEmpty(t, session)

	token, err := sessionGenerator.Verify(session)
	assert.NoError(t, err)
	require.NotEmpty(t, token)
	assert.Equal(t, token.Subject(), userId.String())
	assert.False(t, time.Time{}.Equal(token.IssuedAt()))
	assert.False(t, time.Time{}.Equal(token.Expiration()))

	sessionDuration, _ := time.ParseDuration(sessionLifespan)
	assert.True(t, token.IssuedAt().Add(sessionDuration).Equal(token.Expiration()))
}

func TestGenerator_Verify_Error(t *testing.T) {
	manager := JwkManager{}
	cfg := config.Session{}
	sessionGenerator, err := session.NewManager(&manager, cfg)
	assert.NoError(t, err)
	require.NotEmpty(t, sessionGenerator)

	tests := []struct {
		Name    string
		Input   string
		WantErr error
	}{
		{
			Name:  "expired",
			Input: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImtleTEiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE2NTIxNzM4NTAsImlhdCI6MTY1MjE3Mzc5MCwic3ViIjoiYzU0YTZlZTUtZjNmMy00YzExLTlhMWMtYWUzOTgyYTQxMzYyIn0.AEzZ0M1_3HnOtqd8Dz-BliHkEUc4c5mu97eXhoErgG7qbVWisJP0qfz_KrwL9VYFOYuDAmfRZ3ABnaOg-S53wlRndfL-ulk68lY34otGZfXKhk2P3GJRH8Dq7hW83KnwkSPF5_iOaIIDfUwrWOaavvtLJFgg5RcehuwLkYEA5X17ek6cUNsqz7Vw-x2REReh_f31f5zneqKN9CeVnup5_ZgtMYpOXVvXAORs3b7y2oMwFdXs-hVal9ZVunNPo3iZmaTFMHUSNXX8MceOy_dUofxtd9JDzliiPrjNWDjU5Jx5paLBA5CUc4SctBURi2oJABbkeE1l4ug6-rTOYB04-UW8XAnPZONBTnv3AjtzvScvkpUj-OFKVQLGgcXZHUo1J7ftLaezpWrGTbhlC8TVvXdX1ms5w9D1uqEUZ94lhvVSW_lGGX2DGqMWaT6tOcSpDHFQ0NR5FD3MiNGV-z43AUOOSzilAKS2WaHDS7v43PeJ75xzAAS_7xOoc6L3Z9msdToQIauLYuCrivoOVcCqrEHugknpxO8M0xo6f1fHws8RocT3S7B76YJUIBeAj2F31wJne5xtbiRF5GWiV8uS3ZTXqrPp7y4U6Btf-h6mvEos_Q9o9w5hck-8lixUs5mObPDsT-W6PdEehRaSL7-13dy1GpB8wMP5fGlnRSff9y4",
		},
		{
			Name:  "wrong signature",
			Input: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwOi8vbG9jYWxob3N0OjgwODAiLCJleHAiOjE2NDg2MjgzODAsImlhdCI6MTY0ODYyNDc4MCwiaXNzIjoiSG90M3RlOWlJbHVwbm5scDBld1E1RWJmIiwic3ViIjoiaU9ielRWdUc2bjl6Um1TY2ZMb0RFT1NiIn0.QZPEyEaGCJikNP2slVTGdsT3x8CuT4ynd5tdj-7c38Aa54277MPgGbapQ7JGrvwyjhAihzvvlqCxn2oX3zIFdu0HmSlxAXQ6Ah1K0KlQabneG8XNSed3sgp9xM0BYV1rB2SCuyXwE3U3zj5zFc4g4-v2Y1hpn7z4n3n9IlnnShK7NTUaaELlWPD8FQyp8mzZmJVSDoWbCMdywGHkX5ZWMUAwPfvC17kYZj6nqXC5ZJm3i2u-488cDeE5NxCFe-0ey14NtNtM9xTaPy5U8zvoqeCik1-ZNbxR_NJC4H25Cth2pm__e-W4KepGy7i-cLZ1T_DqNNk8HX9zX_Quj88FJw",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := sessionGenerator.Verify(test.Input)
			assert.Error(t, err)
		})
	}
}

func TestGenerator_DeleteCookie(t *testing.T) {
	manager := JwkManager{}
	cfg := config.Session{}
	sessionGenerator, err := session.NewManager(&manager, cfg)
	assert.NoError(t, err)
	require.NotEmpty(t, sessionGenerator)

	cookie, err := sessionGenerator.DeleteCookie()
	assert.NoError(t, err)
	assert.Equal(t, -1, cookie.MaxAge)
	assert.Equal(t, "hqservice", cookie.Name)
}
