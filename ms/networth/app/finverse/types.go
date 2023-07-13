package finverse

import "time"

const PROVIDER_NAME = "finverse"

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountID string                 `json:"account_id"`
	Other     map[string]interface{} `json:"-"`
}

type Transactions struct {
	TotalTransactions int                    `json:"total_transactions"`
	Transactions      []interface{}          `json:"transactions"`
	Other             map[string]interface{} `json:"-"`
}

// Request payload types
type CreateCustomerToken struct {
	ClientID     string `json:"client_id"  validate:"required"`
	ClientSecret string `json:"client_secret"  validate:"required"`
	GrantType    string `json:"grant_type"  validate:"required"`
}

type CreateLinkToken struct {
	ClientID             string   `json:"client_id" validate:"required"`
	UserID               string   `json:"user_id" validate:"required"`
	RedirectURI          string   `json:"redirect_uri" validate:"required"`
	State                string   `json:"state" validate:"required"`
	GrantType            string   `json:"grant_type" validate:"required"`
	ResponseMode         string   `json:"response_mode" validate:"required"`
	ResponseType         string   `json:"response_type"`
	AutomaticDataRefresh string   `json:"automatic_data_refresh"`
	Countries            []string `json:"countries,omitempty"`
	InstitutionId        string   `json:"institution_id"`
	InstitutionStatus    string   `json:"institution_status"`
	Language             string   `json:"language"`
	LinkMode             string   `json:"link_mode"`
	ProductsSupported    []string `json:"products_supported,omitempty"`
	UIMode               string   `json:"ui_mode"`
	UserType             []string `json:"user_type,omitempty"`
}

type ExchangeAccessToken struct {
	Code string `json:"code" validate:"required"`
}

// Response payload types
type CustomerToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
	TokenType   string `json:"token_type"`
}

type LinkToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	IssuedAt    time.Time `json:"issued_at"`
	LinkURL     string    `json:"link_url"`
	TokenType   string    `json:"token_type"`
}

type AccessToken struct {
	AccessToken     string    `json:"access_token"`
	ExpiresIn       int       `json:"expires_in"`
	IssuedAt        time.Time `json:"issued_at"`
	LoginIdentityID string    `json:"login_identity_id"`
	LinkURL         string    `json:"link_url"`
	TokenType       string    `json:"token_type"`
}
