package handlers

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/dal/test"
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

func (s sessionManager) GenerateJWT(uuid uuid.UUID) (string, error) {
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

var credentials = []ent.WebauthnCredential{
	func() ent.WebauthnCredential {
		uId, _ := uuid.FromString(userId)
		aaguid, _ := uuid.FromString("adce0002-35bc-c60a-648b-0b25f1f05503")
		return ent.WebauthnCredential{
			ID:              "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjH",
			UserID:          uId,
			PublicKey:       "pQECAyYgASFYIPG9WtGAri-mevonFPH4p-lI3JBS29zjuvKvJmaP4_mRIlggOjHw31sdAGvE35vmRep-aPcbAAlbuc0KHxQ9u6zcHog",
			AttestationType: "none",
			Aaguid:          aaguid,
			SignCount:       1650958750,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		}
	}(),
	func() ent.WebauthnCredential {
		uId, _ := uuid.FromString(userId)
		aaguid, _ := uuid.FromString("adce0002-35bc-c60a-648b-0b25f1f05503")
		return ent.WebauthnCredential{
			ID:              "AaFdkcD4SuPjF-jwUoRwH8-ZHuY5RW46fsZmEvBX6RNKHaGtVzpATs06KQVheIOjYz-YneG4cmQOedzl0e0jF951ukx17Hl9jeGgWz5_DKZCO12p2-2LlzjK",
			UserID:          uId,
			PublicKey:       "pQECAyYgASFYIPG9WtGAri-mevonFPH4p-lI3JBS29zjuvKvJmaP4_mRIlggOjHw31sdAGvE35vmRep-aPcbAAlbuc0KHxQ9u6zcHoj",
			AttestationType: "none",
			Aaguid:          aaguid,
			SignCount:       1650958750,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		}
	}(),
}

var sessionData = []ent.WebauthnSessionData{
	func() ent.WebauthnSessionData {
		id, _ := uuid.NewV4()
		uId, _ := uuid.FromString(userId)
		return ent.WebauthnSessionData{
			ID:               id,
			Challenge:        "tOrNDCD2xQf4zFjEjwxaP8fOErP3zz08rMoTlJGtnKU",
			UserID:           uId,
			UserVerification: string(protocol.VerificationRequired),
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			Operation:        svcs.WebauthnOperationRegistration,
			// AllowedCredentials: nil,
		}
	}(),
	func() ent.WebauthnSessionData {
		id, _ := uuid.NewV4()
		return ent.WebauthnSessionData{
			ID:               id,
			Challenge:        "gKJKmh90vOpYO55oHpqaHX_oMCq4oTZt-D0b6teIzrE",
			UserID:           uuid.UUID{},
			UserVerification: string(protocol.VerificationRequired),
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			Operation:        svcs.WebauthnOperationAuthentication,
			// AllowedCredentials: nil,
		}
	}(),
}

var uId, _ = uuid.FromString(userId)

var emails = []ent.Email{
	{
		ID:      uId,
		Address: "john.doe@example.com",
		// PrimaryEmail: &ent.PrimaryEmail{ID: uId},
	},
}

var users = []ent.User{
	func() ent.User {

		return ent.User{
			ID:        uId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			// Emails:    emails,
		}
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
		s.repo = &dal.Repo{}
		s.app = app.New(&defaultConfig, s.repo)
		s.srv = &HttpDeps{
			s.app,
			&defaultConfig,
		}
		return
	}
	dialect := "postgres"
	db, err := test.StartDB("user_test", dialect)
	s.NoError(err)
	repo, err := dal.New(ctx, db.DatabaseUrl)
	s.NoError(err)

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
