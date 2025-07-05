package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(claim jwt.Claims, key string) (string, error) {
	// This function should generate a JWT token based on the provided claims.
	// Implementation details would depend on the JWT library used.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(key))
}

func ParseJwt(claims jwt.Claims, token, key string) error {
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	if !jwtToken.Valid {
		return fmt.Errorf("jwt invalid")
	}
	return nil
}
