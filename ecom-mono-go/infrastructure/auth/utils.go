package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(signingKey string, customClaims map[string]any) (string, error) {
	claims := jwt.MapClaims(customClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}