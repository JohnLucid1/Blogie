package db

import (
	// "time"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: Get fron .env
var KEY = []byte("4e5986f1-fbf3-49bf-8e47-b4768681dbe7")

type MyClaims struct {
	jwt.RegisteredClaims
	Username string
	Password string
}

func GenerateToken(username, password string) (string, error) {

	new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(168 * time.Hour)},
		},
		Username: username,
		Password: password,
	})
	string_token, err := new_token.SignedString(KEY)
	if err != nil {
		return "", err
	}
	return string_token, nil
}
