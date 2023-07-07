package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/stretchr/testify/suite"
)

type assetSuite struct {
	Suite
}

func TestAssetSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(assetSuite))
}

func (s *assetSuite) TestAssetHandler_ListByUserId() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/asset")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodGet, "/networth/asset", nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusOK, rec.Code) {
		s.Equal(rec.Code, http.StatusOK)
		assets := []ent.Asset{}
		err := json.Unmarshal(rec.Body.Bytes(), &assets)
		s.NoError(err)
		s.Equal(userId, assets[0].UserID.String())
	}
}

func (s *assetSuite) TestAssetHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/asset")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	body := `{
		"sheet": 1,
		"section": 1,
		"type": "debit",
		"provider_name": "saltedge",
		"currency": "USD",
		"value": 1.123456789,
		"description": "Test asset"
	}`
	req := httptest.NewRequest(http.MethodPost, "/networth/asset", strings.NewReader(body))
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusOK, rec.Code) {
		s.Equal(rec.Code, http.StatusOK)
		assets, err := s.app.GetAssetSvc().ListByUser(ctx, uuid.FromStringOrNil(userId))
		s.NoError(err)
		s.Equal(2, len(assets)) // 1st from fixture, 2nd from create
		s.Equal(1.1235, assets[1].Value)
	}
}
