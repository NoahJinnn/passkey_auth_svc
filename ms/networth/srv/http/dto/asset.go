package dto

type AssetTableRequest struct {
	Sheet       int32  `json:"sheet,omitempty"`
	Section     int32  `json:"section,omitempty"`
	Description string `json:"description,omitempty"`
}
