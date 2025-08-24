package auth

import "go-initial-project/validator"

// LoginRequest godoc
// @Description Login payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginRequest) Validate() error {
	return validator.Validate.Struct(r)
}
