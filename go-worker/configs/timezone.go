package configs

import (
	"log"
	"os"
	"time"

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

func (timezone Timezone) GetTimeZone() (*time.Location, error) {
	err := timezone.Load()

	// Ensure no error get timezone from env
	if err != nil {
		log.Panic("Error get timezone from env: ", err.Error())

		return &time.Location{}, err
	}

	zone, err := time.LoadLocation(timezone.Zone)

	// Ensure no error getting timezone
	if err != nil {
		log.Panic("Error get timezone", err.Error())

		return &time.Location{}, err
	}

	return zone, nil
}
