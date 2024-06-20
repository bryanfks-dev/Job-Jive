package responses

import (
	"log"
	"strconv"
	"strings"
	"time"

	"configs"
	"db"
	"models"
)

type AttendanceReponseWrapper struct {
	UserId  int                     `json:"user_id"`
	Month   int                     `json:"month"`
	Records AttendanceResponseArray `json:"records"`
}

type AttendanceReponse struct {
	UserId   int     `json:"user_id,omitempty"`
	Date     string  `json:"date"`
	CheckIn  *string `json:"check_in_time"`
	CheckOut *string `json:"check_out_time"`
}

type (
	AttendanceReponseWrapperArray [3]AttendanceReponseWrapper
	AttendanceResponseArray       []AttendanceReponse
)

func (attendance_response *AttendanceReponse) GetAttendanceOnDate(date string, user_id int) error {
	attendance_response.UserId = user_id

	stmt := "SELECT TIME(a1.Date_Time), TIME(a2.Date_Time) FROM `attendances` a1 LEFT JOIN `attendances` a2 ON DATE(a2.Date_Time) = DATE(a1.Date_Time) AND a2.User_ID = a1.User_ID AND a2.Type = 'Check-Out' WHERE a1.Type = 'Check-In' AND a1.User_ID = ? AND DATE(a1.Date_Time) = ?"

	err := db.Conn.QueryRow(stmt, user_id, date).
		Scan(&attendance_response.CheckIn, &attendance_response.CheckOut)

	return err
}

func (attendace_response_array *AttendanceResponseArray) GetAttendaceUsingMonth(curr_date time.Time, target_month int, user_id int) error {
	var month_digit string

	if target_month < 10 {
		month_digit = "0" + strconv.Itoa(target_month)
	} else {
		month_digit = strconv.Itoa(target_month)
	}

	first_date :=
		strconv.Itoa(curr_date.Year()) + "-" + month_digit + "-01"

	curr_date_formatted := curr_date.Format(time.DateOnly)

	stmt1 := "WITH RECURSIVE dateList AS (SELECT '" + first_date + "' AS Date UNION ALL (SELECT ADDDATE(Date, INTERVAL 1 DAY) FROM dateList WHERE Date < LEAST('" + curr_date_formatted + "', LAST_DAY('" + first_date + "'))))"
	stmt2 := `SELECT d.Date, TIME(ci.Date_Time), TIME(co.Date_Time) FROM dateList d LEFT JOIN attendances ci ON DATE(ci.Date_Time) = d.Date AND ci.Type = "Check-In" AND ci.User_ID = ? LEFT JOIN attendances co ON DATE(co.Date_Time) = d.Date AND co.Type = "Check-Out" AND co.User_ID = ? ORDER BY d.Date DESC`

	row, err := db.Conn.Query(stmt1+stmt2, user_id, user_id)

	if err != nil {
		return err
	}

	defer row.Close()

	var attendace_response AttendanceReponse

	for row.Next() {
		err := row.Scan(
			&attendace_response.Date,
			&attendace_response.CheckIn,
			&attendace_response.CheckOut)

		if err != nil {
			return err
		}

		*attendace_response_array = append(*attendace_response_array, attendace_response)
	}

	return nil
}

func (attendance_wrappers *AttendanceReponseWrapperArray) Create(months []int, user_id int) error {
	for idx, month := range months {
		curr_wrapper := &attendance_wrappers[idx]

		var attendance_responses AttendanceResponseArray

		err := attendance_responses.createRecord(month, user_id)

		// Ensure no error create record
		if err != nil {
			return err
		}

		curr_wrapper.UserId = user_id
		curr_wrapper.Month = month
		curr_wrapper.Records = attendance_responses
	}

	return nil
}

func (attendance_responses_array *AttendanceResponseArray) createRecord(month int, user_id int) error {
	tz, err := configs.Timezone{}.GetTimeZone()

	if err != nil {
		return err
	}

	curr_date := time.Now().In(tz)

	err = (*attendance_responses_array).GetAttendaceUsingMonth(curr_date, month, user_id)

	// Ensure no error get user attendance
	if err != nil {
		log.Panic("Error get user attendance in month "+strconv.Itoa(month)+": ", err.Error())

		return err
	}

	return nil
}

func (attendance_response *AttendanceReponse) CreateSingle(attendance models.Attendance) {
	date_time := strings.Split(attendance.Date_Time, " ")

	attendance_response.UserId = attendance.UserId
	attendance_response.Date = date_time[0]

	if attendance.Type == "Check-In" {
		attendance_response.CheckIn = &date_time[1]
	} else {
		attendance_response.CheckOut = &date_time[1]
	}
}
