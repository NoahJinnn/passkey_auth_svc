package dom

type SeBodyReq struct {
	Data interface{} `json:"data"`
}

type SeBodyResp struct {
	Data interface{} `json:"data"`
}

type CreateCustomerData struct {
	Identifier string `json:"identifier"`
}

type CreateCustomerResp struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateConnectSessionData struct {
	CustomerId string  `json:"customer_id"`
	Consent    Consent `json:"consent"`
	Attempt    Attempt `json:"attempt"`
}

type Consent struct {
	FromDate   string   `json:"from_date"`
	PeriodDays int      `json:"period_days"`
	Scopes     []string `json:"scopes"`
}

type Attempt struct {
	FromDate    string   `json:"from_date"`
	FetchScopes []string `json:"fetch_scopes"`
}

type CreateCustomerSessionResp struct {
	ExpiresAt  string `json:"expires_at"`
	ConnectUrl string `json:"connect_url"`
}
