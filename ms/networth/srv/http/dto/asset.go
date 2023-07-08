package dto

type AssetBodyRequest struct {
	Sheet        int32   `json:"sheet,omitempty"`
	Section      int32   `json:"section,omitempty"`
	Type         string  `json:"type,omitempty"`
	ProviderName string  `json:"provider_name,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Value        float64 `json:"value" validate:"required"`
	Description  string  `json:"description,omitempty"`
}
