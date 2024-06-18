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

func GetUserTodayLatestAttendanceHandler(w http.ResponseWriter, r *http.Request) {
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

		curr_date_time := time.Now().In(loc)

		attendance, err :=
			models.Attendance{}.GetTodayLatestAttendance(
				curr_date_time.Format(time.DateOnly), int(token["id"].(float64)))

		// Ensure no error get user today latest attendance
		if err != nil && err != sql.ErrNoRows {
			log.Panic("Error get user latest attendance: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": attendance,
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
		check_type := "Check-In"

		last_check, err :=
			models.Attendance{}.GetTodayLatestAttendance(attend_form.Date, user_id)

		if err != sql.ErrNoRows && last_check.Type == "Check-In" {
			check_type = "Check-Out"
		}

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

		log.Printf("User `%s` just %s", user.FullName, check_type)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": attendance,
		})
	}
}
