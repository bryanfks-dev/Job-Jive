package forms

import (
	"configs"
	"errors"
	"log"
	"time"
)

type AttendForm struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

var (
	ErrTimeNotSync = errors.New("client and server time are not sync")
)

func (attend_form AttendForm) Validate() (bool, error) {
	// Try to parse datetime
	loc, err := configs.Timezone{}.GetTimeZone()

	// Ensure no error get timezone location
	if err != nil {
		return false, err
	}

	client_check_time, err := time.ParseInLocation(time.DateTime,
		attend_form.Date+" "+attend_form.Time, loc)

	// If there is an error on parsing datetime, that means
	// string that want to parse is not contains date or time
	if err != nil {
		return false, err
	}

	// Check if client and server time sync
	// This is crucial, because client could send request
	// for invalid date time

	server_check_time := time.Now().In(loc)

	time_diff := server_check_time.Sub(client_check_time)

	const threshold = 2 * time.Second

	if time_diff < -threshold || time_diff > threshold {
		log.Println("Client check time: ", client_check_time)
		log.Println("Server check Time: ", server_check_time)
		log.Println("Server and Client check time difference: ", time_diff)
		log.Println("WARNING: Server and client time are not sync")

		return false, ErrTimeNotSync
	}

	return true, nil
}
