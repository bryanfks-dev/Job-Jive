package apis

import (
	"encoding/json"
	"net/http"

	"auths"
	"models"
)

func GetConfigsHandler(w http.ResponseWriter, r *http.Request) {
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

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_admin, res := auths.AdminMiddleware(r)

		if !valid_admin {
			json.NewEncoder(w).Encode(res)

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
