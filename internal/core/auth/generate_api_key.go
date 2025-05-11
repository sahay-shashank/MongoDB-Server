package auth

import (
	"log"

	"github.com/google/uuid"
)

// GenerateAPIKey generates a globally unique API key using UUID
func GenerateAPIKey() string {
	// Generate a new UUID and return it as a string
	apiKey := uuid.New().String()
	log.Printf("Generated API Key: %s", apiKey)
	return apiKey
}
