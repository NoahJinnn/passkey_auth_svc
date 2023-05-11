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
	"os"
	"time"

	"github.com/hellohq/hqservice/ms/networth/config"
)

const (
	API_URL = "https://www.saltedge.com/api/v5"
)

type ISeAccountInfoSvc interface {
	CreateCustomer(ctx context.Context, reqBody interface{}) (interface{}, error)
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

func (svc *seSvc) CreateCustomer(ctx context.Context, reqBody interface{}) (interface{}, error) {
	url := fmt.Sprintf("%s/customers", API_URL)

	response, err := doReq("POST", url, reqBody, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return response, nil
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

func doReq(method string, url string, reqBody interface{}, credentials *config.SaltEdgeConfig) (interface{}, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	headers := signedHeaders(request.URL.String(), request.Method, reqBody, credentials)
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

	var result struct {
		Data interface{} `json:"data"`
	}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		return &result, nil
	} else {
		fmt.Println("Error Response:", string(respBody))
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}
}

func signedHeaders(url, method string, params interface{}, credentials *config.SaltEdgeConfig) map[string]string {
	privateKeyBytes, err := os.ReadFile("configs/saltedge-pki/private.pem")
	if err != nil {
		panic(err)
	}

	privateKey, err := parsePrivateKey(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	expiresAt := time.Now().Add(60 * time.Second).Unix()
	payload := fmt.Sprintf("%d|%s|%s|", expiresAt, method, url)
	if method == "POST" {
		payloadBytes, err := json.Marshal(params)
		if err != nil {
			panic(err)
		}
		payload += string(payloadBytes)
	}

	signature, err := sign(payload, privateKey)
	if err != nil {
		panic(err)
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
