package middleware

import (
	"context"
	"net/http"
	"strings"

	"quickfeed/auth"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		_, claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		email := claims["email"].(string)

		// salvar no contexto
		ctx := context.WithValue(r.Context(), "userEmail", email)

		next(w, r.WithContext(ctx))
	}
}