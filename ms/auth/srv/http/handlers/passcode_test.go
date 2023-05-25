package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/validator"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/mail"
	test "github.com/hellohq/hqservice/ms/auth/test/mock/app"
	testRepo "github.com/hellohq/hqservice/ms/auth/test/mock/dal"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mailer struct{}

var renderer, _ = mail.NewRenderer()

func passcodes() []*ent.Passcode {
	now := time.Now()
	return []*ent.Passcode{{
		ID:        uuid.FromStringOrNil("08ee61aa-0946-4ecf-a8bd-e14c604329e2"),
		UserID:    uuid.FromStringOrNil(userId),
		TTL:       300,
		Code:      "$2a$12$gBPH9jnbXFmwAGwZMSzYkeXx7oOTElzhvHfiDgj.D7G8q4znvHpMK",
		CreatedAt: now,
		UpdatedAt: now,
	}}
}

func (m mailer) Send(email []string, subject string, body string) error {
	return nil
}

func TestPasscodeHandler_Init(t *testing.T) {
	appl := test.NewApp(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, nil, emails, nil))
	srv := &HttpDeps{
		appl,
		&defaultCfg,
		&sharedCfg,
	}
	passcodeHandler := NewPasscodeHandler(srv, sessionManager{})

	body := dto.PasscodeInitRequest{
		UserId: userId,
	}
	bodyJson, err := json.Marshal(body)
	require.NoError(t, err)

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	req := httptest.NewRequest(http.MethodPost, "/passcode/login/initialize", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, passcodeHandler.Init(c)) {
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	}
}

func TestPasscodeHandler_Finish(t *testing.T) {
	appl := test.NewApp(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, passcodes(), emails, nil))
	srv := &HttpDeps{
		appl,
		&defaultCfg,
		&sharedCfg,
	}
	passcodeHandler := NewPasscodeHandler(srv, sessionManager{})

	body := dto.PasscodeFinishRequest{
		Id:   "08ee61aa-0946-4ecf-a8bd-e14c604329e2",
		Code: "123456",
	}
	bodyJson, err := json.Marshal(body)
	require.NoError(t, err)

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, passcodeHandler.Finish(c)) {
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	}
}
