package models

import (
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

func CreateToken(credential string) (string, error) {
	// Init token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"credential": credential,
		"exp": time.Now().Add(time.Hour * (24 * 7)).Unix(),
	})

	// Hash token
	token_string, err := token.SignedString(secret_key)

	return token_string, err
}

func VerifyToken(token_string string) (bool, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})

	return token.Valid, err
}