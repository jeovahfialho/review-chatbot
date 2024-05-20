package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for database connection
	"github.com/jmoiron/sqlx"          // SQLX library for SQL database access
)

// InitDB initializes the database connection
func InitDB() (*sqlx.DB, error) {
	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Create the Data Source Name (DSN) for connecting to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	// Connect to the database using SQLX
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err // Return an error if the connection fails
	}
	return db, nil // Return the database connection
}
