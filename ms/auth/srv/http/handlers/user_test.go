package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hellohq/hqservice/internal/http/validator"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
)

func (s *integrationSuite) TestUserHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	body := UserCreateBody{Email: "noah.jin@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewUserHandler(s.srv, &sessionManager{})

	if s.NoError(handler.Create(c)) {
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
