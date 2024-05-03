package forms

import (
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"models"
)

var (
	postMu sync.Mutex
)

func LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postMu.Lock()
		defer postMu.Unlock()

		r.ParseForm()

		// Get login credential
		login_cred := map[string] string {
			"email": r.FormValue("email"),
			"password": r.FormValue("password"),
			"remember": r.FormValue("remember"),
		}

		user_data := models.User.Search(models.User{}, login_cred["email"])

		// User data is available in users table
		if user_data.Id != 0 {
			err := bcrypt.
				CompareHashAndPassword([]byte(user_data.Password), []byte(login_cred["password"]))

			// Not nil value in err possibly cause of hash and password 
			// values are not match, otherwise, if err is nil, means hash 
			// and password values are match
			if err == nil {
				w.WriteHeader(http.StatusOK) // 200 OK

				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
	}
}