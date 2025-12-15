package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	// Create users table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id   SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		dob  DATE NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	fmt.Println("✅ Migration successful! Users table created.")
}
