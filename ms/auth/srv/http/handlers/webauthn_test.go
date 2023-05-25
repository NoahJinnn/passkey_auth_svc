package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/auth/app/wa"
	test "github.com/hellohq/hqservice/ms/auth/test/mock/app"
	testRepo "github.com/hellohq/hqservice/ms/auth/test/mock/dal"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	userIdBytes = []byte{0x37, 0x45, 0x37, 0x77, 0x53, 0x56, 0x75, 0x49, 0x51, 0x79, 0x47, 0x68, 0x63, 0x79, 0x47, 0x77, 0x37, 0x5f, 0x42, 0x71, 0x42, 0x41}
	credentials = []*ent.WebauthnCredential{
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
)

var waSessionData = []*ent.WebauthnSessionData{
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
			Operation:        wa.WebauthnOperationRegistration,
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
			Operation:        wa.WebauthnOperationAuthentication,
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

	appl := test.NewApp(nil, nil, &defaultCfg, testRepo.NewRepo(nil, users, credentials, waSessionData, nil, nil, nil))
	handler := NewWebauthnHandler(&HttpDeps{
		Appl:      appl,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}, &sessionManager{})
	err = handler.InitRegistration(c)
	require.NoError(t, err)
	if assert.NoError(t, handler.InitRegistration(c)) {
		creationOptions := protocol.CredentialCreation{}

		b := bytes.NewBuffer(rec.Body.Bytes())
		d := json.NewDecoder(b)
		err = d.Decode(&creationOptions)
		assert.NoError(t, err)

		respUserId, ok := (creationOptions.Response.User.ID).(string)
		assert.True(t, ok)
		assert.Equal(t, userIdBytes, []byte(respUserId))

		assert.NotEmpty(t, creationOptions.Response.Challenge)
		assert.Equal(t, defaultCfg.Webauthn.RelyingParty.Id, creationOptions.Response.RelyingParty.ID)
		assert.Equal(t, creationOptions.Response.AuthenticatorSelection.ResidentKey, protocol.ResidentKeyRequirementRequired)
		assert.Equal(t, creationOptions.Response.AuthenticatorSelection.UserVerification, protocol.VerificationRequired)
		assert.Equal(t, creationOptions.Response.User.Name, users[0].Edges.Emails[0].Address)
		assert.Equal(t, creationOptions.Response.User.DisplayName, users[0].Edges.Emails[0].Address)
		assert.Equal(t, creationOptions.Response.User.Icon, "")
		assert.True(t, *creationOptions.Response.AuthenticatorSelection.RequireResidentKey)

	}
}

func TestWebauthnHandler_FinishRegistration(t *testing.T) {
	body := `{
"id": "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
"rawId": "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
"type": "public-key",
"response": {
"attestationObject": "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjeSZYN5YgOjGh0NBcPZHZgW4_krrmihjLHmVzzuoMdl2NFYmehnq3OAAI1vMYKZIsLJfHwVQMAWgGhXZHA-Erj4xfo8FKEcB_PmR7mOUVuOn7GZhLwV-kTSh2hrVc6QE7NOikFYXiDo2M_mJ3huHJkDnnc5dHtIxfedbpMdex5fY3hoFs-fwymQjtdqdvti5c4x6UBAgMmIAEhWCDxvVrRgK4vpnr6JxTx-KfpSNyQUtvc47ryryZmj-P5kSJYIDox8N9bHQBrxN-b5kXqfmj3GwAJW7nNCh8UPbus3B6I",
"clientDataJSON": "eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoidE9yTkRDRDJ4UWY0ekZqRWp3eGFQOGZPRXJQM3p6MDhyTW9UbEpHdG5LVSIsIm9yaWdpbiI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsImNyb3NzT3JpZ2luIjpmYWxzZX0"
}
}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webauthn/registration/finalize", strings.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	token := jwt.New()
	err := token.Set(jwt.SubjectKey, userId)
	require.NoError(t, err)
	c.Set("session", token)

	appl := test.NewApp(nil, nil, &defaultCfg, testRepo.NewRepo(nil, users, nil, waSessionData, nil, nil, nil))
	handler := NewWebauthnHandler(&HttpDeps{
		Appl:      appl,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}, &sessionManager{})

	if assert.NoError(t, handler.FinishRegistration(c)) {
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		assert.Regexp(t, `{"credential_id":".*"}`, rec.Body.String())
	}

	req2 := httptest.NewRequest(http.MethodPost, "/webauthn/registration/finalize", strings.NewReader(body))
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	token2 := jwt.New()
	err = token.Set(jwt.SubjectKey, userId)
	require.NoError(t, err)
	c2.Set("session", token2)

	err = handler.FinishRegistration(c2)
	if assert.Error(t, err) {
		httpError := errorhandler.ToHttpError(err)
		assert.Equal(t, http.StatusBadRequest, httpError.Code)
		assert.Equal(t, "Stored challenge and received challenge do not match: sessionData not found", err.Error())
	}
}

func TestWebauthnHandler_BeginLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webauthn/login/begin", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	token := jwt.New()
	err := token.Set(jwt.SubjectKey, userId)
	require.NoError(t, err)
	c.Set("session", token)

	appl := test.NewApp(nil, nil, &defaultCfg, testRepo.NewRepo(nil, users, nil, waSessionData, nil, nil, nil))
	handler := NewWebauthnHandler(&HttpDeps{
		Appl:      appl,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}, &sessionManager{})
	require.NoError(t, err)

	if assert.NoError(t, handler.InitLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		assertionOptions := protocol.CredentialAssertion{}
		err = json.Unmarshal(rec.Body.Bytes(), &assertionOptions)
		assert.NoError(t, err)
		assert.NotEmpty(t, assertionOptions.Response.Challenge)
		assert.Equal(t, assertionOptions.Response.UserVerification, protocol.VerificationRequired)
		assert.Equal(t, defaultCfg.Webauthn.RelyingParty.Id, assertionOptions.Response.RelyingPartyID)
	}
}

func TestWebauthnHandler_FinishLogin(t *testing.T) {
	body := `{
		"id": "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
		"rawId": "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
		"type": "public-key",
		"response": {
		"authenticatorData": "SZYN5YgOjGh0NBcPZHZgW4_krrmihjLHmVzzuoMdl2MFYmezOw",
		"clientDataJSON": "eyJ0eXBlIjoid2ViYXV0aG4uZ2V0IiwiY2hhbGxlbmdlIjoiZ0tKS21oOTB2T3BZTzU1b0hwcWFIWF9vTUNxNG9UWnQtRDBiNnRlSXpyRSIsIm9yaWdpbiI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsImNyb3NzT3JpZ2luIjpmYWxzZX0",
		"signature": "MEYCIQDi2vYVspG6pf38I4GyQCPOojGbvX4nwSPXCi0hm80twAIhAO3EWjhAnj0UpjU_l0AH5sEh3zq4LDvkvo3AUqaqfGYD",
		"userHandle": "7E7wSVuIQyGhcyGw7_BqBA"
		}
		}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webauthn/login/finalize", strings.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	appl := test.NewApp(nil, nil, &defaultCfg, testRepo.NewRepo(nil, users, credentials, waSessionData, nil, nil, nil))
	handler := NewWebauthnHandler(&HttpDeps{
		Appl:      appl,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}, &sessionManager{})
	if assert.NoError(t, handler.FinishLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		cookies := rec.Result().Cookies()
		if assert.NotEmpty(t, cookies) {
			for _, cookie := range cookies {
				if cookie.Name == "hqservice" {
					assert.Equal(t, userId, cookie.Value)
				}
			}
		}
	}

	req2 := httptest.NewRequest(http.MethodPost, "/webauthn/login/finalize", strings.NewReader(body))
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	err := handler.FinishLogin(c2)
	if assert.Error(t, err) {
		httpError := errorhandler.ToHttpError(err)
		assert.Equal(t, http.StatusUnauthorized, httpError.Code)
		assert.Equal(t, "Stored challenge and received challenge do not match: sessionData not found", err.Error())
	}
}
