package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DatabaseDSN string
}

var (
	AppConfig *config
)

// LoadConfig loads the config from the .env file
func LoadConfig() {
	// Load .env file
	godotenv.Load()

	// Create config struct
	AppConfig = &config{
		// Get database connection string from .env file to AppConfig.DatabaseDSN (string)
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
	}
}
