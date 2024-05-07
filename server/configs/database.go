package configs

import "os"

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (db Database) Get() Database {
	return Database{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
