package dto

import "github.com/hellohq/hqservice/ms/networth/app/dao"

type ItemTableBody struct {
	Category    dao.Category `json:"category"  validate:"required"`
	Sheet       int32        `json:"sheet"  validate:"required"`
	Section     int32        `json:"section"  validate:"required"`
	Description string       `json:"description"  validate:"required"`
}
