package auths

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"configs"
	"models"

	"golang.org/x/crypto/bcrypt"
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
	postMu sync.Mutex
	_token_expire_time = map[bool]int{
		true: 30,
		false: 6,
	}
)

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		w.WriteHeader(http.StatusOK)
	}
}

func verifyPassword(hashed_pwd string, cred_pwd string) error {
	// Comparing hashed password from database and login credential
	err :=
		bcrypt.CompareHashAndPassword([]byte(hashed_pwd), []byte(cred_pwd))

	// Not nil value in err possibly cause of hash and password
	// values are not match, otherwise, means hash and password
	// values are match
	return err
}

func verifyUser(cred UserCred) (models.User, error) {
	user, err :=
		models.User{}.GetUsingEmail(cred.Email)

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
		models.Admin{}.GetUsingUsername(cred.Username)

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

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var login_cred UserCred

		err := req_json.Decode(&login_cred)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		user, err := verifyUser(login_cred)

		if err != nil {
			// Incorrect password or user not found
			if err == bcrypt.ErrMismatchedHashAndPassword || err == sql.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid credential",
				})

				return
			}

			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Generate jwt token
		token, err := 
			models.CreateToken(user.Id, "user", _token_expire_time[login_cred.Remember])

		// Ensure jwt token generated
		if err != nil {
			log.Panic("Error generate token: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "Could not generate token",
			})

			return
		}

		// Update user first login if first login date haven't made
		if user.FirstLogin == nil {
			// BUG: user first login cannot be updated
			var zone configs.Timezone
			err = zone.Load()

			// Ensure no error get timezone from env
			if err != nil {
				log.Panic("Error get timezone from env: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			tz, err := time.LoadLocation(zone.Zone)

			// Ensure no error getting timezone
			if err != nil {
				log.Panic("Error get timezone", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			// Get current time within current timezone
			curr_date := time.Now().In(tz).String()

			err = user.UpdateFistLogin(curr_date)

			// Ensure no error updating user
			if err != nil {
				log.Panic("Error update first_login user: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}
		}

		log.Println("user", user.FullName, "logged in")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"token": token,
		})
	}
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		// Set HTTP header
		w.Header().Set("Content-Type", "application/json")

		// Decode json to struct
		req_json := json.NewDecoder(r.Body)

		var login_cred AdminCred

		err := req_json.Decode(&login_cred)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "bad request",
			})

			return
		}

		admin, err := verifyAdmin(login_cred)

		if err != nil {
			// Incorrect password or user not found
			if err == bcrypt.ErrMismatchedHashAndPassword || err == sql.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid credential",
				})

				return
			}

			log.Panic("Error get admin", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Generate jwt token
		token, err := models.CreateToken(admin.Id, "admin", 1)

		// Ensure jwt token generated
		if err != nil {
			log.Panic("Error generate token", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		log.Println("Admin", login_cred.Username, "logged in")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"token": token,
		})
	}
}
