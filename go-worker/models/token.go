package models

import (
	"fmt"
	"time"

	"configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getJWTSecret() []byte {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	jwt := configs.JWT.Get(configs.JWT{})

	return []byte(jwt.Secret)
}

var secret = getJWTSecret()

func CreateToken(record_id int, role string) (string, error) {
	// Init token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   record_id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * (24 * 7 * 7)).Unix(),
	})

	// Hash token
	token_string, err := token.SignedString(secret)

	return token_string, err
}

func ClaimsToken(token_string string) (jwt.MapClaims, error) {
	token, err :=
		jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
			// Parse token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return secret, nil
		})

	// Ensure there is no error in parsing token
	if err != nil {
		return nil, err
	}

	// Try claims token
	if jwt_map, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return jwt_map, nil
	}

	return nil, err
}
