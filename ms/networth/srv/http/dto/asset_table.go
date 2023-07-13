package dto

type AssetTableBody struct {
	Sheet       int32  `json:"sheet"  validate:"required"`
	Section     int32  `json:"section"  validate:"required"`
	Description string `json:"description"  validate:"required"`
}
