package apis

import (
	"encoding/json"
	"net/http"
	"sync"

	"auths"
)

var (
	postMu sync.Mutex
)

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		token_is_valid, res := auths.AuthorizedToken(r)

		w.Header().Set("Content-Type", "application/json")

		if !token_is_valid {
			json.NewEncoder(w).Encode(res)

			return
		}

		// Check role logic.. etc
	}
}
