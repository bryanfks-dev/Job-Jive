package main

import (
	"log"
	"time"

	"configs"
	"db"
	"models"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

type Schedule struct {
	At time.Time
	Every time.Duration
}

func loadConfig() (configs.Database, error) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return configs.Database{}, err
	}

	return configs.Database.Get(configs.Database{}), err
}

func initResetEmployeeSalary(cron *cron.Cron) {
	err := cron.AddFunc("@monthly", func() {
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
}

func initCheckInHandler(cron *cron.Cron) {
	// Change this later into real code
	_, err := 
		models.ConfigJson.LoadConfig(models.ConfigJson{})
	
	// Ensure no error fetching config
	if err != nil {
		panic(err.Error())
	}

	cron.AddFunc("@midnight", func()  {
		// code here
	})
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

	initResetEmployeeSalary(cron)
	initCheckInHandler(cron)

	log.Println("Start cron job")
	cron.Start()

	defer cron.Stop()

	select {}
}
