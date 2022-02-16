package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(SecretKey), nil

	})

	if err != nil {
		return nil, err
	}

	return t, nil

}
