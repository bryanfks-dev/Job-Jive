package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"auths"
	"db"
	"forms"
	"models"
	"responses"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	postMu sync.Mutex
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		users, err := models.User{}.Get()

		// Ensure no error fetching user data
		if err != nil {
			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data []responses.UserResponse

		for _, user := range users {
			var user_data responses.UserResponse

			err := user_data.Create(user)

			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "not found",
					})

					return
				}

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			response_data = append(response_data, user_data)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var user_form forms.UserForm

		err := req_json.Decode(&user_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		user_form.Sanitize()

		valid, err := user_form.ValidateCreate()

		// Ensure no error validating form
		if !valid {
			if err != forms.ErrInvalidEmail && err != forms.ErrEmailExist &&
				err != forms.ErrInvalidPhoneNumber && err != forms.ErrInvalidNIK {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		tx, err := db.Conn.Begin()

		if err != nil {
			log.Panic("Error starting database transaction: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		defer tx.Rollback()

		// Prepare new user field
		user := models.User{
			FullName:     user_form.FullName,
			Email:        user_form.Email,
			Password:     user_form.PhoneNumber, // Default value
			PhoneNumber:  user_form.PhoneNumber,
			DateOfBirth:  user_form.BirthDate,
			Address:      user_form.Address,
			NIK:          user_form.NIK,
			Gender:       user_form.Gender,
			DepartmentId: &user_form.DepartmentId,
			Photo:        user_form.Photo,
		}

		// Ensure no error inserting user
		id, err := user.Insert()

		if err != nil {
			log.Panic("Error insert user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Assign inserted id into current field
		user.Id = id

		err = models.Salary{
			UserId:  id,
			Initial: 0,
			Current: 0,
		}.Insert()

		// Ensure no error when inserting salary
		if err != nil {
			log.Panic("Error insert salary: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		configs, err := 
			models.ConfigJson{}.LoadConfig()

		// Ensure no error load conifg json file
		if err != nil {
			log.Panic("Error load config json: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		err = 
			models.AttendanceStats{}.Insert(id, configs.AbsenceQuota)

		// Ensure no error insert attendance stats
		if err != nil {
			log.Panic("Error insert attendance stats: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		if err := tx.Commit(); err != nil {
			log.Panic("Error committing to database: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data responses.UserResponse

		err = response_data.Create(user)

		// Ensure no error creating response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Printf("New user `%s` has been created", user.FullName)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Retrieve value from url
		id, err := strconv.Atoi(r.PathValue("id"))

		// Ensure user provides a valid record id
		if err != nil || id <= 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "invalid user id",
			})

			return
		}

		_, err = models.User{}.GetUsingId(id)

		// Ensure no error fetching user using id
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid user id",
				})

				return
			}

			log.Println("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		req_json := json.NewDecoder(r.Body)

		var user_form forms.UserForm

		err = req_json.Decode(&user_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		valid, err := user_form.ValidateUpdate(id)

		// Ensure no error validating form
		if !valid {
			if err != forms.ErrInvalidEmail && err != forms.ErrEmailExist &&
				err != forms.ErrInvalidPhoneNumber && err != forms.ErrInvalidNIK {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		user, err := models.User{}.GetUsingId(id)

		// Ensure no error when fetching data
		if err != nil {
			log.Panic("Error get user", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Update records
		user.FullName = user_form.FullName
		user.Email = user_form.Email
		user.DateOfBirth = user_form.BirthDate
		user.Address = user_form.Address
		user.NIK = user_form.NIK
		user.Gender = user_form.Gender
		user.PhoneNumber = user_form.PhoneNumber
		user.DepartmentId = &user_form.DepartmentId

		// Try update user credentials
		if strings.TrimSpace(user_form.NewPassword) != "" {
			hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(user_form.NewPassword), 11)

			// Ensure no error hashing new password
			if err != nil {
				log.Println("Error hashing password: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "server error",
				})

				return
			}

			user.Password = string(hashed_pwd)
		}

		err = user.Update()

		// Ensure no error when updating user
		if err != nil {
			log.Println("Error update user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "server error",
			})

			return
		}

		var response_data responses.UserResponse

		err = response_data.Create(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "server error",
			})

			return
		}

		log.Printf("User `%s` record has been updated\n", user.FullName)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": response_data,
		})
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Retrieve value from url
		id, err := strconv.Atoi(r.PathValue("id"))

		// Ensure user provide a valid record id
		if err != nil || id <= 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "invalid user id",
			})

			return
		}

		user, err := models.User{}.GetUsingId(id)

		// Ensure no error get user data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid user id",
				})

				return
			}

			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		err = user.Delete()

		// Ensure no error when deleting data
		if err != nil {
			log.Panic("Error delete user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Printf("User `%s` deleted", user.FullName)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": models.User{
				Id:    id,
				Photo: user.Photo,
			},
		})
	}
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		token := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		user, err :=
			models.User{}.GetUsingId(int(token["id"].(float64)))

		// Ensure no error when getting user information
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid user id",
				})

				return
			}

			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data responses.UserResponse

		err = response_data.Create(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
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

func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Retrieve value from url
		query := strings.TrimSpace(r.PathValue("query"))

		users, err := models.User{}.Search(query)

		// Ensure no erro search user
		if err != nil {
			log.Panic("Error search user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data []responses.UserResponse

		for _, user := range users {
			var user_data responses.UserResponse

			err := user_data.Create(user)

			// Ensure no error create response
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			response_data = append(response_data, user_data)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}
