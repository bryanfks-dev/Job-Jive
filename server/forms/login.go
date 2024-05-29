package forms

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"db"
	"models"
)

type (
	UserCred  models.UserCred
	AdminCred models.AdminCred
)

var (
	postMu                  sync.Mutex
	errUserNotFound         = errors.New("user not found")
	errDBConnNotEstablished = errors.New("cannot established connection with database")
)

func verifyPassword(hashed_pwd string, cred_pwd string) (bool, error) {
	// Database failed to find user
	if len(hashed_pwd) == 0 {
		return false, errUserNotFound
	}

	// Comparing hashed password from database and login credential
	err := bcrypt.CompareHashAndPassword([]byte(hashed_pwd), []byte(cred_pwd))

	// Not nil value in err possibly cause of hash and password
	// values are not match, otherwise, means hash and password
	// values are match
	return err == nil, err
}

func userVerified[T UserCred | AdminCred](cred T) (bool, error) {
	if db.ConnectionEstablished() {
		switch any(cred).(type) {
		case UserCred:
			{
				hashed_pwd := models.User.GetHashedPassword(models.User{}, any(cred).(UserCred).Email)

				return verifyPassword(hashed_pwd, any(cred).(UserCred).Password)
			}
		case AdminCred:
			{
				hashed_pwd := models.Admin.GetHashedPassword(models.Admin{}, any(cred).(AdminCred).Username)

				return verifyPassword(hashed_pwd, any(cred).(AdminCred).Password)
			}
		}
	}

	return false, errDBConnNotEstablished
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var login_cred UserCred

		err := req_json.Decode(&login_cred)

		if err != nil {
			panic(err.Error())
		}

		verified, err := userVerified(login_cred)

		var response map[string]interface{}

		// Errors other than user not found and invalid password
		if err != nil && (err != errUserNotFound && err != bcrypt.ErrMismatchedHashAndPassword) {
			response = map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Server error",
			}
		} else {
			// User is veified
			if verified {
				token, err := models.CreateToken(login_cred.Password)

				// Failed generate token
				if err != nil {
					response = map[string]interface{}{
						"status":  http.StatusInternalServerError,
						"message": "Server error",
					}
				} else {
					log.Println(login_cred.Email, " logged in")

					response = map[string]interface{}{
						"status":  http.StatusOK,
						"message": "Login success",
						"token":   token,
					}
				}
			} else { // User not found or invalid password
				response = map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Invalid credentials",
				}
			}
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}
}
