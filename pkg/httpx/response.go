package httpx

import (
	"encoding/json"
	"net/http"
)

type (
	Resp struct {
		resp *http.Response
		body []byte
	}
)

func (r *Resp) Body() []byte {
	return r.body
}

func (r *Resp) Unmarshal(v any) error {
	return json.Unmarshal(r.body, &v)
}

func (r *Resp) Status() int {
	return r.resp.StatusCode
}

func (r *Resp) Headers() http.Header {
	return r.resp.Header
}

func (r *Resp) Cookies() []*http.Cookie {
	return r.resp.Cookies()
}

func (r *Resp) Ok() bool {
	return r.resp.StatusCode >= 200 && r.resp.StatusCode <= 299
}

func (r *Resp) Get() *http.Response {
	return r.resp
}
