package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"responses"
	"time"

	"auths"
	"configs"
	"forms"
	"models"

	"github.com/golang-jwt/jwt/v5"
)

const (
	CHECK_IN  = "Check-In"
	CHECK_OUT = "Check-Out"
)

func GetUserAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		loc, err := configs.Timezone{}.GetTimeZone()

		// Ensure no error get timezone location
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		curr_date_time := time.Now().In(loc)

		token := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		curr_month := int(curr_date_time.Month())

		// Create month
		months := [3]int{curr_month, curr_month - 1, curr_month - 2}

		var response_data responses.AttendanceReponseWrapperArray

		err = response_data.Create(months[:], int(token["id"].(float64)))

		// Ensure no error creating response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func GetUserAttendanceTodayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		token := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		loc, err := configs.Timezone{}.GetTimeZone()

		// Ensure no error get timezone location
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		curr_date_time := time.Now().In(loc).Format(time.DateOnly)

		var response_data responses.AttendanceReponse

		err = response_data.GetAttendanceOnDate(
			curr_date_time, int(token["id"].(float64)))

		// Ensure no error get user today latest attendance
		if err != nil && err != sql.ErrNoRows {
			log.Panic("Error get user today attendance: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		response_data.Date = curr_date_time

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func AttendUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		token := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var attend_form forms.AttendForm

		err := req_json.Decode(&attend_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		attend_form.Sanitize()

		valid, err := attend_form.Validate()

		if !valid {
			w.WriteHeader(http.StatusBadRequest)

			if err != forms.ErrTimeNotSync {
				json.NewEncoder(w).Encode(map[string]any{
					"error": "bad request",
				})

				return
			}

			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		user_id := int(token["id"].(float64))

		var response_data responses.AttendanceReponse

		err =
			response_data.GetAttendanceOnDate(attend_form.Date, user_id)

		// Ensure no error get attendance on date
		if err != nil && err != sql.ErrNoRows {
			log.Panic("Error get user today attendance: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Check whether new attendance is valid
		if response_data.CheckIn != nil && response_data.CheckOut != nil {
			w.WriteHeader(http.StatusAlreadyReported)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "already reported",
			})

			return
		}

		user, err :=
			models.User{}.GetUsingId(user_id)

		// Ensure no error fetching user data
		if err != nil {
			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Decide whether attendance type is check-in or check-out
		check_type := CHECK_IN

		if response_data.CheckIn != nil && response_data.CheckOut == nil {
			check_type = CHECK_OUT

			// Validate user check out record
			configs, err := models.ConfigJson{}.LoadConfig()

			if err != nil {
				log.Panic("Error load config json: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			min_check_out_time, err :=
				time.Parse(time.TimeOnly, configs.CheckOutTime)

			// Ensure no error parsing min check out time
			if err != nil {
				log.Panic("Error parse time: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			user_check_out_time, err :=
				time.Parse(time.TimeOnly, attend_form.Time)

			// Ensure no error parsing user check out time
			if err != nil {
				log.Panic("Error parse time: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			if min_check_out_time.After(user_check_out_time) {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "user attend too early",
				})

				return
			}
		}

		// Do update hour code
		
		attendance := models.Attendance{
			UserId:    user_id,
			Date_Time: attend_form.Date + " " + attend_form.Time,
			Type:      check_type,
		}

		err = attendance.Insert()

		// Ensure no error insert attendance
		if err != nil {
			log.Panic("Error insert attendance: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		response_data.Date = attend_form.Date
		response_data.CheckOut = &attend_form.Time

		log.Printf("User `%s` just %s", user.FullName, check_type)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": attendance,
		})
	}
}

func GetUserAttendanceStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		token := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		user_id := int(token["id"].(float64))

		stats, err :=
			models.AttendanceStats{}.GetUsingUserId(user_id)

		// Ensure no error get attendance stats
		if err != nil {
			log.Panic("Error get attendance stats: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": stats,
		})
	}
}
