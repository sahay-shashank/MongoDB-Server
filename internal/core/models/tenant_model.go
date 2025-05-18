package models

import (
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tenant struct {
	TenantID bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OrgName  string        `bson:"org_name" json:"org_name"`
	Email    string        `bson:"email" json:"email"`
	Service  string        `bson:"service" json:"service"`
	APIKey   string        `bson:"api_key" json:"api_key"`
}

func NewTenant(req RegistrationRequest) Tenant {
	return Tenant{
		TenantID: bson.NewObjectID(),
		OrgName:  req.OrgName,
		Email:    req.Email,
		Service:  req.Service,
		APIKey:   generateAPIKey(),
	}
}

// GenerateAPIKey generates a globally unique API key using UUID
func generateAPIKey() string {
	// Generate a new UUID and return it as a string
	apiKey := uuid.New().String()
	log.Printf("Generated API Key: %s", apiKey)
	return apiKey
}
