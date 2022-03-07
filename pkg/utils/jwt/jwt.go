package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var Secret = []byte("Megrez")

// TODO: 测试 token 过期时间
var TokenExpireDuration = time.Hour * 24 * 7

type Claims struct {
	ID uint
	jwt.RegisteredClaims
}

// GenerateToken generate tokens by user id
func GenerateToken(id uint) (string, error) {
	c := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Megrez",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ParseToken parse claims from token
func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return claims, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return claims, fmt.Errorf("invalid token")
}
