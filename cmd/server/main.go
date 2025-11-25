package main

import (
	"log"

	"github.com/rlevidev/e-learning-go/internal/config/database"
)

func main() {
	// Initialize database connection and run migrations
	_, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Database initialized successfully!")
}
