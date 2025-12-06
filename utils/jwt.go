package utils

import (
	"fmt"
	"os"
	"time"

	"backend/models"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT creates a new access token
func GenerateJWT(userID string, email string) (string, *models.ErrorJson) {
	// Define token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return signedString , nil
}

// VerifyJWT checks the validity of the token and returns the claims if valid
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
