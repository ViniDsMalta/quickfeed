package middleware

import (
	"context"
	"net/http"
	"strings"

	"quickfeed/auth"
)

func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		_, claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		email := claims["email"].(string)

	
		ctx := context.WithValue(
			r.Context(),
			"userEmail",
			email,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}