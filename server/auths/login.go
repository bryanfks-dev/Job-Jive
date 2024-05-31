package auths

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"models"
)

type UserCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type AdminCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	postMu          sync.Mutex
	errUserNotFound = errors.New("user not found")
)

func verifyPassword(hashed_pwd string, cred_pwd string) error {
	// Database failed to find user
	if hashed_pwd == "" {
		return errUserNotFound
	}

	// Comparing hashed password from database and login credential
	err := bcrypt.CompareHashAndPassword([]byte(hashed_pwd), []byte(cred_pwd))

	// Not nil value in err possibly cause of hash and password
	// values are not match, otherwise, means hash and password
	// values are match
	return err
}

func verifyUser(cred UserCred) (models.User, error) {
	user, err :=
		models.User.GetUsingEmail(models.User{}, cred.Email)

	// Ensure user is exist
	if err != nil {
		return models.User{}, err
	}

	// Verify user password
	err = verifyPassword(user.Password, cred.Password)

	// Ensure password matched
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func verifyAdmin(cred AdminCred) (models.Admin, error) {
	admin, err :=
		models.Admin.GetUsingUsername(models.Admin{}, cred.Username)

	// Ensure admin is exist
	if err != nil {
		return models.Admin{}, err
	}

	// Verify user password
	err = verifyPassword(admin.Password, cred.Password)

	// Ensure password matched
	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var login_cred UserCred

		err := req_json.Decode(&login_cred)

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
			})

			return
		}

		user, err := verifyUser(login_cred)

		if err != nil {
			// Incorrect password or user not found
			if err == bcrypt.ErrMismatchedHashAndPassword || err == errUserNotFound {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Invalid credential",
				})

				return
			}

			// Other errors
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Generate jwt token
		token, err := models.CreateToken(user.Id, "user")

		// Ensure jwt token generated
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		log.Println(login_cred.Email, "logged in")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Login success",
			"token":   token,
		})
	}
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var login_cred AdminCred

		err := req_json.Decode(&login_cred)

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Invalid credential",
			})

			return
		}

		admin, err := verifyAdmin(login_cred)

		if err != nil {
			// Incorrect password or user not found
			if err == bcrypt.ErrMismatchedHashAndPassword || err == errUserNotFound {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Invalid credential",
				})

				return
			}

			// Other errors
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			})

			return
		}

		// Generate jwt token
		token, err := models.CreateToken(admin.Id, "admin")

		// Ensure jwt token generated
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Could not generate token",
			})

			return
		}

		log.Println("Admin", login_cred.Username, "logged in")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Login success",
			"token":   token,
		})
	}
}
