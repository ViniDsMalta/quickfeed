package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"quickfeed/database"
)

type CreateCompanyRequest struct {
	Name string `json:"name"`
}

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	
	slug := strings.ToLower(strings.ReplaceAll(req.Name, " ", "-"))

	
	email := r.Context().Value("userEmail").(string)

	var ownerID int

	err = database.DB.QueryRow(
		"SELECT id FROM users WHERE email=$1",
		email,
	).Scan(&ownerID)

if err != nil {
	http.Error(w, "user not found", http.StatusInternalServerError)
	return
}

	_, err = database.DB.Exec(
		"INSERT INTO companies (name, slug, owner_id) VALUES ($1, $2, $3)",
		req.Name,
		slug,
		ownerID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("company created"))
}