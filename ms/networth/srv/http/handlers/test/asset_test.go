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

	req := httptest.NewRequest(http.MethodGet, "/networth/assets", nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusOK, rec.Code) {
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
	description := "Test asset"
	tests := []struct {
		name               string
		body               string
		expectedStatusCode int
		expectedAsset      ent.Asset
	}{
		{
			name: "success",
			body: `{
				"sheet": 1,
				"section": 1,
				"type": "debit",
				"provider_name": "saltedge",
				"currency": "USD",
				"value": 1.123456,
				"description": "Test asset"
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: ent.Asset{
				Sheet:        1,
				Section:      1,
				Type:         "debit",
				ProviderName: "saltedge",
				Currency:     "USD",
				Value:        1.123456,
				Description:  description,
			},
		},
		{
			name: "missing value",
			body: `{
				"sheet": 1,
				"section": 1,
				"type": "debit",
				"provider_name": "saltedge",
				Currency: "USD",
				"description": "Test asset"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset: ent.Asset{
				Sheet:        1,
				Section:      1,
				Type:         "debit",
				ProviderName: "saltedge",
				Currency:     "USD",
				Value:        1.123456,
				Description:  description,
			},
		},
		{
			name: "default value",
			body: `{
				"value": 1.123456
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: ent.Asset{
				Sheet:        0,
				Section:      0,
				Type:         "",
				ProviderName: "",
				Currency:     "",
				Value:        1.123456,
				Description:  "",
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
			token, err := s.sessionManager.GenerateJWT(userId)
			s.Require().NoError(err)
			cookie, err := s.sessionManager.GenerateCookie(token)
			s.NoError(err)

			req := httptest.NewRequest(http.MethodPost, "/networth/assets/asset", strings.NewReader(tt.body))
			req.AddCookie(cookie)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			s.e.ServeHTTP(rec, req)
			s.Equal(tt.expectedStatusCode, rec.Code)
			if rec.Code == 200 {
				assets, err := s.app.GetAssetSvc().ListByUser(ctx, uuid.FromStringOrNil(userId))
				s.NoError(err)
				s.Equal(2, len(assets)) // 1st from fixture, 2nd from create
				s.Equal(tt.expectedAsset.Value, assets[1].Value)
				s.Equal(tt.expectedAsset.Sheet, assets[1].Sheet)
				s.Equal(tt.expectedAsset.Section, assets[1].Section)
				s.Equal(tt.expectedAsset.Type, assets[1].Type)
				s.Equal(tt.expectedAsset.ProviderName, assets[1].ProviderName)
				s.Equal(tt.expectedAsset.Currency, assets[1].Currency)
				s.Equal(tt.expectedAsset.Description, assets[1].Description)
				s.app.GetAssetSvc().Delete(ctx, uuid.FromStringOrNil(userId), assets[1].ID)
			}
		})
	}

	// body := `{
	// 	"sheet": 1,
	// 	"section": 1,
	// 	"type": "debit",
	// 	"provider_name": "saltedge",
	// 	"currency": "USD",
	// 	"value": 1.123456,
	// 	"description": "Test asset"
	// }`

}
