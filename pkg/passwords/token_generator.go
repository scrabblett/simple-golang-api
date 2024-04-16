package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"simple-golang-api/internal/repository/users/model"
	"time"
)

type CustomClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJwtToken(info model.UserCredentials) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		UserId: info.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckJwtToken(tokenString string) (bool, error) {
	claims := &CustomClaims{}

	parser := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, parser)
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
