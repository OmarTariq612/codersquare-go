package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func SignJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return tokenString, err
}

func VerifyJWTCustom(tokenString string, claims jwt.Claims) (verifiedClaims jwt.Claims, err error, expired bool) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		// expired (signature is valid but token is expired)
		if er, ok := err.(*jwt.ValidationError); ok && !er.Is(jwt.ErrSignatureInvalid) && er.Is(jwt.ErrTokenExpired) {
			return nil, err, true
		}
		return nil, err, false // invalid or anything else
	}
	return token.Claims, nil, false
}
