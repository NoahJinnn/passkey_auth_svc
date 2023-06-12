package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/handlers"
	"github.com/stretchr/testify/suite"
)

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

type userSuite struct {
	Suite
}

func (s *userSuite) TestUserHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	body := handlers.UserCreateBody{Email: "noah.jin@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusOK, rec.Code) {
		user := dto.CreateUserResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		s.NoError(err)
		s.False(user.ID.IsNil())

		count, err := s.repo.GetUserRepo().Count(ctx, user.ID)
		s.NoError(err)
		s.Equal(1, count)

		email, err := s.repo.GetEmailRepo().GetByAddress(ctx, body.Email)
		s.NoError(err)
		s.NotNil(email)
	}
}

func (s *userSuite) TestUserHandler_Create_CaseInsensitive() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	body := handlers.UserCreateBody{Email: "JANE.DOE@EXAMPLE.COM"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusOK, rec.Code) {
		user := dto.CreateUserResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		s.NoError(err)
		s.False(user.ID.IsNil())

		count, err := s.repo.GetUserRepo().Count(ctx, user.ID)
		s.NoError(err)
		s.Equal(1, count)

		email, err := s.repo.GetEmailRepo().GetByAddress(ctx, strings.ToLower(body.Email))
		s.NoError(err)
		s.NotNil(email)
	}
}

func (s *userSuite) TestUserHandler_Create_UserExists() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/user")
	s.Require().NoError(err)

	body := handlers.UserCreateBody{Email: "john.doe@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusConflict, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.NoError(err)
		s.Equal(http.StatusConflict, httpError.Code)
	}
}

func (s *userSuite) TestUserHandler_Create_InvalidEmail() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"email": 123"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusBadRequest, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.NoError(err)
		s.Equal(http.StatusBadRequest, httpError.Code)
	}
}

func (s *userSuite) TestUserHandler_Create_EmailMissing() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"bogus": 123}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusBadRequest, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.NoError(err)
		s.Equal(http.StatusBadRequest, httpError.Code)
	}
}

func (s *userSuite) TestUserHandler_Get() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/user")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"

	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userId), nil)
	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusOK, rec.Code) {
		s.Equal(rec.Code, http.StatusOK)
		user := ent.User{}
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		s.NoError(err)
		s.Equal(userId, user.ID.String())
		s.Equal(len(user.Edges.WebauthnCredentials), 0)
	}
}

func (s *userSuite) TestUserHandler_GetUserWithWebAuthnCredential() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/user_with_webauthn_credential")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"

	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userId), nil)
	rec := httptest.NewRecorder()
	req.AddCookie(cookie)

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusOK, rec.Code) {
		s.Equal(rec.Code, http.StatusOK)
		var user struct {
			ID                  string                   `json:"id"`
			WebauthnCredentials []ent.WebauthnCredential `json:"webauthn_credentials"`
		}
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		s.Require().NoError(err)
		s.Equal(userId, user.ID)
		s.Equal(1, len(user.WebauthnCredentials))
	}
}

func (s *userSuite) TestUserHandler_Get_InvalidUserId() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"

	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodGet, "/users/invalidUserId", nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusForbidden, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.Require().NoError(err)
		s.Equal(http.StatusForbidden, httpError.Code)
	}
}

func (s *userSuite) TestUserHandler_GetUserIdByEmail_InvalidEmail() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"email": "123"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusBadRequest, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.Require().NoError(err)
		s.Equal(http.StatusBadRequest, httpError.Code)
	}
}

func (s *userSuite) TestUserHandler_GetUserIdByEmail_InvalidJson() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`"email": "123}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
}

func (s *userSuite) TestUserHandler_GetUserIdByEmail_UserNotFound() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"email": "unknownAddress@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	s.Equal(http.StatusNotFound, rec.Code)
}

func (s *userSuite) TestUserHandler_Logout() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	userId, _ := uuid.NewV4()

	token, err := s.sessionManager.GenerateJWT(userId.String())
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)

	if s.Equal(http.StatusNoContent, rec.Code) {
		cookie := rec.Header().Get("Set-Cookie")
		s.NotEmpty(cookie)

		split := strings.Split(cookie, ";")
		s.Equal("Max-Age=0", strings.TrimSpace(split[2]))
	}
}
