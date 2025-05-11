package models

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/auth"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tenant struct {
	TenantID bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`

	OrgName string `bson:"org_name" json:"org_name"`
	Email   string `bson:"email" json:"email"`
	APIKey  string `bson:"api_key" json:"api_key"`
}

func NewTenant(req RegistrationRequest) Tenant {
	return Tenant{
		TenantID: bson.NewObjectID(),
		OrgName:  req.OrgName,
		Email:    req.Email,
		APIKey:   auth.GenerateAPIKey(),
	}
}
