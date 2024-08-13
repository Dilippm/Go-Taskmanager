package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Exported environment variables
var (
	PORT      string
	MONGO_URI string
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set environment variables
	PORT = os.Getenv("PORT")
	MONGO_URI = os.Getenv("MONGODB_URI")
}
