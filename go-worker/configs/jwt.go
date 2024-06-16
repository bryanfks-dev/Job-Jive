package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type JWT struct {
	Secret string
}

func (jwt *JWT) Load() error {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return err
	}

	jwt.Secret = os.Getenv("JWT_SECRET")

	return nil
}
