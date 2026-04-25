package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"quickfeed/database"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, "Erro ao gerar hash", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		req.Email,
		string(hash),
	)

	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("sucessfull register"))
}