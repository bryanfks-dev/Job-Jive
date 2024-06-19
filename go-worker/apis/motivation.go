package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"ai"
	"auths"
	"responses"

	"github.com/golang-jwt/jwt/v5"
	"github.com/googleapis/gax-go/v2/apierror"
)

func GetUserMotivation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		postMu.Lock()
		defer postMu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		token, ok :=
			r.Context().Value(auths.TOKEN_KEY).(jwt.MapClaims)

		// Ensure token extract from request
		if !ok {
			log.Panic("ERROR get token from context")

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "token not found",
			})

			return
		}

		user_id := int(token["id"].(float64))

		// Do absence analysis codes here

		promt := "give me a single motivation for my work, i have a bad day right now :("

		motivation, err := ai.GeminiGenerate(promt)

		// Ensure no error promting to gemini
		if err != nil {
			if api_err, ok := err.(*apierror.APIError); ok &&
				api_err.HTTPCode() == http.StatusTooManyRequests {
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "too many request",
				})

				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}
		motivation = strings.ReplaceAll(motivation, "*", "")

		motivation =
			regexp.MustCompile(`[^a-zA-Z0-9 ',./&-?]+`).ReplaceAllString(motivation, "")

		response_data := responses.MotivationResponse{
			UserId:     user_id,
			Motivation: `"` + strings.TrimSpace(motivation) + `"`,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"data": response_data,
		})
	}
}
