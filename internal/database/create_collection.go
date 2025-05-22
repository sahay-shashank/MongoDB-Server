package database

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/net/context"
)

func CreateCollection(dbName string, collectionName string, opts ...options.Lister[options.CreateCollectionOptions]) details.APIDetails {
	db := mongoClient.Database(dbName)
	errCollection := db.CreateCollection(context.Background(), collectionName, opts...)
	if errCollection != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.MongoDBCollectionFailure,
			Message:           details.GetMessage(details.MongoDBCollectionFailure),
			AdditionalDetails: errCollection,
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBCollectionSuccess,
		Message:    details.GetMessage(details.MongoDBCollectionSuccess),
	}
}
