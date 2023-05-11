package dom

// TODO: Create struct for DOM, DTO and swagger docs for client to integrate
type HttpBody struct {
	Data interface{} `json:"data"`
}

type CreateCustomerReq struct {
	Identifier string `json:"identifier"`
}

type CreateCustomerResp struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	BlockedAt  string `json:"blocked_at,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
