package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"quickfeed/database"
	"quickfeed/handlers"
	"quickfeed/middleware"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system variables")
	}

	connStr := "host=localhost port=5433 user=postgres password=123456 dbname=quickfeed sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database.DB = db

	log.Println("db succesfull connected")

	
	r := chi.NewRouter()

	
	r.Get("/health", handlers.HealthHandler)
	r.Get("/{slug}", handlers.GetCompanyBySlugHandler)

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	
	r.With(middleware.JWTAuth).Get(
		"/profile",
		handlers.ProfileHandler,
	)

	r.With(middleware.JWTAuth).Route("/companies", func(r chi.Router) {

		r.Post("/", handlers.CreateCompanyHandler)

		r.Get("/", handlers.ListCompaniesHandler)
	})

	log.Println("running in http://localhost:8080")

	log.Fatal(
		http.ListenAndServe(
			":8080",
			enableCORS(r),
		),
	)
}