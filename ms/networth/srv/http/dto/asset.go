package dto

type AssetBodyRequest struct {
	Sheet int `json:"sheet,omitempty"`
	// Section holds the value of the "section" field.
	Section int `json:"section,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// ProviderName holds the value of the "provider_name" field.
	ProviderName string `json:"provider_name,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
}
