package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"models"

	"github.com/golang-jwt/jwt/v5"
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

		// Get authorization header
		auth_header := r.Header.Get("Authorization")

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

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

		// Claims user id from token
		token_map, err := models.ClaimsToken(token)

		// Ensure no error when claimming token
		if err != nil {
			if err == jwt.ErrTokenExpired {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": http.StatusUnauthorized,
					"message": "Token expired",
				})

				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusInternalServerError,
				"message": "Unable to claim token",
			})

			return
		}

		log.Println(token_map)
	}
}
