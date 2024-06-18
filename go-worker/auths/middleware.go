package auths

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	"models"

	"github.com/golang-jwt/jwt/v5"
)

type context_key string

const (
	TOKEN_KEY context_key = "token"
)

func AuthenticationMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get authorization header
		auth_header := r.Header.Get("Authorization")

		w.Header().Set("Content-Type", "application/json")

		login_routes := []string{"/auth/verify-token", "/auth/user/login", "/auth/admin/login"}
		is_login_route := slices.Contains(login_routes, r.URL.Path)

		is_valid_token := true

		// Ensure user has authorization in header
		if auth_header == "" {
			if !is_login_route {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "authorization header missing",
				})

				return
			}

			is_valid_token = false
		}

		// Ensure the header starts with "Bearer " and extract the token
		if !strings.HasPrefix(auth_header, "Bearer ") {
			if !is_login_route {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "invalid authorization header format",
				})

				return
			}

			is_valid_token = false
		}

		// Extract the token from the header
		token := strings.TrimPrefix(auth_header, "Bearer ")

		if token == "" {
			if !is_login_route {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "token missing",
				})

				return
			}

			is_valid_token = false
		}

		// Claims user id from token
		extract_token, err := models.ClaimsToken(token)

		// Ensure no error when claimming token
		if err != nil {
			if !is_login_route {
				if err == jwt.ErrTokenExpired {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(map[string]any{
						"error": "token expired",
					})

					return
				}

				log.Panic("Unable to claim token: ", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "server error",
				})

				return
			}

			is_valid_token = false
		}

		if is_login_route {
			if is_valid_token {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "forbidden",
				})

				return
			}

			// Should be either login or verify-token http handler function
			next.ServeHTTP(w, r)

			return
		}

		// Add token to current api context
		ctx :=
			context.WithValue(r.Context(), TOKEN_KEY, extract_token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt_claims, ok :=
			r.Context().Value(TOKEN_KEY).(jwt.MapClaims)

		w.Header().Set("Content-Type", "application/json")

		// Ensure token is available
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "No token found in context",
			})

			return
		}

		// Check user role
		if jwt_claims["role"].(string) != "user" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "forbidden",
			})

			return
		}

		// Check if user is exist in database
		user, err :=
			models.User{}.GetUsingId(int(jwt_claims["id"].(float64)))

		// Ensure no error when getting user data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "Invalid user id",
				})

				return
			}

			// Other errors
			log.Panic("Error get user: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Check for specific user role, either manager or employee
		department_head, err :=
			models.DepartmentHead{}.GetUsingDepartmentId(user.DepartmentId)

		// Ensure no error when getting department head data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "Invalid department id",
				})

				return
			}

			// Other errors
			log.Panic("Error get department: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		// Specify user position
		as := "employee"

		if user.Id == *department_head.ManagerId {
			as = "manager"
		}

		// Edit token value
		jwt_claims["as"] = as

		ctx :=
			context.WithValue(r.Context(), TOKEN_KEY, jwt_claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt_claims, ok := r.Context().Value(context_key("token")).(jwt.MapClaims)

		w.Header().Set("Content-Type", "application/json")

		// Ensure token is available
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "No token found in context",
			})

			return
		}

		// Check user role
		if jwt_claims["role"].(string) != "admin" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "forbidden",
			})

			return
		}

		// Check if admin is exist in database
		_, err :=
			models.Admin{}.GetUsingId(int(jwt_claims["id"].(float64)))

		// Ensure no error when getting admin data
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"error": "Invalid admin id",
				})

				return
			}

			// Other errors
			log.Panic("Error get admin: ", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "server error",
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
