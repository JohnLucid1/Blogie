package db

import (
	"database/sql"
	// "errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CMPPasswords(inputPassword, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
	return err == nil
}

func VerifyPassword(DB *sql.DB, hashed_token string) (bool, error) {
	var claims MyClaims
	tok, err := jwt.ParseWithClaims(hashed_token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		fmt.Println("cannot parse with claims")
		return false, err
	}
	if tok.Valid {
		username := claims.Username
		password := claims.Password
		if err != nil {
			fmt.Println("Can't hash password")
			return false, err
		}

		db_password, err := GetPassFromUsername(DB, username)
		if err != nil {
			fmt.Println("Can't get password from claims username")
			return false, err
		}

		answ := CMPPasswords(password, db_password)
		return answ, nil
	}
	return false, jwt.ErrTokenExpired
}
