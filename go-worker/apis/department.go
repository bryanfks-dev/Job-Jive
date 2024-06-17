package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "manager user not found",
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

		// Ensure the department name is unique
		_, err =
			models.Department{}.GetUsingDepartmentName(department_form.DepartmentName)

		// Emsure no error fetching department data
		if err != nil {
			if err != sql.ErrNoRows {
				log.Panic("Error get department using name", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "department name exist, user other name",
			})
		}

		tx, err := db.Conn.Begin()

		// Ensure no error starting database transaction
		if err != nil {
			log.Panic("Error starting database transaction", err.Error())

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
			log.Panic("Error insert department", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		department.Id = id

		err =
			models.DepartmentHead{
				DepartmentId: id,
				ManagerId:    nil,
			}.Insert()

		// Ensure no error inserting department_head
		if err != nil {
			log.Panic("Error insert departmend_head", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Ensure no error commiting to database
		if err := tx.Commit(); err != nil {
			log.Panic("Error commiting to database", err.Error())

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
			log.Panic("Error create response", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Println("New department", department_form.DepartmentName, "has been created")

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
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
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
			log.Panic("Error get department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		valid, err := department_form.Validate()

		if !valid {
			if err != forms.ErrManagerIdExist {
				log.Panic("Error validate department: ", err.Error())

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

		log.Printf("Department %s record has been updated\n", department.Name)

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
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "Bad request",
			})

			return
		}

		department, err :=
			models.Department{}.GetUsingId(id)

		// Ensure no error when fetching department data
		if err != nil {
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

		log.Println("Department", department.Name, "deleted")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"id": id,
			},
		})
	}
}
