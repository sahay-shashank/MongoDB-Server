package database

import (
	"os"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongoClient *mongo.Client

func InitDatabase() details.APIDetails {
	var err error
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return details.APIDetails{
			Error:      true,
			StatusCode: details.EnvVariableNotDefined,
			Message:    details.GetMessage(details.EnvVariableNotDefined),
			AdditionalDetails: map[string]interface{}{
				"Hint": "Set MONG0DB_URI Environment Variable",
			},
		}
	}
	clientOptions := options.Client().ApplyURI(uri)
	mongoClient, err = mongo.Connect(clientOptions)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.MongoDBConnectionFailure,
			Message:           details.GetMessage(details.MongoDBConnectionFailure),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBConnectionSuccessful,
		Message:    details.GetMessage(details.MongoDBConnectionSuccessful),
	}
}
