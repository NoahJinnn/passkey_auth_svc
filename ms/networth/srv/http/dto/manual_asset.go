package dto

type ManualTableBody struct {
	AssetTableID string  `json:"asset_table_id"`
	AssetType    string  `json:"asset_type"`
	Value        float64 `json:"value"`
}
