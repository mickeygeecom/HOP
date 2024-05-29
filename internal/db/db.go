package db

import (
	"database/sql"
	"log"
	"os"

	// Import the MySQL driver here, to isolate it from the rest of the application
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB ... global variable to hold the database connection
var DB *sql.DB

// InitDB ... initializes a connection to the database using the credentials
func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get database credentials from environment variables
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create a Data Source Name (DSN) for the database connection
	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	// Open a connection to the database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Ping the database to check if the connection is successful
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	// Update quiz join codes after successful database connection
	UpdateQuizJoinCodes()
}
