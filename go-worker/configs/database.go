package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (db *Database) Load() error {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return err
	}

	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.Database = os.Getenv("DB_DATABASE")

	return nil
}
