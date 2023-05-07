package handlers

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/sharedDal"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	test "github.com/hellohq/hqservice/ms/auth/test/dal"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/suite"
)

var ctx = context.Background()
var defaultConfig = config.Config{
	Webauthn: config.WebauthnSettings{
		RelyingParty: config.RelyingParty{
			Id:          "localhost",
			DisplayName: "Test Relying Party",
			Icon:        "",
			Origin:      "http://localhost:8080",
		},
		Timeout: 60000,
	},
}

type sessionManager struct {
}

var userId = "ec4ef049-5b88-4321-a173-21b0eff06a04"

func (s sessionManager) GenerateJWT(uuid string) (string, error) {
	return userId, nil
}

func (s sessionManager) GenerateCookie(token string) (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hanko",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}, nil
}

func (s sessionManager) DeleteCookie() (*http.Cookie, error) {
	return &http.Cookie{
		Name:     "hanko",
		Value:    "",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}, nil
}

func (s sessionManager) Verify(token string) (jwt.Token, error) {
	return nil, nil
}

var uId, _ = uuid.FromString(userId)

var emails = []*ent.Email{
	{
		ID:      uId,
		UserID:  uId,
		Address: "john.doe@example.com",
	},
}

var users = []*ent.User{
	func() *ent.User {

		user := &ent.User{
			ID:        uId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		user.Edges.PrimaryEmail = &ent.PrimaryEmail{ID: uId, UserID: uId}
		user.Edges.Emails = emails
		return user
	}(),
}

type integrationSuite struct {
	suite.Suite
	repo *dal.Repo
	app  app.Appl
	db   *test.TestDB
	srv  *HttpDeps
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(integrationSuite))
}

func (s *integrationSuite) SetupSuite() {
	if testing.Short() {
		return
	}
	dialect := "postgres"
	db, err := test.StartDB("integration_test", dialect)
	s.NoError(err)
	entClient := sharedDal.CreateEntClient(ctx, db.DatabaseUrl)
	repo := dal.New(entClient)

	s.repo = repo
	s.db = db
	s.app = app.New(&defaultConfig, repo)
	s.srv = &HttpDeps{
		s.app,
		&defaultConfig,
	}
}

func (s *integrationSuite) TearDownSuite() {
	if s.db != nil {
		s.NoError(test.PurgeDB(s.db))
	}
}
