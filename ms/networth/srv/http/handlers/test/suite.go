package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/hellohq/hqservice/internal/db"
	"github.com/hellohq/hqservice/internal/db/pgsql"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	server "github.com/hellohq/hqservice/ms/networth/srv/http"
	"github.com/hellohq/hqservice/ms/networth/srv/http/handlers"
	testDal "github.com/hellohq/hqservice/ms/networth/test/mock/dal"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	ctx = context.Background()
	// TODO: Test 3rd provider with test configuration
	defaultCfg = config.Config{
		SaltEdge: &config.SaltEdge{
			AppId:  "",
			Secret: "",
			PK:     "",
		},
		Finverse: &config.Finverse{
			AppId:       "",
			Secret:      "",
			ClientID:    "",
			RedirectURI: "",
		},
		Plaid: &config.Plaid{
			ClientID: "",
			Secret:   "",
			Env:      "dev",
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
)

type Suite struct {
	suite.Suite
	repo           *dal.NwRepo
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
	testDb, err := testDal.StartDB("integration_test_nw", dialect)
	s.NoError(err)
	pgClient := pgsql.NewPgClient(testDb.DatabaseUrl)

	dbClient := &db.Db{PgClient: pgClient}
	repo := dal.New(dbClient)
	jwkRepo := session.NewJwkRepo(dbClient)
	jwkManager, err := session.NewDefaultManager(sharedCfg.Secrets.Keys, jwkRepo)
	s.NoError(err)
	sessionManager, err := session.NewManager(jwkManager, sharedCfg.Session)
	s.NoError(err)

	s.repo = repo
	s.sessionManager = sessionManager
	s.testDb = testDb
	s.app = app.New(&defaultCfg, repo)
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
		fmt.Println("Tearing down test database")
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
