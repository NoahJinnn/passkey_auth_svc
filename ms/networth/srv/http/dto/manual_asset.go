package dto

type ManualAssetBody struct {
	AssetTableID string  `json:"asset_table_id" validate:"required"`
	AssetType    string  `json:"asset_type" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Value        float64 `json:"value" validate:"required"`
	ProviderName string  `json:"provider_name" validate:"required"`
}
