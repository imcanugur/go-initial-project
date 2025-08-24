package userrequests

import "go-initial-project/validator"

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name"  validate:"required,min=2,max=50"`
	Email     string `json:"email"      validate:"required,email"`
	Phone     string `json:"phone"      validate:"required"`
}

func (r *CreateUserRequest) Validate() error {
	return validator.Validate.Struct(r)
}
