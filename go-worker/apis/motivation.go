package apis

import (
	"net/http"
)

func GetUserMotivation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()
	}
}
