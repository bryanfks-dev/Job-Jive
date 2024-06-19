package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Gemini struct {
	APIKey string
}

func (gemini *Gemini) Load() error {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return err
	}

	gemini.APIKey = os.Getenv("GEMINI_API_KEY")

	return nil
}
