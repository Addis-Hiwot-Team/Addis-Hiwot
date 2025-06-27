package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(claim jwt.Claims, key string) (string, error) {
	// This function should generate a JWT token based on the provided claims.
	// Implementation details would depend on the JWT library used.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(key))
}
