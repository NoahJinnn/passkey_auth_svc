package saltedge

// TODO: Create struct for DOM, DTO and swagger docs for client to integrate
type HttpBody struct {
	Data interface{} `json:"data"`
}

type CreateCustomerReq struct {
	Identifier string `json:"identifier"`
}

type CustomerResp struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	BlockedAt  string `json:"blocked_at,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type RemoveCustomerResp struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type CreateConnectSessionReq struct {
	CustomerId           string  `json:"customer_id"`
	IncludeFakeProviders bool    `json:"include_fake_providers"`
	Consent              Consent `json:"consent"`
	Attempt              Attempt `json:"attempt"`
}

type Consent struct {
	Scopes []string `json:"scopes"`
}

type Attempt struct {
	ReturnTo string `json:"return_to"`
}

type CreateConnectSessionResp struct {
	ConnectUrl string `json:"connect_url"`
	ExpiresAt  string `json:"expires_at"`
}
