package httpx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/pkg/errors"
)

type Req struct {
	baseUrl     string
	defaultOpts *Opts
	opts        *Opts
	request     *http.Request
}

type Opts struct {
	Headers map[string]string
	Query   map[string]string
}

func NewReq(baseUrl string, defaultHeaders map[string]string, defaultQuery map[string]string) *Req {
	return &Req{
		baseUrl: baseUrl,
		defaultOpts: &Opts{
			Headers: defaultHeaders,
			Query:   defaultQuery,
		},
		opts: nil,
	}
}

func (r *Req) Send() (*Resp, error) {
	if r.request == nil {
		return nil, fmt.Errorf("request is not initialized")
	}

	client := &http.Client{}
	resp, err := client.Do(r.request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	defer func() {
		r.request = nil
	}()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		result, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read failed response body")
		}
		defer resp.Body.Close()

		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("HTTP CALL FAILED [status code]: %d  [response]: %+v", resp.StatusCode, string(result)))
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

func (r *Req) InitReq(ctx context.Context, method string, path string, body []byte) *Req {
	req, err := http.NewRequestWithContext(ctx, method, r.baseUrl+path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.request = req
	return r
}

func (r *Req) WithOpts(opts *Opts) *Req {
	if r.request == nil {
		return r
	}
	applyOptions(r.request, opts)

	return r
}

func (r *Req) WithDefaultOpts() *Req {
	if r.request == nil {
		return r
	}
	applyOptions(r.request, r.defaultOpts)
	return r
}

func (r *Req) WithHeaders(headers map[string]string) *Req {
	if r.request == nil {
		return r
	}

	for k, v := range headers {
		r.request.Header.Set(k, v)
	}

	return r
}

func (r *Req) WithQuery(query map[string]string) *Req {
	if r.request == nil {
		return r
	}

	q := r.request.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	r.request.URL.RawQuery = q.Encode()

	return r
}

func applyOptions(req *http.Request, opts *Opts) {
	if opts == nil {
		return
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for k, v := range opts.Query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}

func (r *Req) String() string {
	// base url
	str := fmt.Sprintf("Base url:\n%s", r.baseUrl)

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
