package apis

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"

	"auths"
	"models"
)

var (
	postMu sync.Mutex
)

type UserResponseData struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	BirthDate   string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Department  string `json:"department"`
}

type UserFields struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	BirthDate    string `json:"date_of_birth"`
	Address      string `json:"address"`
	NIK          string `json:"nik"`
	Gender       string `json:"gender"`
	DepartmentId int    `json:"department_id"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_admin, res := auths.AdminMiddleware(r)

		if !valid_admin {
			json.NewEncoder(w).Encode(res)

			return
		}

		users, err := models.User.GetUsers(models.User{})

		// Ensure no error fetching employees data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		var response_data []UserResponseData

		for _, user := range users {
			department, err :=
				models.Department.GetUsingId(models.Department{}, user.DepartmentId)

			if err != nil {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Server error",
				})

				return
			}

			response_data = append(response_data, UserResponseData{
				Id:          user.Id,
				FullName:    user.FullName,
				BirthDate:   user.DateOfBirth,
				PhoneNumber: user.PhoneNumber,
				Gender:      user.Gender,
				Department:  department.Name,
			})
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   response_data,
		})
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_admin, res := auths.AdminMiddleware(r)

		if !valid_admin {
			json.NewEncoder(w).Encode(res)

			return
		}

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var user_fields UserFields

		err := req_json.Decode(&user_fields)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		
	}
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_user, res := auths.UserMiddleware(r)

		if !valid_user {
			json.NewEncoder(w).Encode(res)

			return
		}

		user, err :=
			models.User.GetUsingId(models.User{}, res["id"].(int))

		// Ensure no error when getting user information
		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Invalid token value",
				})

				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   user,
		})
	}
}
