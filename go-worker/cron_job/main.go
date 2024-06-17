package main

import (
	"fmt"
	"log"
	"time"

	"configs"
	"db"
	"models"

	"github.com/robfig/cron"
)

type Schedule struct {
	At    time.Time
	Every time.Duration
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
		models.ConfigJson{}.LoadConfig()

	// Ensure no error fetching db_config
	if err != nil {
		panic(err.Error())
	}

	cron.AddFunc("@midnight", func() {
		// code here
	})
}

func main() {
	var db_config = configs.Database{}

	err := db_config.Load()

	// Ensure no error fetching db_config
	if err != nil {
		panic(err.Error())
	}

	err =
		db.Connect(db_config.User, db_config.Password, db_config.Host, db_config.Port, db_config.Database)

	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	cron := cron.New()

	initResetEmployeeSalary(cron)
	initCheckInHandler(cron)

	fmt.Println("Start cron job")
	cron.Start()

	defer cron.Stop()

	select {}
}
