package dto

type ManualItemBody struct {
	ItemTableID  string  `json:"item_table_id" validate:"required"`
	Category     string  `json:"category" validate:"required"`
	Type         string  `json:"type" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Value        float64 `json:"value" validate:"required"`
	ProviderName string  `json:"provider_name" validate:"required"`
}
