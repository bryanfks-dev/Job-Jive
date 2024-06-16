package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"forms"
	"models"
)

func GetConfigsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		config_json, err :=
			models.ConfigJson{}.LoadConfig()

		// Ensure no error loading config file
		if err != nil {
			log.Panic("Error load config json: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": config_json,
		})
	}
}

func SaveConfigsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == http.MethodPut {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var config_form forms.ConfigForm

		err := req_json.Decode(&config_form)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		config_form.Sanitize()

		valid, err := config_form.Validate()

		// Ensure form is valid
		if !valid {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			})

			return
		}

		err = models.ConfigJson(config_form).WriteFile()

		// Ensure no error writting config file
		if err != nil {
			log.Panic("Error write config file: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": models.ConfigJson(config_form),
		})
	}
}
