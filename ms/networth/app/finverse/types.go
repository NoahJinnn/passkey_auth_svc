package finverse

// Request payload types
type CreateCustomerToken struct {
	ClientId     string `json:"client_id"  validate:"required"`
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
	InstitutionID        string   `json:"institution_id"`
	InstitutionStatus    string   `json:"institution_status"`
	Language             string   `json:"language"`
	LinkMode             string   `json:"link_mode"`
	ProductsSupported    []string `json:"products_supported,omitempty"`
	UIMode               string   `json:"ui_mode"`
	UserType             []string `json:"user_type,omitempty"`
}

// Response payload types
type CustomerToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
	TokenType   string `json:"token_type"`
}
