package forms

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"db"
	"models"
)

type (
	UserCred models.UserCred
	AdminCred models.AdminCred
)

var (
	postMu sync.Mutex
	errUserNotFound = errors.New("user not found")
	errDBConnNotEstablished = errors.New("cannot established connection with database")
)

func sendJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

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

func userVerified[T UserCred|AdminCred](cred T) (bool, error) {
	if db.ConnectionEstablished() {
		switch any(cred).(type) {
			case UserCred: {
				hashed_pwd := models.User.GetHashedPassword(models.User{}, any(cred).(UserCred).Email)

				return verifyPassword(hashed_pwd, any(cred).(UserCred).Password)
			}
			case AdminCred: {
				hashed_pwd := models.Admin.GetHashedPassword(models.Admin{}, any(cred).(AdminCred).Username)

				return verifyPassword(hashed_pwd, any(cred).(AdminCred).Password)
			}
		}
	}

	return false, errDBConnNotEstablished
}

func LoginUserAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Interal Server Error", http.StatusInternalServerError)
			return
		}

		// Get login credential
		login_cred := UserCred{
			Email: strings.TrimSpace(r.Form.Get("email")),
			Password: r.Form.Get("password"),
			Remember: r.Form.Get("remember") != "",
		}

		verified, err := userVerified(login_cred)

		if verified {
			sendJson(w, http.StatusOK, map[string]string{
				"message": "Login Success",
			})
		} else if !verified && (err == errUserNotFound || err == bcrypt.ErrMismatchedHashAndPassword) {
			sendJson(w, http.StatusUnauthorized, map[string]string{
				"message": "Invalid Login Credentials",
			})
		}
	}
}

func LoginAdminAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Interal Server Error", http.StatusInternalServerError)
			return
		}

		// Get login credential
		login_cred := AdminCred{
			Username: strings.TrimSpace(r.Form.Get("username")),
			Password: r.Form.Get("password"),
		}

		verified, err := userVerified(login_cred)

		if verified {
			sendJson(w, http.StatusOK, map[string]string{
				"message": "Login Success",
			})
		} else if !verified && (err == errUserNotFound || err == bcrypt.ErrMismatchedHashAndPassword) {
			sendJson(w, http.StatusUnauthorized, map[string]string{
				"message": "Invalid Login Credentials",
			})
		}
	}
}