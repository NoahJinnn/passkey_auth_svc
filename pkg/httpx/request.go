package httpx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Req struct {
	BaseUrl string
	Headers map[string][]string
	Query   map[string][]string
	Request *http.Request
}

func NewReq(baseUrl string) *Req {
	return &Req{
		BaseUrl: baseUrl,
		Headers: map[string][]string{},
		Query:   map[string][]string{},
	}
}

// Appends request header
func (r *Req) SetHeader(key string, value ...string) {
	r.Headers[key] = append(r.Headers[key], value...)
}

// Unsets request header
func (r *Req) UnsetHeader(key string) {
	delete(r.Headers, key)
}

// Appends request query
func (r *Req) SetQ(key string, value ...string) {
	r.Query[key] = append(r.Query[key], value...)
}

// Unsets request query
func (r *Req) UnsetQ(key string) {
	delete(r.Query, key)
}

// Appends request query
func (r *Req) OverrideQ(q map[string][]string) {
	r.Query = q
}

func (r *Req) Send(method string, path string, body []byte) (*Resp, error) {
	if r.Request == nil {
		_, err := r.PrepareReq(method, path, body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to prepare request")
		}
	}
	client := &http.Client{}
	resp, err := client.Do(r.Request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("request failed: %+v", resp)
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}
	fmt.Printf("request: %+v\n", resp.Body)

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return &Resp{
		resp: resp,
		body: result,
	}, nil
}

func (r *Req) Get(path string) (*Resp, error) {
	return r.Send(http.MethodGet, path, nil)
}

func (r *Req) Post(path string, body []byte) (*Resp, error) {
	return r.Send(http.MethodPost, path, body)
}

func (r *Req) Put(path string, body []byte) (*Resp, error) {
	return r.Send(http.MethodPut, path, body)
}

func (r *Req) Delete(path string, body []byte) (*Resp, error) {
	return r.Send(http.MethodDelete, path, body)
}

func (r *Req) String() string {
	// base url
	str := fmt.Sprintf("Base url:\n%s", r.BaseUrl)

	// header
	if len(r.Headers) > 0 {
		str = str + fmt.Sprintf("\nHeader:\n%s", printHeaders(r.Headers))
	}
	return str
}

func (r *Req) PrepareReq(method string, path string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, r.BaseUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// set headers
	for key, values := range r.Headers {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	// set query
	q := req.URL.Query()
	for key, values := range r.Query {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()

	r.Request = req

	return req, err
}

func printHeaders(data map[string][]string) string {
	lines := []string{}

	for key, arr := range data {
		lines = append(lines, fmt.Sprintf("\t%s: %s", key, strings.Join(arr, ", ")))
	}

	return strings.Join(lines, "\n")
}
