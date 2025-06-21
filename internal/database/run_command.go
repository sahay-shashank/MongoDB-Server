package database

import (
	"context"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RunCommand(dbName string, command bson.D, opts ...options.Lister[options.RunCmdOptions]) details.APIDetails {
	db := mongoClient.Database(dbName)
	var result bson.M
	err := db.RunCommand(context.Background(), command, opts...).Decode(&result)
	if err != nil {
		return details.APIDetails{
			StatusCode:        details.MongoDBCommandFailure,
			Message:           details.GetMessage(details.MongoDBCommandFailure),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode:        details.MongoDBCommandSuccessful,
		Message:           details.GetMessage(details.MongoDBCommandSuccessful),
		AdditionalDetails: result,
	}
}
