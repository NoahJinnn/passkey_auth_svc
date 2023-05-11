package svcs

import (
	"bytes"
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
	"io"
	"net/http"
	"time"

	"github.com/hellohq/hqservice/ms/networth/app/dom"
	"github.com/hellohq/hqservice/ms/networth/config"
)

const (
	API_URL = "https://www.saltedge.com/api/v5"
)

type ISeAccountInfoSvc interface {
	CreateCustomer(ctx context.Context, ccr *dom.CreateCustomerReq) (*dom.CreateCustomerResp, error)
	CreateConnectSession(ctx context.Context, ccsr *dom.CreateConnectSessionReq) (*dom.CreateConnectSessionResp, error)
}

type seSvc struct {
	cfg *config.Config
}

func NewSeAccountInfoSvc(cfg *config.Config) ISeAccountInfoSvc {
	return &seSvc{
		cfg: cfg,
	}
}

func (svc *seSvc) CreateCustomer(ctx context.Context, ccr *dom.CreateCustomerReq) (*dom.CreateCustomerResp, error) {
	url := fmt.Sprintf("%s/customers", API_URL)

	resp, err := doReq("POST", url, ccr, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result dom.CreateCustomerResp
	err = json.Unmarshal(resp, &dom.HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateConnectSession(ctx context.Context, ccsr *dom.CreateConnectSessionReq) (*dom.CreateConnectSessionResp, error) {
	url := fmt.Sprintf("%s/connect_sessions/create", API_URL)

	resp, err := doReq("POST", url, ccsr, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result dom.CreateConnectSessionResp
	err = json.Unmarshal(resp, &dom.HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func doReq(method string, url string, reqBody interface{}, cred *config.SaltEdgeConfig) ([]byte, error) {
	b, err := json.Marshal(dom.HttpBody{
		Data: reqBody,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	headers := signedHeaders(req.URL.String(), req.Method, b, cred)
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

	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		fmt.Println("Error Response:", string(body))
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}
}

func signedHeaders(url, method string, body []byte, cred *config.SaltEdgeConfig) map[string]string {
	var signature string
	expiresAt := time.Now().Add(60 * time.Second).Unix()

	if cred.PK != "" {
		pk, err := parsePrivateKey([]byte((cred.PK)))
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
	}

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["App-id"] = cred.AppId
	headers["Secret"] = cred.Secret
	headers["Expires-at"] = fmt.Sprintf("%d", expiresAt)
	headers["Signature"] = signature

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
