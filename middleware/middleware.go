package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/amosehiguese/restaurant-api/auth"
)

type resp map[string]any

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateJWT(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"msg": "Authentication required",
				"code": http.StatusUnauthorized,
			})
			return 
		}

		err = auth.ValidateAdminRoleJWT(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"msg": "Authorization failed. Admin-only route",
				"code": http.StatusUnauthorized,
			})
			return 
		}

		next.ServeHTTP(w, r)
	})
}

func JWTAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateJWT(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"msg": "Authentication required",
				"code": http.StatusUnauthorized,
			})
			return 
		}

		err = auth.ValidateUserRoleJWT(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"msg": "Authorization failed. Only registered users are allowed to perform this action",
				"code": http.StatusUnauthorized,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}