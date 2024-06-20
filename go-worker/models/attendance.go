package models

import (
	"db"
	"time"
)

type Attendance struct {
	UserId    int
	Date_Time string
	Type      string
}

func (attendance Attendance) Insert() error {
	stmt := "INSERT INTO `attendances` (User_ID, Date_Time, Type) VALUES(?, ?, ?)"

	_, err := db.Conn.Exec(stmt,
		attendance.UserId, attendance.Date_Time, attendance.Type)

	return err
}

func (attendance Attendance) GetAttendancePerWeek(user_id int) (int, error) {
	get_config, err := ConfigJson{}.LoadConfig()

	if err != nil {
		return 0, err
	}

	check_in_time := get_config.CheckInTime

	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -7)

	stmt := `SELECT COUNT(*) FROM attendances WHERE User_ID = ? AND Type = 'Check-In' AND
		DATE(Date_Time) >= ? AND DATE(Date_Time) <= ? AND TIME(Date_Time) > ?`

	count := 0

	err =
		db.Conn.QueryRow(stmt, user_id, sevenDaysAgo.Format(time.DateOnly), now.Format("2006-01-02"), check_in_time).Scan(&count)

	return count, err
}
