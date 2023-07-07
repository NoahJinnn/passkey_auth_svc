package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/stretchr/testify/suite"
)

func TestPasscodeSuite(t *testing.T) {
	suite.Run(t, new(passcodeSuite))
}

type passcodeSuite struct {
	Suite
}

func (s *passcodeSuite) TestPasscodeHandler_Init() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/passcode")
	s.Require().NoError(err)

	tests := []struct {
		name               string
		body               dto.PasscodeInitRequest
		expectedStatusCode int
	}{
		{
			name: "init with user id",
			body: dto.PasscodeInitRequest{
				UserId: "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "unknown user id",
			body: dto.PasscodeInitRequest{
				UserId: "04603148-036d-403b-bf34-cfe237974ef9",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			bodyJson, err := json.Marshal(tt.body)
			s.NoError(err)

			req := httptest.NewRequest(http.MethodPost, "/passcode/login/initialize", bytes.NewReader(bodyJson))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			s.e.ServeHTTP(rec, req)

			s.Equal(tt.expectedStatusCode, rec.Code)
		})
	}
}

// func TestPasscodeHandler_Finish(t *testing.T) {
// 	appl := app.New(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, passcodes(), emails, nil))
// 	srv := &handlers.HttpDeps{
// 		Appl:      appl,
// 		Cfg:       &defaultCfg,
// 		SharedCfg: &sharedCfg,
// 	}
// 	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager{})

// 	body := dto.PasscodeFinishRequest{
// 		Id:   "08ee61aa-0946-4ecf-a8bd-e14c604329e2",
// 		Code: "123456",
// 	}
// 	bodyJson, err := json.Marshal(body)
// 	require.NoError(t, err)

// 	e := echo.New()
// 	e.Validator = validator.NewCustomValidator()
// 	req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
// 	req.Header.Set("Content-Type", "application/json")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, passcodeHandler.Finish(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
// 	}
// }

// func TestPasscodeHandler_Finish_WrongCode(t *testing.T) {
// 	appl := app.New(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, passcodes(), emails, nil))
// 	srv := &handlers.HttpDeps{
// 		Appl:      appl,
// 		Cfg:       &defaultCfg,
// 		SharedCfg: &sharedCfg,
// 	}
// 	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager{})

// 	body := dto.PasscodeFinishRequest{
// 		Id:   "08ee61aa-0946-4ecf-a8bd-e14c604329e2",
// 		Code: "012345",
// 	}
// 	bodyJson, err := json.Marshal(body)
// 	require.NoError(t, err)

// 	e := echo.New()
// 	e.Validator = validator.NewCustomValidator()
// 	req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
// 	req.Header.Set("Content-Type", "application/json")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	err = passcodeHandler.Finish(c)
// 	if assert.Error(t, err) {
// 		httpError := errorhandler.ToHttpError(err)
// 		assert.Equal(t, http.StatusUnauthorized, httpError.Code)
// 	}
// }

// func TestPasscodeHandler_Finish_WrongCode_3_Times(t *testing.T) {
// 	appl := app.New(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, passcodes(), emails, nil))
// 	srv := &handlers.HttpDeps{
// 		Appl:      appl,
// 		Cfg:       &defaultCfg,
// 		SharedCfg: &sharedCfg,
// 	}
// 	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager{})

// 	body := dto.PasscodeFinishRequest{
// 		Id:   "08ee61aa-0946-4ecf-a8bd-e14c604329e2",
// 		Code: "012345",
// 	}
// 	bodyJson, err := json.Marshal(body)
// 	require.NoError(t, err)

// 	e := echo.New()
// 	e.Validator = validator.NewCustomValidator()
// 	for i := 0; i < 3; i++ {
// 		req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
// 		req.Header.Set("Content-Type", "application/json")
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)

// 		err = passcodeHandler.Finish(c)
// 		if i < 2 {
// 			if assert.Error(t, err) {
// 				httpError := errorhandler.ToHttpError(err)
// 				assert.Equal(t, http.StatusUnauthorized, httpError.Code)
// 			}
// 		} else {
// 			if assert.Error(t, err) {
// 				httpError := errorhandler.ToHttpError(err)
// 				assert.Equal(t, http.StatusGone, httpError.Code)
// 			}
// 		}
// 	}
// }

// func TestPasscodeHandler_Finish_WrongId(t *testing.T) {
// 	appl := app.New(&mailer{}, renderer, &defaultCfg, testRepo.NewRepo(nil, users, nil, nil, passcodes(), emails, nil))
// 	srv := &handlers.HttpDeps{
// 		Appl:      appl,
// 		Cfg:       &defaultCfg,
// 		SharedCfg: &sharedCfg,
// 	}
// 	passcodeHandler := handlers.NewPasscodeHandler(srv, sessionManager{})

// 	body := dto.PasscodeFinishRequest{
// 		Id:   "1bc9a074-577d-497e-87da-8eaf50f32a26",
// 		Code: "123456",
// 	}
// 	bodyJson, err := json.Marshal(body)
// 	require.NoError(t, err)

// 	e := echo.New()
// 	e.Validator = validator.NewCustomValidator()
// 	req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
// 	req.Header.Set("Content-Type", "application/json")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	err = passcodeHandler.Finish(c)
// 	if assert.Error(t, err) {
// 		httpError := errorhandler.ToHttpError(err)
// 		assert.Equal(t, http.StatusUnauthorized, httpError.Code)
// 	}
// }
