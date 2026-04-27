package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"quickfeed/database"
	"quickfeed/handlers"
	"quickfeed/middleware"
)

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

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", middleware.JWTAuth(handlers.ProfileHandler))
	log.Println("running in http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}