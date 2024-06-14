package configs

import "os"

type JWT struct {
	Secret string
}

func (jwt JWT) Get() JWT {
	return JWT{
		Secret: os.Getenv("JWT_SECRET"),
	}
}
