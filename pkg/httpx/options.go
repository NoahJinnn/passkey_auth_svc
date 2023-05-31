package httpx

type Opts struct {
	Headers map[string]string
	Query   map[string]string
	Body    []byte
}

// func WithBody(body []byte) ClientOption {
// 	return func(c *Req) {
// 		c.Body = body
// 	}
// }
