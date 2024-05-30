package apis

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"models"
)

type (
	User models.User
)

var (
	postMu sync.Mutex
)

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		auth_header := r.Header.Get("Authorization")

		// Ensure user has authorization in header
		if auth_header == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Authorization header missing",
			})

			return
		}

		// Ensure the header starts with "Bearer " and extract the token
		if !strings.HasPrefix(auth_header, "Bearer ") {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Invalid authorization header format",
			})

			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(auth_header, "Bearer ")

		if token == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Token missing",
			})

			return
		}

		tokenVerified, _ := models.VerifyToken(token)

		if !tokenVerified {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusUnauthorized,
				"message": "Invalid token",
			})

			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			
		})
	}
}
