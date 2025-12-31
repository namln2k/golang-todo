package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Get the DB_URL from environment variables set in docker-compose.yml
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// 2. Open a connection to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// 3. Ping the database to ensure connection is active, retry until successful
	for {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Could not connect to the database, retrying: %v", err)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Successfully connected to PostgreSQL via Docker Compose!")

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/hello-world", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	if err := r.Run(":8888"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
