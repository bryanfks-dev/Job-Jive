package apis

import (
	"encoding/json"
	"net/http"

	"auths"
	"models"

	"github.com/golang-jwt/jwt/v5"
)

func GetConfigsHandler(w http.ResponseWriter, r *http.Request) {
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

		config_json, err :=
			models.ConfigJson.LoadConfig(models.ConfigJson{})

		// Ensure no error loading config file
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   config_json,
		})
	}
}

func SaveConfigsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
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

		var config models.ConfigJson

		err := req_json.Decode(&config)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		// Validating values
		if config.AbsenceQuota <= 0 || config.DailyWorkHours <= 0 || 
			config.WeekyWorkHours <= 0 || config.CheckInTime >= config.CheckOutTime || 
			config.WeekyWorkHours > config.DailyWorkHours {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		err = models.ConfigJson.WriteFile(config)

		// Ensure no error writting config file
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Updated",
		})
	}
}
