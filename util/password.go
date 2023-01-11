package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func CheckPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
