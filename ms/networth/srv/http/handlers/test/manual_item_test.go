package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/stretchr/testify/suite"
)

type manualItemSuite struct {
	Suite
}

func TestManualItemSuite(t *testing.T) {
	// t.Parallel()
	suite.Run(t, new(manualItemSuite))
}

func (s *manualItemSuite) TestManualItemHandler_Create() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/manual_item")
	s.Require().NoError(err)
	tests := []struct {
		name               string
		body               string
		expectedStatusCode int
		expectedAsset      *ent.ManualItem
	}{
		{
			name: "success",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusOK,
			expectedAsset: &ent.ManualItem{
				ItemTableID:  "12345",
				Category:     "asset",
				Type:         "debit",
				Description:  "TCB",
				Value:        2000.00,
				ProviderName: "manual",
			},
		},
		{
			name: "failed validation of category",
			body: `{
				"item_table_id": "12345",
				"category": "commercial",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "failed validation of provider name",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "test"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing item_table_id field",
			body: `{
				"category": "asset",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing category field",
			body: `{
				"item_table_id": "12345",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing type field",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"description": "TCB",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing description field",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"type": "debit",
				"value": 2000.00,
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing value field",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"type": "debit",
				"description": "TCB",
				"provider_name": "manual"
			}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedAsset:      nil,
		},
		{
			name: "missing provider_name field",
			body: `{
				"item_table_id": "12345",
				"category": "asset",
				"type": "debit",
				"description": "TCB",
				"value": 2000.00
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

			req := httptest.NewRequest(http.MethodPost, "/networth/manual_items/manual_item", strings.NewReader(tt.body))
			req.AddCookie(cookie)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			s.e.ServeHTTP(rec, req)
			s.Equal(tt.expectedStatusCode, rec.Code)
			if rec.Code == 200 {
				assets, err := s.app.GetProviderSvc().AllManualItem(ctx, uuid.FromStringOrNil(userId))
				s.NoError(err)
				s.Equal(1, len(assets))
				s.Equal(tt.expectedAsset.ItemTableID, assets[0].ItemTableID)
				s.Equal(tt.expectedAsset.Category, assets[0].Category)
				s.Equal(tt.expectedAsset.Type, assets[0].Type)
				s.Equal(tt.expectedAsset.Description, assets[0].Description)
				s.Equal(tt.expectedAsset.Value, assets[0].Value)
				s.Equal(tt.expectedAsset.ProviderName, assets[0].ProviderName)
				s.app.GetProviderSvc().DeleteManualItem(ctx, uuid.FromStringOrNil(userId), assets[0].ID)
			}
		})
	}
}
