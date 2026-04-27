package middleware

import (
	"net/http"
	"strings"

	"quickfeed/auth"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "whithout a token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		_, err := auth.ValidateToken(token)

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}