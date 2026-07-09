package auth

func HashPassword(password string) (string, error) {
	return password, nil
}

func CheckPasswordHash(password, hash string) bool {
	return password == hash
}
