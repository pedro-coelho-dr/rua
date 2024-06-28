package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	// Verify that environment variables are loaded
	requiredEnvVars := []string{"DB_USER", "DB_PASSWORD", "DB_NAME", "DB_HOST"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Environment variable %s not set", envVar)
		}
	}
}
