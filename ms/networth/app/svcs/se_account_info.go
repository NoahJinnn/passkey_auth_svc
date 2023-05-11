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
	CreateConnectSession(ctx context.Context, reqBody interface{}) (interface{}, error)
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

	response, err := doReq("POST", url, ccr, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result dom.CreateCustomerResp
	err = json.Unmarshal(response, &dom.HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateConnectSession(ctx context.Context, reqBody interface{}) (interface{}, error) {
	url := fmt.Sprintf("%s/connect_sessions/create", API_URL)

	response, err := doReq("POST", url, reqBody, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return response, nil
}

func doReq(method string, url string, reqBody interface{}, credentials *config.SaltEdgeConfig) ([]byte, error) {
	seBody := dom.HttpBody{
		Data: reqBody,
	}

	body, err := json.Marshal(seBody)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	headers := signedHeaders(request.URL.String(), request.Method, body, credentials)
	request.Header = make(http.Header)
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		return respBody, nil
	} else {
		fmt.Println("Error Response:", string(respBody))
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}
}

func signedHeaders(url, method string, params []byte, credentials *config.SaltEdgeConfig) map[string]string {
	var signature string
	expiresAt := time.Now().Add(60 * time.Second).Unix()

	if credentials.PK != "" {
		privateKey, err := parsePrivateKey([]byte((credentials.PK)))
		if err != nil {
			panic(err)
		}

		payload := fmt.Sprintf("%d|%s|%s|", expiresAt, method, url)
		if method == "POST" {
			payload += string(params)
		}

		signature, err = sign(payload, privateKey)
		if err != nil {
			panic(err)
		}
	}

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["App-id"] = credentials.AppId
	headers["Secret"] = credentials.Secret
	headers["Expires-at"] = fmt.Sprintf("%d", expiresAt)
	headers["Signature"] = signature

	return headers
}

func parsePrivateKey(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func sign(payload string, privateKey *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256([]byte(payload))
	signatureBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}
