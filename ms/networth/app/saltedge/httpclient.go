package saltedge

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"time"

	urlpkg "net/url"

	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/pkg/httpx"
)

type SeClient struct {
	cred *config.SaltEdge
	req  *httpx.Req
}

const BASE_URL = "https://www.saltedge.com/api/v5"

func NewSeClient(cred *config.SaltEdge) *SeClient {
	req := httpx.NewReq(BASE_URL, map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"App-id":       cred.AppId,
		"Secret":       cred.Secret,
	}, nil)

	return &SeClient{
		cred: cred,
		req:  req,
	}
}

func (cl *SeClient) DoReq(ctx context.Context, method string, path string, query map[string][]string, reqBody interface{}) ([]byte, error) {
	var b []byte
	if reqBody != nil {
		var err error
		b, err = json.Marshal(HttpBody{
			Data: reqBody,
		})
		if err != nil {
			return nil, err
		}
	}

	u, err := urlpkg.Parse(BASE_URL + path)
	if err != nil {
		return nil, err
	}
	signedHeaders := cl.SignedHeaders(u.String(), method, b)

	resp, err := cl.req.
		InitReq(ctx, method, path, b).
		WithDefaultOpts().
		WithHeaders(signedHeaders).
		Send()
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (cl *SeClient) SignedHeaders(url, method string, body []byte) map[string]string {
	var signature string
	expiresAt := time.Now().Add(60 * time.Second).Unix()
	headers := make(map[string]string)

	if cl.cred.PK != "" {
		pk, err := parsePrivateKey([]byte((cl.cred.PK)))
		if err != nil {
			panic(err)
		}

		payload := fmt.Sprintf("%d|%s|%s|", expiresAt, method, url)
		if method == "POST" {
			payload += string(body)
		}

		signature, err = sign(payload, pk)
		if err != nil {
			panic(err)
		}
		headers["Signature"] = signature
	}
	headers["Expires-at"] = fmt.Sprintf("%d", expiresAt)

	return headers
}

func parsePrivateKey(rawKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(rawKey)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

func sign(payload string, pk *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256([]byte(payload))
	signature, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
