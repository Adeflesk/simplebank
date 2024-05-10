package util

import (
	"fmt"

	bycrpt "golang.org/x/crypto/bcrypt"
)

// HashPassword return the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bycrpt.GenerateFromPassword([]byte(password), bycrpt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil

}

func CheckPassword(password, hash string) error {
	return bycrpt.CompareHashAndPassword([]byte(hash), []byte(password))

}
