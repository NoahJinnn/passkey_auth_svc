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

func (s *assetSuite) TestAssetTableHandler_ListByUserId() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/asset_table")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodGet, "/networth/asset_tables", nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusOK, rec.Code) {
		assets := []ent.AssetTable{}
		err := json.Unmarshal(rec.Body.Bytes(), &assets)
		s.NoError(err)
		s.Equal(userId, assets[0].UserID.String())
	}
}

func (s *assetSuite) TestAssetTableHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/asset_table")
	s.Require().NoError(err)
	description := "Test asset table"
	tests := []struct {
		name               string
		body               string
		expectedStatusCode int
		expectedAsset      ent.AssetTable
	}{
		{
			name: "success",
			body: `{
				"sheet": 1,
				"section": 1,
				"description": "Test asset table"
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: ent.AssetTable{
				Sheet:       1,
				Section:     1,
				Description: description,
			},
		},
		{
			name: "default value",
			body: `{
				"value": 1.123456
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: ent.AssetTable{
				Sheet:       0,
				Section:     0,
				Description: "",
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

			req := httptest.NewRequest(http.MethodPost, "/networth/asset_tables/asset_table", strings.NewReader(tt.body))
			req.AddCookie(cookie)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			s.e.ServeHTTP(rec, req)
			s.Equal(tt.expectedStatusCode, rec.Code)
			if rec.Code == 200 {
				assets, err := s.app.GetAssetSvc().ListByUser(ctx, uuid.FromStringOrNil(userId))
				s.NoError(err)
				s.Equal(2, len(assets)) // 1st from fixture, 2nd from create
				s.Equal(tt.expectedAsset.Sheet, assets[1].Sheet)
				s.Equal(tt.expectedAsset.Section, assets[1].Section)
				s.Equal(tt.expectedAsset.Description, assets[1].Description)
				s.app.GetAssetSvc().Delete(ctx, uuid.FromStringOrNil(userId), assets[1].ID)
			}
		})
	}
}
