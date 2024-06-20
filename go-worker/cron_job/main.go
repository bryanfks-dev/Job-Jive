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

func resetAttendanceStatsWeekly() {
	log.Println("Running weekly reset attendance stats for current week hours")

	tx, err := db.Conn.Begin()

	if err != nil {
		log.Panic("Error start database transaction: ", err.Error())

		return
	}

	defer tx.Rollback()

	_, err = tx.Exec("UPDATE `attendance_statistics` SET Current_Week_Hours = 0")

	if err != nil {
		log.Panic("Error update `attendance_statistic` current_week_hours: ", err.Error())

		return
	}

	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}

	log.Println("Attendance stats weekly hours reset complete")
}

func resetAttendanceStatsMonthly() {
	log.Println("Running monthly reset attendance stats for current month hours")

	tx, err := db.Conn.Begin()

	if err != nil {
		log.Panic("Error start database transaction: ", err.Error())

		return
	}

	defer tx.Rollback()

	_, err = tx.Exec("UPDATE `attendance_statistics` SET Current_Month_Hours = 0")

	if err != nil {
		log.Panic("Error update `attendance_statistic` current_month_hours: ", err.Error())

		return
	}

	if err := tx.Commit(); err != nil {
		log.Panic("Error commit to database: ", err.Error())

		return
	}

	log.Println("Attendance stats monthly hours reset complete")
}

func resetAttendanceStatsYearly() {
	log.Println("Running yearly reset attendance stats for annual leaves")

	// Load the config file for AbsenceQuota
	var config models.ConfigJson
	config, err := config.LoadConfig()

	// Ensure no error load config json
	if err != nil {
		panic(err.Error())
	}

	tx, err := db.Conn.Begin()

	if err != nil {
		log.Panic("Error start database transaction: ", err.Error())

		return
	}

	defer tx.Rollback()

	_, err =
		tx.Exec("UPDATE `attendance_statistics` SET Annual_Leaves = ?", config.AbsenceQuota)

	if err != nil {
		log.Panic("Error update `attendance_statistic` annual_leaves: ", err.Error())

		return
	}

	if err := tx.Commit(); err != nil {
		log.Panic("Error commit to database: ", err.Error())

		return
	}

	log.Println("Attendance stats yearly annual leaves reset complete")
}

func resetCurrentSalary() {
	log.Println("Running monthly reset for current salary")

	tx, err := db.Conn.Begin()

	if err != nil {
		log.Panic("Error start database transaction: ", err.Error())

		return
	}

	defer tx.Rollback()

	_, err = tx.Exec("UPDATE `salaries` SET Current_Salary = Initial_Salary")

	if err != nil {
		log.Panic("Error update `salaries` current_salary: ", err.Error())

		return
	}

	if err := tx.Commit(); err != nil {
		log.Panic("Error commit to database: ", err.Error())

		return
	}

	log.Println("Monthly reset for current salary complete")
}

func subCurrentSalary() {
	log.Println("Running daily subtraction for current salary")

	// Load the config file for AbsenceQuota
	var config models.ConfigJson
	config, err := config.LoadConfig()

	// Ensure no error load config json
	if err != nil {
		panic(err.Error())
	}

	loc, err := configs.Timezone{}.GetTimeZone()

	// Ensure no error load timezone location
	if err != nil {
		panic(err.Error())
	}

	today_date := time.Now().In(loc).Format(time.DateOnly)

	tx, err := db.Conn.Begin()

	if err != nil {
		log.Panic("Error start database transaction: ", err.Error())

		return
	}

	defer tx.Rollback()

	_, err = tx.Exec("UPDATE `salaries` s SET s.Current_Salary = GREATEST(s.Current_Salary - ?, 0) WHERE s.User_ID IN (SELECT u.User_ID FROM `users` u LEFT JOIN `attendances` a ON a.User_ID = u.User_ID AND DATE(a.Date_Time) = ? WHERE (a.Type = 'Check-In' AND TIME(a.Date_Time) > ?) OR a.Type IS NULL)", config.DeductionAmounts, today_date, config.CheckInTime)

	if err != nil {
		log.Panic("Error update salaries: ", err.Error())

		return
	}

	if err := tx.Commit(); err != nil {
		log.Panic("Error commit to database: ", err.Error())

		return
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

	c := cron.New()

	c.AddFunc("@weekly", resetAttendanceStatsWeekly)
	c.AddFunc("@monthly", resetAttendanceStatsMonthly)
	c.AddFunc("@yearly", resetAttendanceStatsYearly)
	c.AddFunc("@monthly", resetCurrentSalary)
	c.AddFunc("0 59 23 * * *", subCurrentSalary)

	fmt.Println("Start cron job")

	c.Start()
	defer c.Stop()

	select {}
}
