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
