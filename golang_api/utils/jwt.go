package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {

	// Here, we are checking whether the token we got from user, uses the same signing method
	// that we used to encrypt the original token.

	parsed_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing Method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not Parse token")
	}

	if !parsed_token.Valid {
		return 0, errors.New("Invalid Token !")
	}

	claims, ok := parsed_token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// .() <- This is used for Type Checking

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
