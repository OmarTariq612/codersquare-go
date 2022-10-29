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

// func VerifyJWT(tokenString string) jwt.Claims {
// 	token, _ := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_KEY")), nil
// 	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
// 	return token.Claims
// }

// func VerifyJWT(tokenString string) (*types.JWTObject, error) {
// 	jwtObj := &types.JWTObject{}
// 	_, err := jwt.ParseWithClaims(tokenString, jwtObj, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_KEY")), nil
// 	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
// 	return jwtObj, err
// }

func VerifyJWTCustom(tokenString string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	return token.Claims, err
}
