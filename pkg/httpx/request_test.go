package httpx

import (
	"context"
	"io/ioutil"
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
	body := []byte("test")
	endpoint := "/test"

	// Act
	req := s.req.InitReq(s.ctx, http.MethodGet, endpoint, &Opts{
		Body: body,
	})

	// Assert
	s.NotNil(req.request)
	s.Equal(http.MethodGet, req.request.Method)
	s.Equal(s.baseUrl+endpoint, req.request.URL.String())

	requestBody, err := ioutil.ReadAll(req.request.Body)
	s.NoError(err)
	s.Equal(body, requestBody)
}

func (s *TestRequestSuite) Test_InitReq_Fail() {
	// Arrange
	body := []byte("test")
	endpoint := "\\\test"
	errMethod := "wrong"

	// Assert
	s.Panics(
		func() {
			s.req.InitReq(s.ctx, errMethod, endpoint, &Opts{
				Body: body,
			})
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
		method  func(ctx context.Context, path string, opts *Opts) (*Resp, error)
	}{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  req.Get,
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  req.Post,
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  req.Put,
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  req.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  req.Delete,
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			response, err := exec.method(s.ctx, "", nil)

			// Assert
			s.NotNil(response)
			s.NoError(err)
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
		name    string
		baseUrl string
		method  string
		f       func(ctx context.Context, path string, opts *Opts) (*Resp, error)
		opts    *Opts
	}{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  http.MethodGet,
			f:       req.Get,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  http.MethodPost,
			f:       req.Post,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  http.MethodPut,
			f:       req.Put,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  http.MethodPatch,
			f:       req.Patch,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  http.MethodDelete,
			f:       req.Delete,
			opts: &Opts{
				Headers: map[string]string{"key": "value"},
				Query:   map[string]string{"key": "value"},
			},
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			req.InitReq(s.ctx, exec.method, "", exec.opts)
			response, err := exec.f(s.ctx, "", nil)

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
		method  func(ctx context.Context, path string, opts *Opts) (*Resp, error)
	}{
		{
			name:    "GET",
			baseUrl: s.baseUrl,
			method:  s.req.Get,
		},
		{
			name:    "POST",
			baseUrl: s.baseUrl,
			method:  s.req.Post,
		},
		{
			name:    "PUT",
			baseUrl: s.baseUrl,
			method:  s.req.Put,
		},
		{
			name:    "PATCH",
			baseUrl: s.baseUrl,
			method:  s.req.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: s.baseUrl,
			method:  s.req.Delete,
		},
	}

	for _, exec := range requests {
		s.Suite.Run(exec.name, func() {
			// Act
			response, err := exec.method(s.ctx, "", nil)

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}
