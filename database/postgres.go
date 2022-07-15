package database

import (
	"database/sql"
	"fmt"
	"stream-it-consulting/innovation-team/example-rest-api/config"
	"time"

	_ "github.com/lib/pq"
)

// Initializes the database
func Initialize() *sql.DB {
	// Create database connection
	db, err := sql.Open("postgres", config.AppConfig.DatabaseDSN)

	// Check if connection is successful
	if err != nil {
		fmt.Println("Can't connect to database: ", err)
	}

	// Database connection settings
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(200)

	// Return the database connection
	return db
}
