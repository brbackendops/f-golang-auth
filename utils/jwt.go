package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func JWTVerifyAndDecode(tokenString string, claims jwt.MapClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	return token, err
}

func JWTVerify(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return err == nil
}

func JWTSign(data jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tok, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tok, nil
}
