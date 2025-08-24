package userrequests

import "go-initial-project/validator"

type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName  string `json:"last_name"  validate:"omitempty,min=2,max=50"`
	Email     string `json:"email"      validate:"omitempty,email"`
	Phone     string `json:"phone"      validate:"omitempty"`
}

func (r *UpdateUserRequest) Validate() error {
	return validator.Validate.Struct(r)
}
