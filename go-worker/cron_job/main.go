package main

import (
	"log"
	
	"configs"
	"db"
	"models"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func loadConfig() (configs.Database, error) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return configs.Database{}, err
	}

	return configs.Database.Get(configs.Database{}), err
}

func main() {
	config, err := loadConfig()

	// Ensure no error fetching config
	if err != nil {
		panic(err.Error())
	}

	err =
		db.Connect(config.User, config.Password, config.Host, config.Port, config.Database)

	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	cron := cron.New()
	err = cron.AddFunc( /* "0 0 0 1 * *" */ "1 * * * * *", func() {
		tx, err := db.Conn.Begin()

		if err != nil {
			panic(err.Error())
		}

		defer tx.Rollback()

		err = models.ResetCurrentSalary()

		if err != nil {
			panic(err.Error())
		}

		if err := tx.Rollback(); err != nil {
			panic(err.Error())
		}

		log.Println("User salary has been reset")
	})

	if err != nil {
		panic(err.Error())
	}

	log.Println("Start cron job")
	cron.Start()

	select {}
}
