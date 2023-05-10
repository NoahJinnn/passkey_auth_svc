package svcs

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
	"os"
	"time"

	"github.com/hellohq/hqservice/ms/networth/app/dom"
	"github.com/hellohq/hqservice/ms/networth/config"
)

const (
	API_URL = "https://www.saltedge.com/api/v5"
)

type ISeSvc interface {
	CreateCustomer() (*dom.SeBodyResp, error)
}

type seSvc struct {
	cfg *config.Config
}

func NewSeSvc(cfg *config.Config) ISeSvc {
	return &seSvc{
		cfg: cfg,
	}
}

func (svc *seSvc) CreateCustomer() (*dom.SeBodyResp, error) {
	url := fmt.Sprintf("%s/customers", API_URL)
	params := dom.SeBodyReq{
		Data: dom.CreateCustomerData{
			Identifier: "my_2unique_identifier",
		},
	}

	body, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	response, err := doReq(request, params, svc.cfg.SaltEdgeConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Decode response
	var resp = dom.SeBodyResp{
		Data: dom.CreateCustomerData{},
	}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println("Response:", resp)

	return &resp, nil
}

// func (svc *seSvc) CreateConnectSession() ([]byte, error) {
// 	url := fmt.Sprintf("%s/connect_sessions/create", API_URL)
// }

func doReq(options *http.Request, reqBody interface{}, credentials *config.SaltEdgeConfig) ([]byte, error) {
	headers := signedHeaders(options.URL.String(), options.Method, reqBody, credentials)

	options.Header = make(http.Header)
	for key, value := range headers {
		options.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(options)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		return body, nil
	} else {
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
