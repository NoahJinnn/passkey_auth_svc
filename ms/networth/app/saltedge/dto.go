package saltedge

type CreateCustomer struct {
	Identifier string `json:"identifier"`
}

type CreateConnectSession struct {
	CustomerId           string  `json:"customer_id"`
	IncludeFakeProviders bool    `json:"include_fake_providers"`
	Consent              Consent `json:"consent"`
	Attempt              Attempt `json:"attempt"`
}
