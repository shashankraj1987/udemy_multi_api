// Package utils provides utility functions for the application.
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password using bcrypt.
// Cost factor is set to 14 for a good balance between security and performance.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

// CheckPassword verifies if a password matches its hash.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
