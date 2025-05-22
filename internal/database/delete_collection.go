package database

import (
	"context"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
)

func DeleteCollection(dbName string, collectionName string) details.APIDetails {
	db := mongoClient.Database(dbName)
	errDeletion := db.Collection(collectionName).Drop(context.Background())
	if errDeletion != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.MongoDBCollectionDeletionFailure,
			Message:           details.GetMessage(details.MongoDBCollectionDeletionFailure),
			AdditionalDetails: errDeletion,
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBCollectionDeletionSuccess,
		Message:    details.GetMessage(details.MongoDBCollectionDeletionSuccess),
	}
}
