package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateJWT(tokenString string) (*Claims, error) {
	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
