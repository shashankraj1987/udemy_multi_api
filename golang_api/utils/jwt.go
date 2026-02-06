// Package utils provides utility functions for the application.
package utils

import (
	"errors"
	"time"
	"udemy-multi-api-golang/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtConfig *config.JWTConfig

// InitJWT initializes JWT configuration.
func InitJWT(cfg *config.JWTConfig) {
	jwtConfig = cfg
}

// GenerateToken generates a JWT token for a user.
func GenerateToken(email string, userID int64) (string, error) {
	if jwtConfig == nil {
		return "", errors.New("JWT config not initialized")
	}

	claims := jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * time.Duration(jwtConfig.TokenExpiryHrs)).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SecretKey))
}

// VerifyToken verifies a JWT token and returns the user ID.
func VerifyToken(token string) (int64, error) {
	if jwtConfig == nil {
		return 0, errors.New("JWT config not initialized")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtConfig.SecretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}

	return int64(userID), nil
}
