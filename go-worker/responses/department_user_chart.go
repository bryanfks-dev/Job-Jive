package responses

import "db"

type DepartmentUserChartResponse struct {
	AttendCount  int `json:"attend_count"`
	AbsenceCount int `json:"absence_count"`
}

func (department_user_chart_response *DepartmentUserChartResponse) GetDepartmentUsersAttendance(today string, department_id int, manager_id int) (int, error) {
	stmt := "SELECT COUNT(*), COALESCE(SUM(a.Type = 'Check-Out'), 0) FROM `users` u LEFT JOIN `attendances` a ON a.User_ID = u.User_ID WHERE DATE(a.Date_Time) = ? AND u.Department_ID = ? AND u.User_ID <> ?"

	employee_count := 0

	err :=
		db.Conn.QueryRow(stmt, today, department_id, manager_id).
			Scan(&employee_count, &department_user_chart_response.AttendCount)

	if err != nil {
		return 0, err
	}

	return employee_count - 1, nil
}

func (department_user_chart_response *DepartmentUserChartResponse) GetDepartmentUsersAttendanceBetween(start_date string, end_date string, department_id int, manager_id int) (int, error) {
	stmt := "SELECT COUNT(*), COALESCE(SUM(a.Type = 'Check-Out'), 0) FROM `users` u LEFT JOIN `attendances` a ON a.User_ID = u.User_ID WHERE (DATE(a.Date_Time) >= ? AND DATE(a.Date_Time) <= ?) AND u.Department_ID = ? AND u.User_ID <> ?"

	employee_count := 0

	err :=
		db.Conn.QueryRow(stmt, start_date, end_date, department_id, manager_id).
			Scan(&employee_count, &department_user_chart_response.AttendCount)

	if err != nil {
		return 0, err
	}

	return employee_count - 1, nil
}
