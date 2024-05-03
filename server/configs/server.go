package configs

import "os"

type Server struct {
	Host string
	Port string
}

func (server Server) Get() Server {
	return Server{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}
}