package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/internal/db"
	"github.com/NoahJinnn/passkey_auth_svc/internal/db/pgsql"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/session"
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	server "github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http/handlers"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/mail"
	testDal "github.com/NoahJinnn/passkey_auth_svc/ms/auth/test/mock/dal"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/suite"
)

var (
	ctx        = context.Background()
	defaultCfg = config.Config{
		Webauthn: config.WebauthnSettings{
			RelyingParty: config.RelyingParty{
				Id:          "localhost",
				DisplayName: "Test Relying Party",
				Icon:        "",
				Origins:     []string{"http://localhost:8080"},
			},
			Timeout: 60000,
		},
	}
	sharedCfg = sharedconfig.Shared{
		Session: sharedconfig.Session{
			Lifespan: "1h",
			Cookie: sharedconfig.Cookie{
				HttpOnly: true,
				SameSite: "strict",
				Secure:   true,
			},
			EnableAuthTokenHeader: true,
		},
		Secrets: sharedconfig.Secrets{
			Keys: []string{"needsToBeAtLeast16Test"},
		},
	}
	userId = "ec4ef049-5b88-4321-a173-21b0eff06a04"
	uId, _ = uuid.FromString(userId)
	emails = []*ent.Email{
		{
			ID:      uId,
			UserID:  &uId,
			Address: "john.doe@example.com",
		},
	}
	users = []*ent.User{
		func() *ent.User {
			user := &ent.User{
				ID:        uId,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			primE := &ent.PrimaryEmail{ID: uId, UserID: &uId}
			primE.Edges.Email = emails[0]
			user.Edges.PrimaryEmail = primE
			user.Edges.Emails = emails
			return user
		}(),
	}
)

type sessionManager struct{}

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

type Suite struct {
	suite.Suite
	repo           *dal.AuthRepo
	app            *app.App
	testDb         *testDal.TestDB
	srv            *handlers.HttpDeps
	sessionManager *session.Manager
	e              *echo.Echo
}

func (s *Suite) SetupSuite() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	dialect := "postgres"
	testDb, err := testDal.StartDB("integration_test_auth", dialect)
	s.NoError(err)
	pgClient := pgsql.NewPgEnt(testDb.DatabaseUrl)

	dbClient := &db.Db{PgEnt: pgClient}
	repo := dal.New(dbClient)
	jwkRepo := session.NewJwkRepo(dbClient)
	jwkManager, err := session.NewDefaultManager(sharedCfg.Secrets.Keys, jwkRepo)
	s.NoError(err)
	sessionManager, err := session.NewManager(jwkManager, sharedCfg.Session)
	s.NoError(err)

	s.repo = repo
	s.sessionManager = sessionManager
	s.testDb = testDb
	s.app = app.New(mailer{}, renderer, &defaultCfg, repo)
	s.srv = &handlers.HttpDeps{
		Appl:      s.app,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}
	srv, err := server.NewServer(s.app, sessionManager, &sharedCfg, &defaultCfg)
	s.NoError(err)

	s.e = srv
}

func (s *Suite) TearDownSuite() {
	if s.testDb != nil {
		s.NoError(testDal.PurgeDB(s.testDb))
	}
}

// LoadFixtures loads predefined data from the path in the database.
func (s *Suite) LoadFixtures(path string) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(s.testDb.DbCon),
		testfixtures.Dialect(s.testDb.Dialect),
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
