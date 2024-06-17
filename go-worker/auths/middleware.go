package auths

import (
	"models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func UserMiddleware(r *http.Request) (bool, map[string]interface{}) {
	// Validate token
	token_valid, res := AuthorizedToken(r)

	if !token_valid {
		return false, res
	}

	jwt_claims := res["token"].(jwt.MapClaims)

	// Check user role
	if jwt_claims["role"].(string) != "user" {
		return false, map[string]interface{}{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
		}
	}

	// Check if user is exist in database
	_, err := models.User.GetUsingId(models.User{}, int(jwt_claims["id"].(float64)))

	// Ensure no error when getting user data
	if err != nil {
		return false, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Invalid user",
		}
	}

	return true, map[string]interface{}{
		"status":     http.StatusOK,
		"jwt_claims": jwt_claims,
	}
}

func AdminMiddleware(r *http.Request) (bool, map[string]interface{}) {
	// Validate token
	token_valid, res := AuthorizedToken(r)

	if !token_valid {
		return false, res
	}

	jwt_claims := res["token"].(jwt.MapClaims)

	// Check user role
	if jwt_claims["role"].(string) != "admin" {
		return false, map[string]interface{}{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
		}
	}

	// Check if admin is exist in database
	_, err := models.Admin.GetUsingId(models.Admin{}, int(jwt_claims["id"].(float64)))

	// Ensure no error when getting user data
	if err != nil {
		return false, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Invalid user",
		}
	}

	return true, map[string]interface{}{
		"status":     http.StatusOK,
		"jwt_claims": jwt_claims,
	}
}
