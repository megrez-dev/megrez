package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

func Parse(tokenString string, keyFunc jwt.Keyfunc) (jwt.MapClaims, error) {
	return jwt.Parse(tokenString, keyFunc)
}
