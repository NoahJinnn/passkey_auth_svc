package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hellohq/hqservice/internal/http/validator"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/handlers"
	test "github.com/hellohq/hqservice/ms/auth/test/mock/app"
	testRepo "github.com/hellohq/hqservice/ms/auth/test/mock/dal"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestEmailSuite(t *testing.T) {
	suite.Run(t, new(emailSuite))
}

type emailSuite struct {
	Suite
}

func TestEmailHandler_ListByUser(t *testing.T) {
	uId1, _ := uuid.NewV4()
	uId2, _ := uuid.NewV4()

	tests := []struct {
		name          string
		userId        uuid.UUID
		data          []*ent.Email
		expectedCount int
	}{
		{
			name:   "should return all user assigned email addresses",
			userId: uId1,
			data: []*ent.Email{
				{
					UserID:  &uId1,
					Address: "test1@gmail.com",
				},
				{
					UserID:  &uId1,
					Address: "test2@gmail.com",
				},
				{
					UserID:  &uId2,
					Address: "test1@gmail.com",
				},
			},
			expectedCount: 2,
		},
		{
			name:   "should return empty list when user has no assigned email addresses",
			userId: uId2,
			data: []*ent.Email{
				{
					UserID:  &uId1,
					Address: "test1@gmail.com",
				},
				{
					UserID:  &uId1,
					Address: "test2@gmail.com",
				},
			},
			expectedCount: 0,
		},
	}

	var emails []*dto.EmailResponse
	for _, currentTest := range tests {
		e := echo.New()
		e.Validator = validator.NewCustomValidator()
		req := httptest.NewRequest(http.MethodGet, "/emails", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		// Mock JWT token
		c := e.NewContext(req, rec)
		token := jwt.New()
		err := token.Set(jwt.SubjectKey, currentTest.userId.String())
		require.NoError(t, err)
		c.Set("session", token)
		repo := testRepo.NewRepo(nil, nil, nil, nil, nil, currentTest.data, nil)
		appl := test.NewApp(nil, nil, &defaultCfg, repo)

		handler := handlers.NewEmailHandler(&handlers.HttpDeps{
			Appl:      appl,
			Cfg:       &defaultCfg,
			SharedCfg: &sharedCfg,
		}, &sessionManager{})

		if assert.NoError(t, handler.ListByUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &emails))
			assert.Equal(t, currentTest.expectedCount, len(emails))
		}
	}
}

func (s *emailSuite) TestEmailHandler_SetPrimaryEmail() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	err := s.LoadFixtures("../../../../test/fixtures/email")
	s.Require().NoError(err)

	newPrimaryEmailId := uuid.FromStringOrNil("8bb4c8a7-a3e6-48bb-b54f-20e3b485ab33")
	userId := "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"

	token, err := s.sessionManager.GenerateJWT(userId)
	s.Require().NoError(err)
	cookie, err := s.sessionManager.GenerateCookie(token)
	s.NoError(err)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/emails/%s/set_primary", newPrimaryEmailId), nil)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	if s.Equal(http.StatusNoContent, rec.Code) {
		email, err := s.repo.GetEmailRepo().GetById(ctx, uuid.FromStringOrNil(userId))
		s.Require().NoError(err)
		s.Equal(newPrimaryEmailId, email.ID)
	}
}

func TestEmailHandler_Delete(t *testing.T) {
	uId, _ := uuid.NewV4()
	emailId1, _ := uuid.NewV4()
	emailId2, _ := uuid.NewV4()
	testUsers := []*ent.User{
		{
			ID: uId,
		},
	}
	testEmails := []*ent.Email{
		{
			ID:      emailId1,
			Address: "test1@gmail.com",
		},
		{
			ID:      emailId2,
			Address: "test2@gmail.com",
		},
	}
	testUsers[0].Edges.Emails = testEmails

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/emails/:id")
	c.SetParamNames("id")
	c.SetParamValues(emailId1.String())
	token := jwt.New()
	err := token.Set(jwt.SubjectKey, uId.String())
	require.NoError(t, err)
	c.Set("session", token)

	repo := testRepo.NewRepo(nil, testUsers, nil, nil, nil, testEmails, nil)
	appl := test.NewApp(nil, nil, &defaultCfg, repo)
	handler := handlers.NewEmailHandler(&handlers.HttpDeps{
		Appl:      appl,
		Cfg:       &defaultCfg,
		SharedCfg: &sharedCfg,
	}, &sessionManager{})

	if assert.NoError(t, handler.Delete(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
