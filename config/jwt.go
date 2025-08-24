package config

func JWTSecret() []byte {
	return []byte(AppConfig.JWT.Secret)
}
