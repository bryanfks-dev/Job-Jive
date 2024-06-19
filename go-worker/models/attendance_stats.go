package models

import "db"

type AttendanceStats struct {
	UserID            int
	CurrentWeekHours  int
	CurrentMonthHours int
	AnnualLeaves      int
}

func (attendance_stats AttendanceStats) GetUsingUserId(user_id int) (AttendanceStats, error) {
	stmt := "SELECT * FROM `attendance_statistics` WHERE User_ID = ?"

	err := db.Conn.QueryRow(stmt, user_id).
		Scan(&attendance_stats.UserID,
			&attendance_stats.CurrentWeekHours,
			&attendance_stats.CurrentMonthHours,
			&attendance_stats.AnnualLeaves)

	return attendance_stats, err
}

func (attendance_stats AttendanceStats) Insert(user_id int, max_annual_leaves int) error {
	stmt := "INSERT INTO `attendance_statistics` (User_ID, Current_Week_Hours, Current_Month_Hours, Annual_Leaves) VALUES (?, ?, ?, ?)"

	_, err := db.Conn.Exec(stmt, user_id, 0, 0, max_annual_leaves)

	return err
}

func (attendance_stats AttendanceStats) UpdateHours(user_id int, value int) error {
	stmt := "UPDATE `attendance_statistics` SET Current_Week_Hours = Current_Week_Hours + ?, Current_Month_Hours = Current_Month_Hours + ? WHERE User_ID = ?"

	_, err := db.Conn.Exec(stmt, value, value, user_id)

	return err
}

func (attendance_stats AttendanceStats) UpdateAnnualLeaves(user_id int) error {
	stmt := "UPDATE `attendance_statistics` SET Annual_Leaves = Annual_Leaves - 1 WHERE User_ID = ?"

	_, err := db.Conn.Exec(stmt, user_id)

	return err
}
