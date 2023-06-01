package httpx

import (
	"context"
	"net/http"
	"net/http/httptest"
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
	s.req = NewReq(s.baseUrl, nil, nil)
}

func (s *TestRequestSuite) Test_NewReq_Success() {
	// Assert
	s.Equal(s.baseUrl, s.req.baseUrl)
	s.Nil(s.req.defaultOpts.Headers)
	s.Nil(s.req.defaultOpts.Query)
}

func (s *TestRequestSuite) Test_InitReq_Success() {
	// Arrange
	endpoint := "/test"

	// Act
	req := s.req.InitReq(s.ctx, http.MethodGet, endpoint, nil)

	// Assert
	s.NotNil(req.request)
	s.Equal(http.MethodGet, req.request.Method)
	s.Equal(s.baseUrl+endpoint, req.request.URL.String())

}

func (s *TestRequestSuite) Test_InitReq_Fail() {
	// Arrange
	endpoint := "\\\test"
	errMethod := "wrong"

	// Assert
	s.Panics(
		func() {
			s.req.InitReq(s.ctx, errMethod, endpoint, nil)
		},
	)

}

func (s *TestRequestSuite) Test_Request_Success() {
	// Arrange
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()
	req := NewReq(svc.URL, nil, nil)

	requests := []struct {
		name    string
		baseUrl string
		method  string
	}{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  http.MethodGet,
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  http.MethodPost,
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  http.MethodPut,
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  http.MethodPatch,
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  http.MethodDelete,
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			response, err := req.InitReq(s.ctx, exec.method, "", nil).
				Send()

			// Assert
			s.NotNil(response)
			s.NoError(err)
		})
	}
}

func (s *TestRequestSuite) Test_Request_Fail() {
	requests := []struct {
		name    string
		baseUrl string
		method  string
	}{
		{
			name:    "GET",
			baseUrl: s.baseUrl,
			method:  http.MethodGet,
		},
		{
			name:    "POST",
			baseUrl: s.baseUrl,
			method:  http.MethodPost,
		},
		{
			name:    "PUT",
			baseUrl: s.baseUrl,
			method:  http.MethodPut,
		},
		{
			name:    "PATCH",
			baseUrl: s.baseUrl,
			method:  http.MethodPatch,
		},
		{
			name:    "DELETE",
			baseUrl: s.baseUrl,
			method:  http.MethodDelete,
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			response, err := s.req.InitReq(s.ctx, exec.method, "", nil).
				Send()

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}
func (s *TestRequestSuite) Test_RequestWithOptions_Success() {
	// Arrange
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()
	req := NewReq(svc.URL, nil, nil)

	requests := []struct {
		name   string
		method string
		opts   *Opts
	}{
		{
			name:   "GET",
			method: http.MethodGet,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:   "POST",
			method: http.MethodPost,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:   "PUT",
			method: http.MethodPut,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:   "PATCH",
			method: http.MethodPatch,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:   "DELETE",
			method: http.MethodDelete,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			response, err := req.InitReq(s.ctx, exec.method, "", nil).
				WithOpts(exec.opts).
				Send()

			// Assert
			s.NotNil(response)
			s.NoError(err)
		})
	}
}
