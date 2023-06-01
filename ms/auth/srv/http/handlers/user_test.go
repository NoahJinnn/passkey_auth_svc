package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/internal/http/validator"
	"github.com/hellohq/hqservice/internal/pgsql"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	test "github.com/hellohq/hqservice/ms/auth/test/mock/dal"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

type userSuite struct {
	suite.Suite
	repo    *dal.AuthRepo
	app     *app.App
	db      *test.TestDB
	srv     *HttpDeps
	handler *UserHandler
	echo    *echo.Echo
}

func (s *userSuite) SetupSuite() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	dialect := "postgres"
	db, err := test.StartDB("integration_test", dialect)
	s.NoError(err)
	entClient := pgsql.CreateEntClient(ctx, db.DatabaseUrl)
	repo := dal.New(entClient)

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	s.echo = e

	s.repo = repo
	s.db = db
	s.app = app.New(nil, nil, &defaultCfg, repo)
	s.srv = &HttpDeps{
		s.app,
		&defaultCfg,
		&sharedCfg,
	}
	s.handler = NewUserHandler(s.srv, &sessionManager{})
}

func (s *userSuite) TearDownSuite() {
	if s.db != nil {
		s.NoError(test.PurgeDB(s.db))
	}
}

// LoadFixtures loads predefined data from the path in the database.
func (s *userSuite) LoadFixtures(path string) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(s.db.DbCon),
		testfixtures.Dialect(s.db.Dialect),
		testfixtures.Directory(path),
		testfixtures.SkipResetSequences(),
	)
	if err != nil {
		return fmt.Errorf("could not create testfixtures: %w", err)
	}

	err = fixtures.Load()
	if err != nil {
		return fmt.Errorf("could not load fixtures: %w", err)
	}

	return nil
}

func (s *userSuite) TestUserHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	body := UserCreateBody{Email: "noah.jin@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := s.echo.NewContext(req, rec)

	if s.NoError(s.handler.Create(c)) && s.Equal(http.StatusOK, rec.Code) {
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

	body := UserCreateBody{Email: "JANE.DOE@EXAMPLE.COM"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := s.echo.NewContext(req, rec)

	if s.NoError(s.handler.Create(c)) && s.Equal(http.StatusOK, rec.Code) {
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
	err := s.LoadFixtures("../../../test/fixtures")
	s.Require().NoError(err)

	body := UserCreateBody{Email: "john.doe@example.com"}
	bodyJson, err := json.Marshal(body)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := s.echo.NewContext(req, rec)
	s.handler.Create(c)
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

	c := s.echo.NewContext(req, rec)
	s.handler.Create(c)

	if s.Equal(http.StatusBadRequest, rec.Code) {
		httpError := errorhandler.HTTPError{}
		err := json.Unmarshal(rec.Body.Bytes(), &httpError)
		s.NoError(err)
		s.Equal(http.StatusBadRequest, httpError.Code)
	}
}
