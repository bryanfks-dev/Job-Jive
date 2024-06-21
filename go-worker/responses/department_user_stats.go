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
	curr_date_str := curr_date.Format(time.DateOnly)

	stmt := "WITH `UserAttendances` AS (SELECT u.User_ID, u.Full_Name, COALESCE(SUM(CASE WHEN a.Type = 'Check-Out' THEN 1 ELSE 0 END), 0) AS AttendanceCount, (DATEDIFF(?, DATE_SUB(?, INTERVAL 3 MONTH)) - COALESCE(SUM(CASE WHEN a.Type = 'Check-Out' THEN 1 ELSE 0 END), 0)) AS AbsenceCount FROM `users` u LEFT JOIN `attendances` a ON u.User_ID = a.User_ID AND a.Date_Time >= DATE_SUB(?, INTERVAL 3 MONTH) WHERE u.Department_ID = ? AND u.User_ID <> ? GROUP BY u.User_ID, u.Full_Name) SELECT User_ID, Full_Name, AttendanceCount, AbsenceCount FROM `UserAttendances` ORDER BY AttendanceCount DESC, Full_Name ASC LIMIT 3"

	row, err := db.Conn.Query(stmt, curr_date_str, curr_date_str, curr_date_str, department_id, manager_id)

	if err != nil {
		return []DepartmentUserStatsResponse{}, err
	}

	defer row.Close()

	var response_users []DepartmentUserStatsResponse

	for row.Next() {
		err := row.Scan(&department_user_stats_response.UserId, &department_user_stats_response.UserFullName, &department_user_stats_response.AttendCount, &department_user_stats_response.AbsenceCount)

		if err != nil {
			return []DepartmentUserStatsResponse{}, err
		}

		response_users = append(response_users, department_user_stats_response)
	}

	return response_users, nil
}

func (department_user_stats_response DepartmentUserStatsResponse) GetWorstDepartmentUserStats(curr_date time.Time, department_id int, manager_id int) ([]DepartmentUserStatsResponse, error) {
	curr_date_str := curr_date.Format(time.DateOnly)

	stmt := "WITH `UserAttendances` AS (SELECT u.User_ID, u.Full_Name, COALESCE(SUM(CASE WHEN a.Type = 'Check-Out' THEN 1 ELSE 0 END), 0) AS AttendanceCount, (DATEDIFF(?, DATE_SUB(?, INTERVAL 3 MONTH)) - COALESCE(SUM(CASE WHEN a.Type = 'Check-Out' THEN 1 ELSE 0 END), 0)) AS AbsenceCount FROM `users` u LEFT JOIN `attendances` a ON u.User_ID = a.User_ID AND a.Date_Time >= DATE_SUB(?, INTERVAL 3 MONTH) WHERE u.Department_ID = ? AND u.User_ID <> ? GROUP BY u.User_ID, u.Full_Name), `TopAttendees` AS (SELECT User_ID FROM `UserAttendances` ORDER BY AttendanceCount DESC, Full_Name ASC LIMIT 3) SELECT ua.User_ID, ua.Full_Name, ua.AttendanceCount, ua.AbsenceCount FROM `UserAttendances` ua LEFT JOIN `TopAttendees` ta ON ua.User_ID = ta.User_ID WHERE ta.User_ID IS NULL ORDER BY ua.AttendanceCount ASC, ua.Full_Name ASC LIMIT 3"

	row, err := db.Conn.Query(stmt, curr_date_str, curr_date_str, curr_date_str, department_id, manager_id)

	if err != nil {
		return []DepartmentUserStatsResponse{}, err
	}

	defer row.Close()

	var response_users []DepartmentUserStatsResponse

	for row.Next() {
		err := row.Scan(&department_user_stats_response.UserId, &department_user_stats_response.UserFullName, &department_user_stats_response.AttendCount, &department_user_stats_response.AbsenceCount)

		if err != nil {
			return []DepartmentUserStatsResponse{}, err
		}

		response_users = append(response_users, department_user_stats_response)
	}

	return response_users, nil
}
