package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	test "github.com/hellohq/hqservice/ms/auth/test/app"
	testRepo "github.com/hellohq/hqservice/ms/auth/test/dal"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userIdBytes = []byte{0xec, 0x4e, 0xf0, 0x49, 0x5b, 0x88, 0x43, 0x21, 0xa1, 0x73, 0x21, 0xb0, 0xef, 0xf0, 0x6a, 0x4}
var credentials = []*ent.WebauthnCredential{
	func() *ent.WebauthnCredential {
		uId, _ := uuid.FromString(userId)
		aaguid, _ := uuid.FromString("adce0002-35bc-c60a-648b-0b25f1f05503")
		return &ent.WebauthnCredential{
			ID:              "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
			UserID:          uId,
			PublicKey:       "pQECAyYgASFYIPG9WtGAri-mevonFPH4p-lI3JBS29zjuvKvJmaP4_mRIlggOjHw31sdAGvE35vmRep-aPcbAAlbuc0KHxQ9u6zcHog",
			AttestationType: "none",
			Aaguid:          aaguid,
			SignCount:       1650958750,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		}
	}(),
	func() *ent.WebauthnCredential {
		uId, _ := uuid.FromString(userId)
		aaguid, _ := uuid.FromString("adce0002-35bc-c60a-648b-0b25f1f05503")
		return &ent.WebauthnCredential{
			ID:              "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjK",
			UserID:          uId,
			PublicKey:       "pQECAyYgASFYIPG9WtGAri-mevonFPH4p-lI3JBS29zjuvKvJmaP4_mRIlggOjHw31sdAGvE35vmRep-aPcbAAlbuc0KHxQ9u6zcHoj",
			AttestationType: "none",
			Aaguid:          aaguid,
			SignCount:       1650958750,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		}
	}(),
}

var sessionData = []*ent.WebauthnSessionData{
	func() *ent.WebauthnSessionData {
		id, _ := uuid.NewV4()
		uId, _ := uuid.FromString(userId)
		return &ent.WebauthnSessionData{
			ID:               id,
			Challenge:        "tOrNDCD2xQf4zFjEjwxaP8fOErP3zz08rMoTlJGtnKU",
			UserID:           uId,
			UserVerification: string(protocol.VerificationRequired),
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			Operation:        svcs.WebauthnOperationRegistration,
			// AllowedCredentials: nil,
		}
	}(),
	func() *ent.WebauthnSessionData {
		id, _ := uuid.NewV4()
		return &ent.WebauthnSessionData{
			ID:               id,
			Challenge:        "gKJKmh90vOpYO55oHpqaHX_oMCq4oTZt-D0b6teIzrE",
			UserID:           uuid.UUID{},
			UserVerification: string(protocol.VerificationRequired),
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			Operation:        svcs.WebauthnOperationAuthentication,
			// AllowedCredentials: nil,
		}
	}(),
}

func TestWebauthnHandler_BeginRegistration(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webauthn/registration/initialize", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	token := jwt.New()
	err := token.Set(jwt.SubjectKey, userId)
	require.NoError(t, err)
	c.Set("session", token)

	appl := test.NewApp(&defaultConfig, testRepo.NewRepo(nil, users, nil, credentials, sessionData, nil))

	handler := NewWebauthnHandler(&HttpDeps{
		Appl: appl,
		Cfg:  &defaultConfig,
	})
	err = handler.BeginRegistration(c)
	require.NoError(t, err)
	if assert.NoError(t, handler.BeginRegistration(c)) {
		creationOptions := protocol.CredentialCreation{}

		b := bytes.NewBuffer(rec.Body.Bytes())
		d := json.NewDecoder(b)
		err = d.Decode(&creationOptions)
		assert.NoError(t, err)

		assert.NotEmpty(t, creationOptions.Response.Challenge)
		assert.Equal(t, userIdBytes, []byte(creationOptions.Response.User.ID))
		assert.Equal(t, defaultConfig.Webauthn.RelyingParty.Id, creationOptions.Response.RelyingParty.ID)
		assert.Equal(t, creationOptions.Response.AuthenticatorSelection.ResidentKey, protocol.ResidentKeyRequirementRequired)
		assert.Equal(t, creationOptions.Response.AuthenticatorSelection.UserVerification, protocol.VerificationRequired)
		assert.True(t, *creationOptions.Response.AuthenticatorSelection.RequireResidentKey)
	}
}
