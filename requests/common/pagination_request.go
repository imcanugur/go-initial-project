package common

import "go-initial-project/validator"

type PaginationRequest struct {
	Page     int    `form:"page" json:"page" validate:"omitempty,min=1"`
	PageSize int    `form:"page_size" json:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy   string `form:"sort_by" json:"sort_by" validate:"omitempty"`
	Order    string `form:"order" json:"order" validate:"omitempty,oneof=asc desc"`
}

func (r *PaginationRequest) Validate() error {
	return validator.Validate.Struct(r)
}
