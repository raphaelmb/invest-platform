package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// TODO: refactor main
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in environment")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	srv := &http.Server{
		Handler: apiCfg.routes(),
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server started on port %v", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
