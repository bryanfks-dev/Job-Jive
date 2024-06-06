package main

import (
	"fmt"
	"os"

	"configs" // ignore: import error
	"db"      // ignore: import error
	"models"  // ignore: import error

	"github.com/joho/godotenv" // ignore: import error
)

func loadConfig() (configs.Database, error) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return configs.Database{}, err
	}

	return configs.Database.Get(configs.Database{}), nil
}

func main() {
	// Check input arg length
	if len(os.Args) < 2 {
		fmt.Println("Usage: utils/addAdmin/main <username> <password>")
		os.Exit(1)
	}

	dbConf, err := loadConfig()

	// Ensure no error fetching config
	if err != nil {
		panic(err.Error())
	}

	// Connect to database
	err = db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()

	admin := models.Admin{
		Username: os.Args[1],
		Password: os.Args[2],
	}

	if db.ConnectionEstablished() {
		err := models.Admin.Insert(admin)

		// Ensure no error insertting admin data
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Admin `%s` has been created", admin.Username)

		return
	}

	fmt.Println("Error establish connection to database")
}
