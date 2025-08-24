package common

import "go-initial-project/validator"

// IDRequest tekil ID taşıyan request'lerde kullanılır (delete, detail vs.)
type IDRequest struct {
	ID uint `uri:"id" json:"id" validate:"required,min=1"`
}

func (r *IDRequest) Validate() error {
	return validator.Validate.Struct(r)
}
