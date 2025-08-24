package auth

import "go-initial-project/validator"

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name"  validate:"required,min=2,max=50"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=6"`
}

func (r *RegisterRequest) Validate() error {
	return validator.Validate.Struct(r)
}
