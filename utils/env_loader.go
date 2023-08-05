package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	// Load environment variables from .env file
	// put .env inside /utils/env folder
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Get the value of an environment variable
	apiKey := os.Getenv("API_KEY")
	return apiKey
}
