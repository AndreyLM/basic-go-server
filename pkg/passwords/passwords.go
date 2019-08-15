package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

const salt = 10

// IsValid - checks if password is valid
func IsValid(hash, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return false
	}
	return true
}

// Encrypt - encrypts password
func Encrypt(password string) (string, error) {
	str, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(str), err
}
