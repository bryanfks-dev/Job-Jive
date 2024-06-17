package responses

import (
	"log"
	"strconv"

	"db"
)

type AttendanceReponseWrapper struct {
	Month   int                     `json:"month"`
	Records AttendanceResponseArray `json:"records"`
}

type AttendanceReponse struct {
	Date     string  `json:"date"`
	CheckIn  string  `json:"check_in_time"`
	CheckOut *string `json:"check_out_time"`
}

type (
	AttendanceReponseWrapperArray [3]AttendanceReponseWrapper
	AttendanceResponseArray       []AttendanceReponse
)

func (attendace_response_array *AttendanceResponseArray) GetAttendaceUsingMonth(month int, user_id int) error {
	stmt := `SELECT DATE(a1.Date_Time) Date, TIME(a1.Date_Time) Check_In, CASE WHEN a2.Type = "Check-Out" THEN TIME(a2.Date_Time) ELSE NULL END Check_Out FROM attendances a1 LEFT JOIN attendances a2 ON a1.User_ID = a2.User_ID AND DATE(a2.Date_Time) = DATE(a1.Date_Time) AND a2.Type = "Check-Out" WHERE a1.Type = "Check-In" AND a1.User_ID = ? AND MONTH(a1.Date_Time) = ? ORDER BY a1.Date_Time DESC`

	row, err := db.Conn.Query(stmt, user_id, month)

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
		curr_wrapper := &(*attendance_wrappers)[idx]

		var attendance_responses AttendanceResponseArray
		
		err := attendance_responses.createRecord(month, user_id)

		// Ensure no error create record
		if err != nil {
			return err
		}

		curr_wrapper.Month = month
		curr_wrapper.Records = attendance_responses
	}

	return nil
}

func (attendance_responses_array *AttendanceResponseArray) createRecord(month int, user_id int) error {
	err := (*attendance_responses_array).GetAttendaceUsingMonth(month, user_id)

	// Ensure no error get user attendance
	if err != nil {
		log.Panic("Error get user attendance in month "+strconv.Itoa(month)+": ", err.Error())

		return err
	}

	return nil
}
