package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	b_pass, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(b_pass), err
}

func CheckPassword(pass, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err == nil
}
