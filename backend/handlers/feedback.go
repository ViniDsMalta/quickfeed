package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"quickfeed/database"
)

type FeedbackRequest struct {
	Message string `json:"message"`
}
type Feedback struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

func CreateFeedbackHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	slug := chi.URLParam(r, "slug")

	var req FeedbackRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}


	var companyID int

	err = database.DB.QueryRow(
		"SELECT id FROM companies WHERE slug=$1",
		slug,
	).Scan(&companyID)

	if err != nil {

		if err == sql.ErrNoRows {
			http.Error(w, "company not found", http.StatusNotFound)
			return
		}

		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}


	_, err = database.DB.Exec(
		"INSERT INTO feedbacks (company_id, message) VALUES ($1, $2)",
		companyID,
		req.Message,
	)

	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("feedback created"))
}

func ListFeedbacksHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	slug := chi.URLParam(r, "slug")

	email := r.Context().Value("userEmail").(string)


	var companyID int

	err := database.DB.QueryRow(
		`SELECT companies.id
		 FROM companies
		 JOIN users
		 ON companies.owner_id = users.id
		 WHERE users.email=$1
		 AND companies.slug=$2`,
		email,
		slug,
	).Scan(&companyID)

	if err != nil {

		if err == sql.ErrNoRows {
			http.Error(w, "company not found or unauthorized", http.StatusForbidden)
			return
		}

		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	
	rows, err := database.DB.Query(
		`SELECT id, message, created_at
		 FROM feedbacks
		 WHERE company_id=$1
		 ORDER BY created_at DESC`,
		companyID,
	)

	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var feedbacks []Feedback

	for rows.Next() {

		var feedback Feedback

		err := rows.Scan(
			&feedback.ID,
			&feedback.Message,
			&feedback.CreatedAt,
		)

		if err != nil {
			http.Error(w, "scan error", http.StatusInternalServerError)
			return
		}

		feedbacks = append(feedbacks, feedback)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(feedbacks)
}