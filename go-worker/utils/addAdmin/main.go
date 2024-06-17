package main

import (
	"fmt"
	"os"

	"configs" // ignore: import error
	"db"      // ignore: import error
	"models"  // ignore: import error
)

func main() {
	// Check input arg length
	if len(os.Args) < 2 {
		fmt.Println("Usage: addAdmin <username> <password>")
		os.Exit(1)
	}

	db_config := configs.Database{}

	err := db_config.Load()

	// Ensure no error fetching config
	if err != nil {
		panic(err.Error())
	}

	// Connect to database
	err = db.Connect(db_config.User, db_config.Password, db_config.Host, db_config.Port, db_config.Database)

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
		// Check if username already exist
		_, err := 
			models.Admin.GetUsingUsername(models.Admin{}, admin.Username)

		if err == nil {
			fmt.Errorf("Error: Admin username already exist")

			return
		}

		err = models.Admin.Insert(admin)

		// Ensure no error insertting admin data
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Admin `%s` has been created", admin.Username)

		return
	}

	fmt.Println("Error establish connection to database")
}
