package models

import "db"

type Attendance struct {
	UserId    int    `json:"user_id"`
	Date_Time string `json:"date_time"`
	Type      string `json:"type"`
}

func (attendance Attendance) GetTodayLatestAttendance(date string, user_id int) (Attendance, error) {
	stmt := "SELECT * FROM `attendances` WHERE User_ID = ? AND DATE(Date_Time) = ?"

	err := db.Conn.QueryRow(stmt, user_id, date).
		Scan(&attendance.UserId,
			&attendance.Date_Time,
			&attendance.Type)

	return attendance, err
}

func (attendance Attendance) Insert() error {
	stmt := "INSERT INTO `attendances` (User_ID, Date_Time, Type) VALUES(?, ?, ?)"

	_, err := db.Conn.Exec(stmt,
		attendance.UserId, attendance.Date_Time, attendance.Type)

	return err
}
