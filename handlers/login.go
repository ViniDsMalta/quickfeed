package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"quickfeed/auth"
	"quickfeed/database"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var storedHash string

	err = database.DB.QueryRow(
		"SELECT password_hash FROM users WHERE email=$1",
		req.Email,
	).Scan(&storedHash)

	if err == sql.ErrNoRows {
		http.Error(w, "this user does not exist", http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, "intern error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedHash),
		[]byte(req.Password),
	)

	if err != nil {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(req.Email)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}