package httpx

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestRequestSuite struct {
	suite.Suite
	ctx     context.Context
	baseUrl string
	req     *Req
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestRequestSuite))
}

func (s *TestRequestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.baseUrl = "http://localhost"
	s.req = NewReq(s.baseUrl)
}

func (s *TestRequestSuite) Test_NewReq_Success() {
	// Assert
	s.Equal(s.baseUrl, s.req.BaseUrl)
	s.Empty(s.req.defaultOpts.Headers)
	s.Empty(s.req.defaultOpts.Query)
}

func (s *TestRequestSuite) Test_PrepareReq_Success() {
	// Arrange
	body := []byte("test")
	endpoint := "/test"

	// Act
	request, err := s.req.PrepareReq(s.ctx, http.MethodGet, endpoint, &Opts{
		Body: body,
	})

	// Assert
	s.NoError(err)
	s.NotNil(request)
	s.Equal(http.MethodGet, request.Method)
	s.Equal(s.baseUrl+endpoint, request.URL.String())

	requestBody, err := ioutil.ReadAll(request.Body)
	s.NoError(err)
	s.Equal(body, requestBody)
}

func (s *TestRequestSuite) Test_PrepareReq_Fail() {
	// Arrange
	body := []byte("test")
	endpoint := "/test"
	errMethod := "wrong"

	// Act
	request, err := s.req.PrepareReq(nil, errMethod, endpoint, &Opts{
		Body: body,
	})

	// Assert
	s.Nil(request)
	s.Error(err)
}

// func (s *TestRequestSuite) Test_Request_ShouldRunSuccesfully() {
// 	// Arrange
// 	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer svc.Close()

// 	requests := []struct {
// 		name    string
// 		baseUrl string
// 		method  func(path string, body []byte) (*Resp, error)
// 	}{
// 		{
// 			name:    "GET",
// 			baseUrl: svc.URL,
// 			method:  s.req.Get,
// 		},
// 		{
// 			name:    "POST",
// 			baseUrl: svc.URL,
// 			method:  s.req.Post,
// 		},
// 		{
// 			name:    "PUT",
// 			baseUrl: svc.URL,
// 			method:  s.req.Put,
// 		},
// 		{
// 			name:    "PATCH",
// 			baseUrl: svc.URL,
// 			method:  s.req.Patch,
// 		},
// 		{
// 			name:    "DELETE",
// 			baseUrl: svc.URL,
// 			method:  s.req.Delete,
// 		},
// 		// {
// 		// 	name:    "CONNECT",
// 		// 	baseUrl: svc.URL,
// 		// 	method:  s.req.Connect,
// 		// },
// 		// {
// 		// 	name:    "OPTIONS",
// 		// 	baseUrl: svc.URL,
// 		// 	method:  s.req.Options,
// 		// },
// 		// {
// 		// 	name:    "TRACE",
// 		// 	baseUrl: svc.URL,
// 		// 	method:  s.req.Trace,
// 		// },
// 	}

// 	for _, req := range requests {
// 		s.Suite.Run(req.name, func() {
// 			// Act
// 			response, err := req.method("")

// 			// Assert
// 			s.NotNil(response)
// 			s.NoError(err)
// 		})
// 	}
// }
