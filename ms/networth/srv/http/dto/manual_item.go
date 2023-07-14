package dto

import "github.com/hellohq/hqservice/ms/networth/app/dao"

type ManualItemBody struct {
	ItemTableID  string       `json:"item_table_id" validate:"required"`
	Category     dao.Category `json:"category" validate:"required"`
	Type         string       `json:"type" validate:"required"`
	Description  string       `json:"description" validate:"required"`
	Value        float64      `json:"value" validate:"required"`
	ProviderName dao.Provider `json:"provider_name" validate:"required"`
}
