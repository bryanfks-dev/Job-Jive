package main

import (
	"db"
	"log"
	"models"

	"github.com/robfig/cron"
)

func main() {
	cron := cron.New()
	err := cron.AddFunc("0 0 0 1 * *", func() {
		tx, err := db.Conn.Begin()

		defer tx.Rollback()

		if err != nil {
			panic(err.Error())
		}

		err = models.ResetCurrentSalary()

		if err != nil {
			panic(err.Error())
		}

		if err := tx.Rollback(); err != nil {
			panic(err.Error())
		}
	})

	if err != nil {
		panic(err.Error())
	}

	log.Println("Start cron job")
	cron.Start()

	select{}
}