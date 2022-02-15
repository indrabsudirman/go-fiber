package utils

import "github.com/golang-jwt/jwt/v4"

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webToken, nil
}
