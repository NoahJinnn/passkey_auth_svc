package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

func TestPasscodeSuite(t *testing.T) {
	suite.Run(t, new(passcodeSuite))
}

type passcodeSuite struct {
	Suite
}

func (s *passcodeSuite) TestPasscodeHandler_Init() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/passcode")
	s.Require().NoError(err)

	emailId := "51b7c175-ceb6-45ba-aae6-0092221c1b84"
	unknownEmailId := "83618f24-2db8-4ea2-b370-ac8335f782d8"
	tests := []struct {
		name               string
		body               dto.PasscodeInitBody
		expectedStatusCode int
	}{
		{
			name: "with userID and emailID",
			body: dto.PasscodeInitBody{
				UserId:  "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5",
				EmailId: &emailId,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "with user id",
			body: dto.PasscodeInitBody{
				UserId: "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "unknown user id",
			body: dto.PasscodeInitBody{
				UserId: unknownEmailId,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "with unknown emailID",
			body: dto.PasscodeInitBody{
				UserId:  "b5dd5267-b462-48be-b70d-bcd6f1bbe7a5",
				EmailId: &unknownEmailId,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			bodyJson, err := json.Marshal(tt.body)
			s.NoError(err)

			req := httptest.NewRequest(http.MethodPost, "/passcode/login/initialize", bytes.NewReader(bodyJson))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			s.e.ServeHTTP(rec, req)

			s.Equal(tt.expectedStatusCode, rec.Code)
		})
	}
}

func (s *passcodeSuite) TestPasscodeHandler_Finish() {
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}
	err := s.LoadFixtures("../../../../test/fixtures/passcode")
	s.Require().NoError(err)

	hashedPasscode, err := bcrypt.GenerateFromPassword([]byte("123456"), 12)
	s.Require().NoError(err)

	passcode := ent.Passcode{
		UserID:   uuid.FromStringOrNil("b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"),
		EmailID:  uuid.FromStringOrNil("51b7c175-ceb6-45ba-aae6-0092221c1b84"),
		TTL:      300,
		Code:     string(hashedPasscode),
		TryCount: 0,
	}

	passcodeWithExpiredTimeout := ent.Passcode{
		UserID:   uuid.FromStringOrNil("b5dd5267-b462-48be-b70d-bcd6f1bbe7a5"),
		EmailID:  uuid.FromStringOrNil("51b7c175-ceb6-45ba-aae6-0092221c1b84"),
		TTL:      0,
		Code:     string(hashedPasscode),
		TryCount: 0,
	}

	tests := []struct {
		name               string
		passcodeId         string
		retryCount         int
		passcode           ent.Passcode
		code               string
		expectedStatusCode int
	}{
		{
			name:               "finish successful",
			passcodeId:         "a2383922-dea3-46c8-be17-85b267c0d135",
			passcode:           passcode,
			code:               "123456",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "with wrong code",
			passcodeId:         "a2383922-dea3-46c8-be17-85b267c0d135",
			passcode:           passcode,
			code:               "654321",
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "with wrong code 3 times",
			passcodeId:         "a2383922-dea3-46c8-be17-85b267c0d135",
			retryCount:         2,
			passcode:           passcode,
			code:               "654321",
			expectedStatusCode: http.StatusGone,
		},
		{
			name:               "after passcode expired",
			passcodeId:         "a2383922-dea3-46c8-be17-85b267c0d135",
			passcode:           passcodeWithExpiredTimeout,
			code:               "123456",
			expectedStatusCode: http.StatusRequestTimeout,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			pc, err := s.repo.GetPasscodeRepo().Create(ctx, &tt.passcode)
			s.Require().NoError(err)
			body := dto.PasscodeFinishBody{
				Id:   pc.ID.String(),
				Code: tt.code,
			}
			bodyJson, err := json.Marshal(body)
			s.Require().NoError(err)

			responseCode := 0
			for i := 0; i <= tt.retryCount; i++ {
				req := httptest.NewRequest(http.MethodPost, "/passcode/login/finalize", bytes.NewReader(bodyJson))
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()

				s.e.ServeHTTP(rec, req)
				responseCode = rec.Code
			}

			s.Equal(tt.expectedStatusCode, responseCode)
			s.repo.GetPasscodeRepo().Delete(ctx, pc)
		})
	}
}
