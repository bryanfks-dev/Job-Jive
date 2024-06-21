package responses

import (
	"time"

	"db"
)

type DepartmentUserStatsWrapperResponse struct {
	BestUsersAttendance  []DepartmentUserStatsResponse `json:"best_users_attendance"`
	WorstUsersAttendance []DepartmentUserStatsResponse `json:"worst_users_attendance"`
}

type DepartmentUserStatsResponse struct {
	UserId       int    `json:"user_id"`
	UserFullName string `json:"user_full_name"`
	AttendCount  int    `json:"attend_count"`
	AbsenceCount int    `json:"absence_count"`
}

func (department_user_stats_response DepartmentUserStatsResponse) GetBestDepartmentUserStats(curr_date time.Time, department_id int, manager_id int) ([]DepartmentUserStatsResponse, error) {
	days_count := int(curr_date.Sub(curr_date.AddDate(0, -3, 0)).Hours() / 24)

	stmt := "WITH UserAttendances AS (SELECT u.User_ID, u.Full_Name, COALESCE(SUM(a.Type = 'Check-Out'), 0) AS att FROM users u LEFT JOIN attendances a ON a.User_ID = u.User_ID AND DATE(a.Date_Time) >= DATE_SUB(?, INTERVAL 3 MONTH) AND u.Department_ID = ? AND u.User_ID <> ? GROUP BY u.User_ID, u.Full_Name) SELECT * FROM UserAttendances ua ORDER BY ua.att, ua.Full_Name DESC LIMIT 3"

	row, err := db.Conn.Query(stmt, curr_date.Format(time.DateOnly), department_id, manager_id)

	if err != nil {
		return []DepartmentUserStatsResponse{}, err
	}

	defer row.Close()

	var response_users []DepartmentUserStatsResponse

	for row.Next() {
		err := row.Scan(&department_user_stats_response.UserId, &department_user_stats_response.UserFullName, &department_user_stats_response.AttendCount)

		if err != nil {
			return []DepartmentUserStatsResponse{}, err
		}

		department_user_stats_response.AbsenceCount = days_count - department_user_stats_response.AttendCount

		response_users = append(response_users, department_user_stats_response)
	}

	return response_users, nil
}

func (department_user_stats_response DepartmentUserStatsResponse) GetWorstDepartmentUserStats(curr_date time.Time, department_id int, manager_id int) ([]DepartmentUserStatsResponse, error) {
	days_count := int(curr_date.Sub(curr_date.AddDate(0, -3, 0)).Hours() / 24)

	stmt := "WITH UserAttendances AS (SELECT u.User_ID, u.Full_Name, COALESCE(SUM(a.Type = 'Check-Out'), 0) AS att FROM users u LEFT JOIN attendances a ON a.User_ID = u.User_ID AND DATE(a.Date_Time) >= DATE_SUB(?, INTERVAL 3 MONTH) WHERE u.Department_ID = ? AND u.User_ID <> ? GROUP BY u.User_ID, u.Full_Name), BestUserAttendances AS (SELECT * FROM UserAttendances ua ORDER BY ua.att, ua.Full_Name DESC LIMIT 3) SELECT ua.* FROM UserAttendances ua, BestUserAttendances bua WHERE ua.User_ID NOT IN (bua.User_ID) ORDER BY ua.att, ua.Full_Name ASC LIMIT 3"

	row, err := db.Conn.Query(stmt, curr_date.Format(time.DateOnly), department_id, manager_id)

	if err != nil {
		return []DepartmentUserStatsResponse{}, err
	}

	defer row.Close()

	var response_users []DepartmentUserStatsResponse

	for row.Next() {
		err := row.Scan(&department_user_stats_response.UserId, &department_user_stats_response.UserFullName, &department_user_stats_response.AttendCount)

		if err != nil {
			return []DepartmentUserStatsResponse{}, err
		}

		department_user_stats_response.AbsenceCount = days_count - department_user_stats_response.AttendCount

		response_users = append(response_users, department_user_stats_response)
	}

	return response_users, nil
}
