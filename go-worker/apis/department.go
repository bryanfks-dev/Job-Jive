package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"auths"
	"db"
	"models"
)

type DepartmentResponseData struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ManagerId   *int   `json:"manager_id"`
	ManagerName string `json:"manager_name"`
}

type DepartmentFields struct {
	DepartmentName string `json:"department_name"`
	ManagerId      int    `json:"manager_id"`
}

func GetDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
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

		departments, err :=
			models.Department.Get(models.Department{})

		// Ensure no error when fetching department datas
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		var response_data []DepartmentResponseData

		for _, department := range departments {
			department_head, err :=
				models.DepartmentHead.GetUsingDepartmentId(
					models.DepartmentHead{}, department.Id)

			var user models.User

			// Ensure no error fetching department
			if err != nil && err != sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Server error",
				})

				return
			}

			if department_head.ManagerId != nil {
				user, err =
					models.User.GetUsingId(models.User{}, *department_head.ManagerId)

				// Ensure no error fetching user
				if err != nil {
					json.NewEncoder(w).Encode(map[string]interface{}{
						"status":  http.StatusInternalServerError,
						"message": "Server error",
					})

					return
				}
			}

			response_data = append(response_data, DepartmentResponseData{
				Id:          department.Id,
				Name:        department.Name,
				ManagerId:   &user.Id,
				ManagerName: user.FullName,
			})
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   response_data,
		})
	}
}

func CreateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
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

		var department_fields DepartmentFields

		err := req_json.Decode(&department_fields)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		id, err := models.Department.Insert(models.Department{
			Name: department_fields.DepartmentName,
		})

		// Ensure no error when inserting department
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		err = models.DepartmentHead.Insert(models.DepartmentHead{
			DepartmentId: int(id),
			ManagerId:    nil,
		})

		// Ensure no error when inserting department_head
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println("New department", department_fields.DepartmentName, "has been created")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Created",
		})
	}
}

func UpdateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
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
				"message": "Bad request",
			})

			return
		}

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var department_fields DepartmentFields

		err = req_json.Decode(&department_fields)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
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

		department, err :=
			models.Department.GetUsingId(models.Department{}, id)

		// Ensure no error when fetching data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		department_head, err :=
			models.DepartmentHead.GetUsingDepartmentId(models.DepartmentHead{}, department.Id)

		// Ensure no error when fetching data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Update records
		department.Name = department_fields.DepartmentName
		department_head.ManagerId = &department_fields.ManagerId

		err = models.Department.Update(department)

		// Ensure no error when updating depatrment
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		err = models.DepartmentHead.Update(department_head)

		// Ensure no error when updating department
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		if err := tx.Commit(); err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println("Department record", department.Name, "updated")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Updated",
		})
	}
}

func DeleteDepartmentHandler(w http.ResponseWriter, r *http.Request) {
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

		department, err :=
			models.Department.GetUsingId(models.Department{}, id)

		// Ensure no error when fetching data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		err = models.Department.Delete(department)

		// Ensure no error when deleting data
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println("Department", department.Name, "deleted")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Deleted",
		})
	}
}
