package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/dal/test"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	suite.Suite
	repo *dal.Repo
	db   *test.TestDB
	srv  *HttpDeps
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (s *userSuite) SetupSuite() {
	if testing.Short() {
		return
	}
	dialect := "postgres"
	db, err := test.StartDB("user_test", dialect)
	s.NoError(err)
	repo, err := dal.New(ctx, db.DatabaseUrl)
	s.NoError(err)

	s.repo = repo
	s.db = db
	s.srv = &HttpDeps{}
}

func (s *userSuite) TearDownSuite() {
	if s.db != nil {
		s.NoError(test.PurgeDB(s.db))
	}
}

func (s *userSuite) TestUserHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	e := echo.New()
	e.Validator = dto.NewCustomValidator()

	body := UserCreateBody{Email: "jane.doe@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewUserHandler(s.srv, &sessionManager{})

	if s.NoError(handler.Create(c)) {
		user := UserCreateBody{}
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		s.NoError(err)
		s.NotEmpty(user.Email)

		// count, err := s.repo.GetUserRepo().Count(uuid.Nil, "")
		// s.NoError(err)
		// s.Equal(1, count)

		// email, err := s.repo.GetEmailRepo().FindByAddress(body.Email)
		// s.NoError(err)
		// s.NotNil(email)
	}
}
