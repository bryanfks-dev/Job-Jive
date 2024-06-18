package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Host string
	Port string
}

func (server *Server) Load() error {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return err
	}

	server.Host = os.Getenv("SERVER_HOST")
	server.Port = os.Getenv("SERVER_PORT")

	return nil
}
