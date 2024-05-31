package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"auths"
	"models"
)

type DepartmentFields struct {
	DepartmentName string `json:"department-name"`
	ManagerId      int    `json:"manager-id"`
}

func CreateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Validate token
		token_valid, res := auths.AuthorizedToken(r)

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		if !token_valid {
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

		err = models.Department.Insert(models.Department{
			Name: department_fields.DepartmentName,
		})

		// Ensure theres no inserting department error
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println("New department", department_fields.DepartmentName, "has been created")

		/* err = models.DepartmentHead.Insert(models.DepartmentHead{
			DepartmentId: int(department_id),
			ManagerId:    &department_fields.ManagerId,
		})

		// Ensure theres no inserting department head error
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Server error",
			})

			return
		} */

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Created",
		})
	}
}
