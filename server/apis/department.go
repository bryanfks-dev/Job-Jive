package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"auths"
	"models"

	"github.com/golang-jwt/jwt/v5"
)

type DepartmentFields struct {
	DepartmentName string `json:"department-name"`
	ManagerId      int    `json:"manager-id"`
}

type ResponseData struct {
	DepartmentName string `json:"department_name"`
	ManagerName    string `json:"manager_name"`
}

func GetDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
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

		jwt_claims := res["token"].(jwt.MapClaims)

		// Check user role
		if jwt_claims["role"].(string) != "admin" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusForbidden,
				"message": "Forbidden",
			})

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

		var response_data []ResponseData

		for _, data := range departments {
			department_head, err :=
				models.DepartmentHead.GetUsingDepartmentId(
					models.DepartmentHead{}, data.Id)

			// Ensure no error fetching department
			if err != nil {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Server error",
				})

				return
			}

			var user models.User

			// Ensure department_head not empty
			if department_head != (models.DepartmentHead{}) {
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

			response_data = append(response_data, ResponseData{
				DepartmentName: data.Name,
				ManagerName:    user.FullName,
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

		// Validate token
		token_valid, res := auths.AuthorizedToken(r)

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		if !token_valid {
			json.NewEncoder(w).Encode(res)

			return
		}

		jwt_claims := res["token"].(jwt.MapClaims)

		// Check user role
		if jwt_claims["role"].(string) != "admin" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusForbidden,
				"message": "Forbidden",
			})

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

		// Ensure no error when inserting department
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
