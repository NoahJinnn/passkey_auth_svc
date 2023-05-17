package saltedge

import (
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

	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/pkg/httpx"
)

type SeClient struct {
	cred *config.SaltEdge
	req  *httpx.Req
}

func NewSeClient(cred *config.SaltEdge) *SeClient {

	req := httpx.NewReq("https://www.saltedge.com/api/v5")
	req.SetHeader("Accept", "application/json")
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("App-id", cred.AppId)
	req.SetHeader("Secret", cred.Secret)

	return &SeClient{
		cred: cred,
		req:  req,
	}
}

func (cl *SeClient) DoReq(method string, url string, query map[string][]string, reqBody interface{}) ([]byte, error) {

	cl.req.OverrideQ(query)

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

	httpReq, err := cl.req.PrepareReq(method, url, b)
	if err != nil {
		return nil, err
	}
	cl.SignedHeaders(httpReq.URL.String(), method, b)

	resp, err := cl.req.SendWithReq(httpReq)
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

	cl.req.SetHeader("Expires-at", fmt.Sprintf("%d", expiresAt))
	cl.req.SetHeader("Signature", signature)

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
