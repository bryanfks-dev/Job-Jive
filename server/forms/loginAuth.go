package forms

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"configs"
	"db"
	"models"
)

var (
	postMu sync.Mutex
)

type UserCred struct {
	Email string
	Password string
	Remember bool
}

type AdminCred struct {
	Username string
	Password string
}

func VerifyPassword(hashed_pwd string, cred_pwd string) bool {
	// Database failed to find user
	if len(hashed_pwd) == 0 {
		return false
	}

	// Comparing hashed password from database and login credential
	err := bcrypt.CompareHashAndPassword([]byte(hashed_pwd), []byte(cred_pwd))

	// Not nil value in err possibly cause of hash and password 
	// values are not match, otherwise, means hash and password 
	// values are match
	return err == nil
}

func userVerified[T UserCred|AdminCred](cred T) bool {
	if db.ConnectionEstablished() {
		switch any(cred).(type) {
			case UserCred: {
				hashed_pwd := models.User.GetHashedPassword(models.User{}, any(cred).(UserCred).Email)

				return VerifyPassword(hashed_pwd, any(cred).(UserCred).Password)
			}
			case AdminCred: {
				hashed_pwd := models.Admin.GetHashedPassword(models.Admin{}, any(cred).(AdminCred).Username)

				return VerifyPassword(hashed_pwd, any(cred).(AdminCred).Password)
			}
		}
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

		if userVerified(login_cred) {
			http.Redirect(w, r, "http://localhost:8000/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "http://localhost:8000" + "/login?msg=Invalid+Login+Credentials", 
			http.StatusSeeOther)
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

		fmt.Println(login_cred)
	}
}