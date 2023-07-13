package dto

type ManualAssetBody struct {
	AssetTableID string  `json:"asset_table_id"`
	AssetType    string  `json:"asset_type"`
	Value        float64 `json:"value"`
	ProviderName string  `json:"provider_name"`
}
