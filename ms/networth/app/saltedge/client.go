package saltedge

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hellohq/hqservice/ms/networth/config"
)

type SeClient struct {
	cred *config.SaltEdgeConfig
}

func NewSeClient(cred *config.SaltEdgeConfig) *SeClient {
	return &SeClient{
		cred: cred,
	}
}

func (cl *SeClient) DoReq(method string, url string, reqBody interface{}) ([]byte, error) {
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

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	headers := cl.SignedHeaders(req.URL.String(), req.Method, b)
	req.Header = make(http.Header)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", string(body))
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return body, nil
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

	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["App-id"] = cl.cred.AppId
	headers["Secret"] = cl.cred.Secret
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
