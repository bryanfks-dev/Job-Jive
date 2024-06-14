package auths

import (
	"net/http"
	"strings"

	"models"

	"github.com/golang-jwt/jwt/v5"
)

func AuthorizedToken(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
	// Get authorization header
	auth_header := r.Header.Get("Authorization")

	// Ensure user has authorization in header
	if auth_header == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return false, map[string]interface{}{
			"error": "Authorization header missing",
		}
	}

	// Ensure the header starts with "Bearer " and extract the token
	if !strings.HasPrefix(auth_header, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)

		return false, map[string]interface{}{
			"error": "Invalid authorization header format",
		}
	}

	// Extract the token from the header
	token := strings.TrimPrefix(auth_header, "Bearer ")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return false, map[string]interface{}{
			"error": "Token missing",
		}
	}

	// Claims user id from token
	token_map, err := models.ClaimsToken(token)

	// Ensure no error when claimming token
	if err != nil {
		if err == jwt.ErrTokenExpired {
			w.WriteHeader(http.StatusUnauthorized)

			return false, map[string]interface{}{
				"error": "Token expired",
			}
		}

		w.WriteHeader(http.StatusInternalServerError)

		return false, map[string]interface{}{
			"error": "Unable to claim token",
		}
	}

	return true, map[string]interface{}{
		"token": token_map,
	}
}
