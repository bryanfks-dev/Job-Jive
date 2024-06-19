package models

import "db"

type Attendance struct {
	UserId    int    `json:"user_id"`
	Date_Time string `json:"date_time"`
	Type      string `json:"type"`
}

func (attendance Attendance) GetCount(user_id int) /* (int, error) */ {
	// Stuck bgt
}

func (attendance Attendance) Insert() error {
	stmt := "INSERT INTO `attendances` (User_ID, Date_Time, Type) VALUES(?, ?, ?)"

	_, err := db.Conn.Exec(stmt,
		attendance.UserId, attendance.Date_Time, attendance.Type)

	return err
}
