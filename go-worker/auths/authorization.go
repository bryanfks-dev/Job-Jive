package auths

import (
	"net/http"
	"strings"

	"models"

	"github.com/golang-jwt/jwt/v5"
)

func AuthorizedToken(r *http.Request) (bool, map[string]interface{}) {
	// Get authorization header
	auth_header := r.Header.Get("Authorization")

	// Ensure user has authorization in header
	if auth_header == "" {
		return false, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Authorization header missing",
		}
	}

	// Ensure the header starts with "Bearer " and extract the token
	if !strings.HasPrefix(auth_header, "Bearer ") {
		return false, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Invalid authorization header format",
		}
	}

	// Extract the token from the header
	token := strings.TrimPrefix(auth_header, "Bearer ")

	if token == "" {
		return false, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Token missing",
		}
	}

	// Claims user id from token
	token_map, err := models.ClaimsToken(token)

	// Ensure no error when claimming token
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return false, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Token expired",
			}
		}

		return false, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Unable to claim token",
		}
	}

	return true, map[string]interface{}{
		"status":  http.StatusOK,
		"token":   token_map,
	}
}
