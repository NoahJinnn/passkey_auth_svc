package dto

type AssetBodyRequest struct {
	Sheet        int     `json:"sheet,omitempty"`
	Section      int     `json:"section,omitempty"`
	Type         string  `json:"type,omitempty"`
	ProviderName string  `json:"provider_name,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Value        float64 `json:"value,omitempty"`
	Description  *string `json:"description,omitempty"`
}
