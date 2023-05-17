package saltedge

// TODO: Create struct for DOM, DTO and swagger docs for client to integrate
type HttpBody struct {
	Data interface{} `json:"data"`
}

// Request payload types
type CreateCustomer struct {
	Identifier string `json:"identifier"  validate:"required"`
}

type CreateConnectSession struct {
	CustomerId           string  `json:"customer_id" validate:"required"`
	IncludeFakeProviders bool    `json:"include_fake_providers" validate:"required"`
	Consent              Consent `json:"consent" validate:"required"`
	Attempt              Attempt `json:"attempt,omitempty"`
}

type Consent struct {
	Scopes []string `json:"scopes" validate:"required"`
}

type Attempt struct {
	ReturnTo string `json:"return_to"`
}

// Response payload types
type Customer struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	BlockedAt  string `json:"blocked_at,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type RemoveCustomer struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type ConnectSession struct {
	ConnectUrl string `json:"connect_url"`
	ExpiresAt  string `json:"expires_at"`
}
