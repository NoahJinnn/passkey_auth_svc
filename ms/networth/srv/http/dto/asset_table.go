package dto

type AssetTableBody struct {
	Sheet       int32  `json:"sheet"`
	Section     int32  `json:"section"`
	Description string `json:"description"`
}
