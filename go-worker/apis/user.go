package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"sync"

	"auths"
	"db"
	"models"

	"github.com/golang-jwt/jwt/v5"
)

var (
	postMu sync.Mutex
)

type UserResponseData struct {
	Id          int     `json:"id"`
	FullName    string  `json:"full_name"`
	Status      string  `json:"status"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	BirthDate   string  `json:"birth_date"`
	PhoneNumber string  `json:"phone_number"`
	Gender      string  `json:"gender"`
	NIK         string  `json:"nik"`
	Department  string  `json:"department"`
	Photo       string  `json:"photo"`
	Salary      float64 `json:"salary"`
	FirstLogin  *string `json:"first_login"`
}

type UserFields struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"date_of_birth"`
	Address      string `json:"address"`
	NIK          string `json:"nik"`
	Gender       string `json:"gender"`
	DepartmentId int    `json:"department_id"`
	Photo        string `json:"photo"`
	NewPassword  string `json:"new_password"`
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
				Email:       user.Email,
				Address:     user.Address,
				BirthDate:   user.DateOfBirth,
				PhoneNumber: user.PhoneNumber,
				Gender:      user.Gender,
				NIK:         user.NIK,
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
				"message": "There is an invalid input field",
			})

			return
		}

		// Validate rules
		// Email validator
		_, err = mail.ParseAddress(user_fields.Email)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		// Phone number validator
		if len(user_fields.PhoneNumber) < 11 || len(user_fields.PhoneNumber) > 13 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		// NIK validator
		if len(user_fields.NIK) != 16 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		user := models.User{
			FullName:     user_fields.FullName,
			Email:        user_fields.Email,
			Password:     user_fields.PhoneNumber,
			PhoneNumber:  user_fields.PhoneNumber,
			DateOfBirth:  user_fields.BirthDate,
			Address:      user_fields.Address,
			NIK:          user_fields.NIK,
			Gender:       user_fields.Gender,
			DepartmentId: user_fields.DepartmentId,
			Photo:        user_fields.Photo,
		}

		// Check if email is unique
		_, err = models.User.GetUsingEmail(models.User{}, user.Email)

		// If err is nil, therefore there is an user that already have this email
		if err == nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Email already in use, please use other email",
			})

			return
		}

		id, err := models.User.Insert(user)

		// Ensure no error when inserting user
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		err = models.Salary.Insert(models.Salary{
			UserId:        id,
			InitialSalary: 0,
			CurrentSalary: 0,
		})

		// Ensure no error when inserting salary
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Printf("New user `%s` has been created", user.FullName)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Created",
		})
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_admin, res := auths.AdminMiddleware(r)

		if !valid_admin {
			json.NewEncoder(w).Encode(res)

			return
		}

		// Retrieve value from url
		id, err := strconv.Atoi(r.PathValue("id"))

		// Ensure user provide a valid record id
		if err != nil || id <= 0 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid user id",
			})

			return
		}

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var user_fields UserFields

		err = req_json.Decode(&user_fields)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		// Validate rules
		// Email validator
		_, err = mail.ParseAddress(user_fields.Email)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		// Phone number validator
		if len(user_fields.PhoneNumber) < 11 || len(user_fields.PhoneNumber) > 13 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		// NIK validator
		if len(user_fields.NIK) != 16 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "There is an invalid input field",
			})

			return
		}

		current_user, err :=
			models.User.GetUsingId(models.User{}, id)

		// Ensure no error fetching user data
		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusBadRequest,
					"message": "Invalid user id",
				})

				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Ensure user email is unique
		result, err :=
			models.User.GetUsingEmail(models.User{}, user_fields.Email)

		if err == nil && result.Id != current_user.Id {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Email already in use, please use other email",
			})

			return
		}

		// Database transaction
		tx, err := db.Conn.Begin()

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		defer tx.Rollback()

		user, err :=
			models.User.GetUsingId(models.User{}, id)

		// Ensure no error when fetching data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		old_name := user.FullName

		// Update records
		user.FullName = user_fields.FullName
		user.DateOfBirth = user_fields.BirthDate
		user.Address = user_fields.Address
		user.NIK = user_fields.NIK
		user.Gender = user_fields.Gender
		user.PhoneNumber = user_fields.PhoneNumber
		user.DepartmentId = user_fields.DepartmentId

		err = models.User.UpdateInformation(user)

		// Ensure no error when updating user
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Update login credentials
		if user_fields.Email != user.Email || strings.TrimSpace(user_fields.NewPassword) != "" {
			err := user.UpdateCredentials(user_fields.Email, user_fields.NewPassword)

			// Ensure no error updating user credentials
			if err != nil {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Server error",
				})
	
				return
			}
		}

		// Ensure no error commiting to database
		if err := tx.Commit(); err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Printf("User `%s` record has been updated\n", old_name)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Updated",
		})
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_admin, res := auths.AdminMiddleware(r)

		if !valid_admin {
			json.NewEncoder(w).Encode(res)

			return
		}

		// Retrieve value from url
		id, err := strconv.Atoi(r.PathValue("id"))

		// Ensure user provide a valid record id
		if err != nil || id <= 0 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		user, err :=
			models.User.GetUsingId(models.User{}, id)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		err = models.User.Delete(user)

		// Ensure no error when deleting data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println("User", user.FullName, "deleted")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Deleted",
		})
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

		token := res["jwt_claims"].(jwt.MapClaims)

		user, err :=
			models.User.GetUsingId(models.User{}, int(token["id"].(float64)))

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

		department_head, err :=
			models.DepartmentHead.GetUsingDepartmentId(models.DepartmentHead{}, user.DepartmentId)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Get user status
		user_status := "Employee"

		if department_head.ManagerId == &user.Id {
			user_status = "Manager"
		}

		// Get user department
		department, err :=
			models.Department.GetUsingId(models.Department{}, user.DepartmentId)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		response_data := UserResponseData{
			FullName:    user.FullName,
			Status:      user_status,
			Email:       user.Email,
			Address:     user.Address,
			BirthDate:   user.DateOfBirth,
			PhoneNumber: user.PhoneNumber,
			Gender:      user.Gender,
			NIK:         user.NIK,
			Department:  department.Name,
			Photo:       user.Photo,
			Salary:      0,
			FirstLogin:  user.FirstLogin,
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   response_data,
		})
	}
}
