package handlers

import (
	"testing"
)

var userIdBytes = []byte{0xec, 0x4e, 0xf0, 0x49, 0x5b, 0x88, 0x43, 0x21, 0xa1, 0x73, 0x21, 0xb0, 0xef, 0xf0, 0x6a, 0x4}

func (us integrationSuite) TestWebauthnHandler_BeginRegistration(t *testing.T) {
	us.T().Log("TestWebauthnHandler_BeginRegistration")
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/webauthn/registration/initialize", nil)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// token := jwt.New()
	// err := token.Set(jwt.SubjectKey, userId)
	// require.NoError(t, err)
	// c.Set("session", token)

	// p := test.NewRepo()
	// handler, err := NewWebauthnHandler(&defaultConfig, p, sessionManager{}, test.NewAuditLogger())
	// require.NoError(t, err)

	// if assert.NoError(t, handler.BeginRegistration(c)) {
	// 	creationOptions := protocol.CredentialCreation{}
	// 	err = json.Unmarshal(rec.Body.Bytes(), &creationOptions)
	// 	assert.NoError(t, err)
	// 	assert.NotEmpty(t, creationOptions.Response.Challenge)
	// 	assert.Equal(t, userIdBytes, []byte(creationOptions.Response.User.ID))
	// 	assert.Equal(t, defaultConfig.Webauthn.RelyingParty.Id, creationOptions.Response.RelyingParty.ID)
	// 	assert.Equal(t, creationOptions.Response.AuthenticatorSelection.ResidentKey, protocol.ResidentKeyRequirementRequired)
	// 	assert.Equal(t, creationOptions.Response.AuthenticatorSelection.UserVerification, protocol.VerificationRequired)
	// 	assert.True(t, *creationOptions.Response.AuthenticatorSelection.RequireResidentKey)
	// }
}
