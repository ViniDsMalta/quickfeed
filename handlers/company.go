package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"quickfeed/database"
	"strings"

	"github.com/go-chi/chi/v5"
)

type CreateCompanyRequest struct {
	Name string `json:"name"`
}
type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
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
func ListCompaniesHandler(w http.ResponseWriter, r *http.Request) {

	email := r.Context().Value("userEmail").(string)

	var ownerID int

	err := database.DB.QueryRow(
		"SELECT id FROM users WHERE email=$1",
		email,
	).Scan(&ownerID)

	if err != nil {
		http.Error(w, "user not found", http.StatusInternalServerError)
		return
	}

	rows, err := database.DB.Query(
		"SELECT id, name, slug FROM companies WHERE owner_id=$1",
		ownerID,
	)

	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var companies []Company

	for rows.Next() {

		var company Company

		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Slug,
		)

		if err != nil {
			http.Error(w, "scan error", http.StatusInternalServerError)
			return
		}

		companies = append(companies, company)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(companies)
}
func GetCompanyBySlugHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	slug := chi.URLParam(r, "slug")

	var company Company

	err := database.DB.QueryRow(
		"SELECT id, name, slug FROM companies WHERE slug=$1",
		slug,
	).Scan(
		&company.ID,
		&company.Name,
		&company.Slug,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			http.Error(w, "company not found", http.StatusNotFound)
			return
		}

		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(company)
}