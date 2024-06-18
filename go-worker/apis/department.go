package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"db"
	"forms"
	"models"
	"responses"
)

func GetDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		departments, err :=
			models.Department{}.Get()

		// Ensure no error when fetching department datas
		if err != nil {
			log.Panic("Error get department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data []responses.DepartmentResponse

		for _, department := range departments {
			var department_data responses.DepartmentResponse

			err := department_data.Create(department)

			// Ensure no error create response
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "manager user not found",
					})

					return
				}

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			response_data = append(response_data, department_data)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func CreateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var department_form forms.DepartmentForm

		err := req_json.Decode(&department_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		department_form.Sanitize()

		valid, err := department_form.ValidateCreate()

		if !valid {
			if err != forms.ErrDepartmentNameExist {
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

		// Ensure no error starting database transaction
		if err != nil {
			log.Panic("Error starting database transaction: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		defer tx.Rollback()

		var department = models.Department{
			Name: department_form.DepartmentName,
		}

		id, err := department.Insert()

		// Ensure no error when inserting department
		if err != nil {
			log.Panic("Error insert department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Assign department id
		department.Id = id

		err =
			models.DepartmentHead{
				DepartmentId: id,
				ManagerId:    nil,
			}.Insert()

		// Ensure no error inserting department_head
		if err != nil {
			log.Panic("Error insert departmend_head: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Ensure no error commiting to database
		if err := tx.Commit(); err != nil {
			log.Panic("Error commiting to database: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data responses.DepartmentResponse

		err = response_data.Create(department)

		// Ensure no error create response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Printf("New department `%s` has been created", department_form.DepartmentName)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func UpdateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
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
				"error": "invalid department id",
			})

			return
		}

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var department_form forms.DepartmentForm

		err = req_json.Decode(&department_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		department, err :=
			models.Department{}.GetUsingId(id)

		// Ensure no error fetching department data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid department id",
				})

				return
			}

			log.Panic("Error get department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		valid, err := department_form.ValidateUpdate()

		if !valid {
			if err != forms.ErrManagerIdExist {
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

		department_head, err :=
			models.DepartmentHead{}.GetUsingDepartmentId(department.Id)

		// Ensure no error when fetching department_head data
		if err != nil {
			log.Panic("Error get department_head: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Update records
		department_head.ManagerId = &department_form.ManagerId

		err = department_head.Update()

		// Ensure no error when updating depatrmen
		if err != nil {
			log.Panic("Error update department_head: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data responses.DepartmentResponse

		err = response_data.Create(department)

		// Ensure no error creating response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Printf("Department `%s` record has been updated\n", department.Name)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}

func DeleteDepartmentHandler(w http.ResponseWriter, r *http.Request) {
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
				"error": "invalid department id",
			})

			return
		}

		department, err :=
			models.Department{}.GetUsingId(id)

		// Ensure no error when fetching department data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid department id",
				})

				return
			}

			log.Panic("Error get department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		err = department.Delete()

		// Ensure no error when deleting data
		if err != nil {
			log.Panic("Error delete department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Printf("Department `%s` deleted", department.Name)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": models.Department{
				Id: department.Id,
			},
		})
	}
}

func SearchDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Retrieve value from url
		query := strings.TrimSpace(r.PathValue("query"))

		departments, err := models.Department{}.Search(query)

		// Ensure no error search department
		if err != nil {
			log.Panic("Error search department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		var response_data []responses.DepartmentResponse

		for _, department := range departments {
			var department_data responses.DepartmentResponse

			err := department_data.Create(department)

			// Ensure no error create response
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "manager user not found",
					})

					return
				}

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			response_data = append(response_data, department_data)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}
