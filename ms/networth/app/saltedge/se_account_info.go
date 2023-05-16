package saltedge

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hellohq/hqservice/ms/networth/config"
)

const (
	API_URL = "https://www.saltedge.com/api/v5"
)

type Ctx = context.Context

type ISeAccountInfoSvc interface {
	Customer(ctx Ctx, customerId string) (*CustomerResp, error)
	CreateCustomer(ctx Ctx, ccr *CreateCustomerReq) (*CustomerResp, error)
	RemoveCustomer(ctx context.Context, customerId string) (*RemoveCustomerResp, error)
	CreateConnectSession(ctx Ctx, ccsr *CreateConnectSessionReq) (*CreateConnectSessionResp, error)
	GetConnectionByCustomerId(ctx Ctx, customerId string) (interface{}, error)
	GetAccountByConnectionId(ctx context.Context, connectionId string) (interface{}, error)
	GetTxByConnectionIdAndAccountId(ctx context.Context, connectionId string, accountId string) (interface{}, error)
}

type seSvc struct {
	client *SeClient
}

func NewSeAccountInfoSvc(cfg *config.Config) ISeAccountInfoSvc {
	client := NewSeClient(cfg.SaltEdgeConfig)
	return &seSvc{
		client: client,
	}
}

func (svc *seSvc) Customer(ctx context.Context, customerId string) (*CustomerResp, error) {
	url := fmt.Sprintf("%s/customers/%s", API_URL, customerId)

	resp, err := svc.client.DoReq("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result CustomerResp
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateCustomer(ctx context.Context, ccr *CreateCustomerReq) (*CustomerResp, error) {
	url := fmt.Sprintf("%s/customers", API_URL)

	resp, err := svc.client.DoReq("POST", url, ccr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result CustomerResp
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) RemoveCustomer(ctx context.Context, customerId string) (*RemoveCustomerResp, error) {
	url := fmt.Sprintf("%s/customers/%s", API_URL, customerId)

	resp, err := svc.client.DoReq("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result RemoveCustomerResp
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateConnectSession(ctx context.Context, ccsr *CreateConnectSessionReq) (*CreateConnectSessionResp, error) {
	url := fmt.Sprintf("%s/connect_sessions/create", API_URL)

	resp, err := svc.client.DoReq("POST", url, ccsr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result CreateConnectSessionResp
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetConnectionByCustomerId(ctx context.Context, customerId string) (interface{}, error) {
	url := fmt.Sprintf("%s/customers/%s", API_URL, customerId)

	resp, err := svc.client.DoReq("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetAccountByConnectionId(ctx context.Context, connectionId string) (interface{}, error) {
	url := fmt.Sprintf("%s/accounts?connection_id=%s", API_URL, connectionId)

	resp, err := svc.client.DoReq("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetTxByConnectionIdAndAccountId(ctx context.Context, connectionId string, accountId string) (interface{}, error) {
	url := fmt.Sprintf("%s/transactions?connection_id=%s&account_id=%s", API_URL, connectionId, accountId)

	resp, err := svc.client.DoReq("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}
