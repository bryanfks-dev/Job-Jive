package apis

import (
	"log"
	"net/http"
	"sync"

	"models"
)

type (
	User models.User
)

var (
	postMu sync.Mutex
)

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		log.Println(r.Header.Get("Authorization"))
	}
}
