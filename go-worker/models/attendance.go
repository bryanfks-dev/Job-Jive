package models

import "db"

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
