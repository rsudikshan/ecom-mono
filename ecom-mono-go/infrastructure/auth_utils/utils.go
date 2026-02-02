package auth_utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	ACCESS_TOKEN TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
	EMAIL_VERIFICATION_TOKEN TokenType = "email_verification"
)


func CreateToken(signingKey string, customClaims map[string]any) (string, error) {
	claims := jwt.MapClaims(customClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}