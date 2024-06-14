package configs

import "os"

type Timezone struct {
	Zone string
}

func (timezone Timezone) Get() Timezone {
	return Timezone{
		Zone: os.Getenv("TIMEZONE_ZONE"),
	}
}
