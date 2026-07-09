package auth

type TokenManager struct {
	secret string
}

func NewTokenManager(secret string) TokenManager {
	return TokenManager{secret: secret}
}
