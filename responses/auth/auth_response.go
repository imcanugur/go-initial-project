package auth

type AuthResponse struct {
	Token string `json:"token"`
	User  any    `json:"user"` // istersen burada UserResponse kullanabilirsin
}
