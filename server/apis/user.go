package apis

import (
	"database/sql"
	"encoding/json"
	"models"
	"net/http"
	"sync"

	"auths"
)

var (
	postMu sync.Mutex
)

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		valid_user, res := auths.UserMiddleware(r)

		if !valid_user {
			json.NewEncoder(w).Encode(res)

			return
		}

		user, err :=
			models.User.GetUsingId(models.User{}, res["id"].(int))

		// Ensure no error when getting user information
		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Invalid token value",
				})
	
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"data":   user,
		})
	}
}
