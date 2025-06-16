package auth

import (
	"errors"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecret") //read from env later

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string,error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims,error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{},error){
		return jwtKey, nil
	})
	fmt.Println(err," : ", token)
	if err!=nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims,nil
}
