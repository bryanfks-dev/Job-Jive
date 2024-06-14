package auths

import (
	"database/sql"
	"log"
	"net/http"

	"models"

	"github.com/golang-jwt/jwt/v5"
)

func UserMiddleware(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
	// Validate token
	token_valid, res := AuthorizedToken(w, r)

	if !token_valid {
		return false, res
	}

	jwt_claims := res["token"].(jwt.MapClaims)

	// Check user role
	if jwt_claims["role"].(string) != "user" {
		w.WriteHeader(http.StatusForbidden)

		return false, map[string]interface{}{
			"error": "Forbidden",
		}
	}

	// Check if user is exist in database
	user, err :=
		models.User{}.GetUsingId(int(jwt_claims["id"].(float64)))

	// Ensure no error when getting user data
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)

			return false, map[string]interface{}{
				"error": "Invalid user",
			}
		}

		log.Panic("Error occured:", err)

		w.WriteHeader(http.StatusInternalServerError)

		return false, map[string]interface{}{
			"error": "Server error",
		}
	}

	// Check for specific user role, either manager or employee
	department_head, err :=
		models.DepartmentHead{}.GetUsingDepartmentId(user.DepartmentId)

	// Ensure no error when getting department head data
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)

			return false, map[string]interface{}{
				"error": "Invalid department",
			}
		}

		w.WriteHeader(http.StatusInternalServerError)

			return false, map[string]interface{}{
				"error": "Server error",
			}
	}

	// Specify user title
	var title string

	if *department_head.ManagerId == user.Id {
		title = "manager"
	} else {
		title = "employee"
	}

	return true, map[string]interface{}{
		"as": title,
		"jwt_claims": jwt_claims,
	}
}

func AdminMiddleware(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
	// Validate token
	token_valid, res := AuthorizedToken(w, r)

	if !token_valid {
		return false, res
	}

	jwt_claims := res["token"].(jwt.MapClaims)

	// Check user role
	if jwt_claims["role"].(string) != "admin" {
		w.WriteHeader(http.StatusForbidden)

		return false, map[string]interface{}{
			"message": "Forbidden",
		}
	}

	// Check if admin is exist in database
	_, err := 
		models.Admin{}.GetUsingId(int(jwt_claims["id"].(float64)))

	// Ensure no error when getting user data
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return false, map[string]interface{}{
			"error": "Invalid user",
		}
	}

	return true, map[string]interface{}{
		"jwt_claims": jwt_claims,
	}
}
