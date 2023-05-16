package saltedge

// TODO: Create struct for DOM, DTO and swagger docs for client to integrate
type HttpBody struct {
	Data interface{} `json:"data"`
}

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

type Consent struct {
	Scopes []string `json:"scopes"`
}

type Attempt struct {
	ReturnTo string `json:"return_to"`
}

type ConnectSession struct {
	ConnectUrl string `json:"connect_url"`
	ExpiresAt  string `json:"expires_at"`
}
