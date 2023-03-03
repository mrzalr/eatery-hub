package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(claims interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"claims": claims,
	})

	return token.SignedString([]byte(os.Getenv("APP_JWTSECRETKEY")))
}
