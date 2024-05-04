package forms

import (
	"net/http"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"models"
	"db"
	"configs"
)

type LoginCred struct {
	Email string
	Password string
	Remember bool
}

var (
	postMu sync.Mutex
)

func userVerified(cred LoginCred) bool {
	if db.ConnectionEstablished() {
		user_data := models.User.Search(models.User{}, cred.Email)

		// Means database failed to find user data
		if user_data.Id == 0 {
			return false
		}

		// Comparing hashed password from database and login credential
		err := bcrypt.CompareHashAndPassword(
			[]byte(user_data.Password), []byte(cred.Password))

		// Not nil value in err possibly cause of hash and password 
		// values are not match, otherwise, means hash and password 
		// values are match
		return err == nil
	}

	return false
}

func createSession(w http.ResponseWriter, r *http.Request, value int) { 
	session, err := configs.UserSession.Store.Get(r, "user-sessions")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	session.Values["key"] = value

	err = session.Save(r, w)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMu.Lock()
		defer postMu.Unlock()

		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Interal Server Error", http.StatusInternalServerError)
			return
		}

		// Get login credential
		login_cred := LoginCred{
			Email: strings.TrimSpace(r.Form.Get("email")),
			Password: r.Form.Get("password"),
			Remember: r.Form.Get("remember") != "",
		}

		if userVerified(login_cred) {
			http.Redirect(w, r, "http://localhost:8000/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "http://localhost:8000" + "/login?msg=Invalid+Login+Credentials", 
			http.StatusSeeOther)
	}
}