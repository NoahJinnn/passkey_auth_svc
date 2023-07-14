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

type itemTableSuite struct {
	Suite
}

func TestItemTableSuite(t *testing.T) {
	suite.Run(t, new(itemTableSuite))
}

func (s *itemTableSuite) TestItemTableHandler_ListByUserId() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/item_table")
	s.Require().NoError(err)

	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodGet, "/networth/item_tables", nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusOK, rec.Code) {
		assets := []ent.ItemTable{}
		err := json.Unmarshal(rec.Body.Bytes(), &assets)
		s.NoError(err)
		s.Equal(userId, assets[0].UserID.String())
	}
}

func (s *itemTableSuite) TestItemTableHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/item_table")
	s.Require().NoError(err)
	tests := []struct {
		name               string
		body               string
		expectedStatusCode int
		expectedAsset      *ent.ItemTable
	}{
		{
			name: "success",
			body: `{
				"sheet": 1,
				"section": 1,
				"description": "Test item table",
				"category": "asset"
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: &ent.ItemTable{
				Sheet:       1,
				Section:     1,
				Description: "Test item table",
				Category:    "asset",
			},
		},
		{
			name: "failed validation of category",
			body: `{
				"sheet": 1,
				"section": 1,
				"description": "Test item table",
				"category": "commercial"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing field",
			body: `{
				"sheet": 1,
				"section": 1,
				"description": "Test item table with missing field"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"
			token, err := s.sessionManager.GenerateJWT(userId)
			s.Require().NoError(err)
			cookie, err := s.sessionManager.GenerateCookie(token)
			s.NoError(err)

			req := httptest.NewRequest(http.MethodPost, "/networth/item_tables/item_table", strings.NewReader(tt.body))
			req.AddCookie(cookie)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			s.e.ServeHTTP(rec, req)
			s.Equal(tt.expectedStatusCode, rec.Code)
			if rec.Code == 200 {
				assets, err := s.app.GetItemTableSvc().ListByUser(ctx, uuid.FromStringOrNil(userId))
				s.NoError(err)
				s.Equal(2, len(assets)) // 1st from fixture, 2nd from create
				s.Equal(tt.expectedAsset.Sheet, assets[1].Sheet)
				s.Equal(tt.expectedAsset.Section, assets[1].Section)
				s.Equal(tt.expectedAsset.Description, assets[1].Description)
				s.Equal(tt.expectedAsset.Category, assets[1].Category)
				s.app.GetItemTableSvc().Delete(ctx, uuid.FromStringOrNil(userId), assets[1].ID)
			}
		})
	}
}
