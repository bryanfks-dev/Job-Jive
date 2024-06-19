package apis

import (
	"auths"
	"database/sql"
	"encoding/json"
	"forms"
	"log"
	"net/http"
	"strconv"

	"models"
	"responses"

	"github.com/golang-jwt/jwt/v5"
)

type context_key string

func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		jwt_claims := r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		get_userID := int(jwt_claims["id"].(float64))
		get_departmentID, err := models.User{}.GetDepartmentId(get_userID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}

		// Fetch manager data
		manager, err := models.User{}.GetUsingId(get_userID)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "Invalid user",
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}

		// Initialize response data
		var response_data []responses.UserResponse

		// Add manager to response data
		var manager_data responses.UserResponse
		err = manager_data.Create(manager)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}
		response_data = append(response_data, manager_data)

		// Fetch employees data
		users, err := models.User{}.GetEmployees(get_userID, get_departmentID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}

		// Add employees to response data
		for _, user := range users {
			var user_data responses.UserResponse
			err := user_data.Create(user)
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "server error",
					})
					return
				}
				log.Panic("Error create response: ", err.Error())
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

// UpdateEmployeeHandler handles PUT request to update employee's salary
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT
	if r.Method == http.MethodPut {

		// Lock the mutex to prevent multiple requests from accessing the same resource at the same time
		postMu.Lock()
		defer postMu.Unlock()

		// Set the content type of the response to application/json
		w.Header().Set("Content-Type", "application/json")

		// Retrieve the id from the URL path and convert it to an integer
		id, err := strconv.Atoi(r.PathValue("id"))

		// Check if the id is valid
		if err != nil || id <= 0 {
			// If the id is invalid, return a bad request error
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Invalid user id",
			})
			return
		}

		// Check if there is a salary record associated with the id
		_, err = models.Salary{}.GetUsingUserId(id)

		// Check if there is an error fetching the salary record
		if err != nil {
			// If there is no salary record, return a bad request error
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "Invalid user id",
				})
				return
			}

			// If there is an error, log it and return a server error
			log.Println("Error get user: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}

		// Decode the request body into a SalaryForm struct
		req_json := json.NewDecoder(r.Body)

		var salaryForm forms.SalaryForm

		err = req_json.Decode(&salaryForm)

		// log.Println(salaryForm)
		// log.Println(req_json)
		// log.Println(id)

		// Check if there is an error decoding the request body
		if err != nil {
			// If there is an error, return a bad request error
			log.Panic("Error decode: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})
			return
		}

		// Retrieve the salary record associated with the id
		salary, err := models.Salary{}.GetUsingUserId(id)

		// Check if there is an error retrieving the salary record
		if err != nil {
			// If there is an error, log it and return a server error
			log.Panic("Error get user", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Update the initial and current values of the salary record
		salary.Initial = salaryForm.Initial
		salary.Current = salaryForm.Current

		// Update the salary record in the database
		err = salary.Update()

		// Check if there is an error updating the salary record
		if err != nil {
			// If there is an error, log it and return a server error
			log.Println("Error update salary", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "server error",
			})
			return
		}

		// Create a SalaryResponse struct from the updated salary record
		var response_data responses.SalaryResponse
		err = response_data.Create(salary)
		if err != nil {
			// If there is an error, return a server error
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})
			return
		}

		// Log that the salary update was successful and return a success response
		log.Println("Success update salary")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})

	}
}
