package main

import (
	"db"
	"fmt"
	"log"
	"models"
	"configs"

	"github.com/robfig/cron"
)

type Config struct {
    AbsenceQuota int `json:"absence_quota"`
}

func resetWeekly() {
	log.Println("Running weekly reset")
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE attendance_statistics SET weekly = 0")
	if err != nil {
		log.Panic(err)
	}

	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
	log.Println("Weekly reset complete")
}

func resetMonthly() {
	log.Println("Running monthly reset")
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE attendance_statistics SET monthly = 0")
	if err != nil {
		log.Panic(err)
	}

	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
	log.Println("Monthly reset complete")
}

func resetKuotaAbsen(absenceQuota int) {
    log.Println("Running yearly reset for kuota absen")
    tx, err := db.Conn.Begin()
    if err != nil {
        log.Panic(err)
    }
    defer tx.Rollback()

    _, err = tx.Exec("UPDATE attendance_statistics SET annual_leave = ?", absenceQuota)
    if err != nil {
        log.Panic(err)
    }

    if err := tx.Commit(); err != nil {
        log.Panic(err)
    }
    log.Println("Yearly reset for kuota absen complete")
}

func resetCurrentSalary() {
	log.Println("Running monthly reset for current salary")
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE salaries SET current_salary = initial_salary")
	if err != nil {
		log.Panic(err)
	}

	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
	log.Println("Monthly reset for current salary complete")
}

func main() {
	var db_config = configs.Database{}

	err := db_config.Load()
	// Ensure no error fetching db_config
	if err != nil {
		panic(err.Error())
	}

	err = db.Connect(db_config.User, db_config.Password, db_config.Host, db_config.Port, db_config.Database)
	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	// Load the config file for AbsenceQuota
    var config models.ConfigJson
    config, err = config.LoadConfig()
    if err != nil {
        log.Panic("Error loading config file: ", err)
    }

	c := cron.New()

	c.AddFunc("@weekly", resetWeekly)
	c.AddFunc("@monthly", resetMonthly)
    c.AddFunc("@every 1m", func() { resetKuotaAbsen(config.AbsenceQuota) })
	c.AddFunc("@monthly", resetCurrentSalary)

	fmt.Println("Start cron job")
	c.Start()
	defer c.Stop()

	select {}
}
