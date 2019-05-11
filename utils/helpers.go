package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), err
}
