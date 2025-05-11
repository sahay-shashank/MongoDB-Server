package database

import (
	"context"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertOneDoc(dbName string, collectionName string, data interface{}, opts ...options.Lister[options.InsertOneOptions]) details.APIDetails {
	db := mongoClient.Database(dbName)
	collection := db.Collection(collectionName)
	result, err := collection.InsertOne(context.Background(), data,opts...)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.MongoDBDocumentInsertionFailure,
			Message:           details.GetMessage(details.MongoDBDocumentInsertionFailure),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode:        details.MongoDBDocumentInsertionSuccessful,
		Message:           details.GetMessage(details.MongoDBDocumentInsertionSuccessful),
		AdditionalDetails: result,
	}
}
