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

func CreateToken(username string) string {
	// Init token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * (24 * 7)).Unix(),
	})

	// Hash token
	token_string, err := token.SignedString(secret_key)

	if err != nil {
		panic(err.Error())
	}

	return token_string
}

func VerifyToken(token_string string) bool {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})

	if err != nil {
		panic(err.Error())
	}

	return token.Valid
}