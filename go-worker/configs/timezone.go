package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Timezone struct {
	Zone string
}

func (timezone *Timezone) Load() error {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return err
	}

	timezone.Zone = os.Getenv("TIMEZONE_ZONE")

	return nil
}
