package main

import (
	"fmt"
	"os"

	"configs" // ignore: import error
	"models" // ignore: import error
	"db" // ignore: import error

	"github.com/joho/godotenv" // ignore: import error
)

func loadConfig() configs.Database {
	// Load .env
	err := godotenv.Load("../.env")

	if err != nil {
		panic(err.Error())
	}

	return configs.Database.Get(configs.Database{})
}

func main() {
	// Check input arg length
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./addAdmin <username> <password>")
		os.Exit(1)
	}

	dbConf := loadConfig()

	// Connect to database
	db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	defer db.Conn.Close()

	admin := models.Admin{
		Username: os.Args[1],
		Password: os.Args[2],
	}

	models.Admin.AddToDB(admin)

	fmt.Println("New Admin has been created")
}