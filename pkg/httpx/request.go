package httpx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Req struct {
	BaseUrl     string
	defaultOpts *Opts
}

func NewReq(baseUrl string) *Req {
	return &Req{
		BaseUrl: baseUrl,
		defaultOpts: &Opts{
			Headers: map[string]string{},
			Query:   map[string]string{},
			Body:    nil,
		},
	}
}

func (r *Req) Send(ctx context.Context, method string, path string, opts *Opts) (*Resp, error) {
	httpReq, err := r.PrepareReq(ctx, method, path, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare request")
	}
	return r.SendReq(httpReq)
}

func (r *Req) SendReq(httpReq *http.Request) (*Resp, error) {
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("request failed: %+v", resp)
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return &Resp{
		resp: resp,
		body: result,
	}, nil
}

func (r *Req) Get(ctx context.Context, path string, opts *Opts) (*Resp, error) {
	return r.Send(ctx, http.MethodGet, path, opts)
}

func (r *Req) Post(ctx context.Context, path string, opts *Opts) (*Resp, error) {
	return r.Send(ctx, http.MethodPost, path, opts)
}

func (r *Req) Put(ctx context.Context, path string, opts *Opts) (*Resp, error) {
	return r.Send(ctx, http.MethodPut, path, opts)
}

func (r *Req) Patch(ctx context.Context, path string, opts *Opts) (*Resp, error) {
	return r.Send(ctx, http.MethodPatch, path, opts)
}

func (r *Req) Delete(ctx context.Context, path string, opts *Opts) (*Resp, error) {
	return r.Send(ctx, http.MethodDelete, path, opts)
}

func (r *Req) PrepareReq(ctx context.Context, method string, path string, opts *Opts) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, r.BaseUrl+path, bytes.NewBuffer(opts.Body))
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range opts.Headers {
		req.Header.Set(key, value)
	}

	// set query
	q := req.URL.Query()
	for key, value := range opts.Query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, err
}

func (r *Req) String() string {
	// base url
	str := fmt.Sprintf("Base url:\n%s", r.BaseUrl)

	// header
	if len(r.defaultOpts.Headers) > 0 {
		str = str + fmt.Sprintf("\nHeader:\n%s", printHeaders(r.defaultOpts.Headers))
	}
	return str
}

func printHeaders(data map[string]string) string {
	lines := []string{}

	for key, val := range data {
		lines = append(lines, fmt.Sprintf("\t%s: %s", key, val))
	}

	return strings.Join(lines, "\n")
}
