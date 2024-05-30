package models

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretKey() []byte {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

var secret_key = getSecretKey()

func CreateToken(user_id int) (string, error) {
	// Init token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * (24 * 7 * 7)).Unix(),
	})

	// Hash token
	token_string, err := token.SignedString(secret_key)

	return token_string, err
}

func VerifyToken(token_string string) (bool, error) {
	token, err :=
		jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return secret_key, nil
		})

	if err != nil {
		return false, err
	}

	return token.Valid, err
}
