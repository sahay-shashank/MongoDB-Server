package database

import (
	"context"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindOneDoc(dbName string, collectionName string, decoder interface{}, filter interface{}, opts ...options.Lister[options.FindOneOptions]) details.APIDetails {
	collection := mongoClient.Database(dbName).Collection(collectionName)
	result := collection.FindOne(context.Background(), filter, opts...)
	err := result.Decode(decoder)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return details.APIDetails{
				StatusCode: details.MongoDBFindEmpty,
				Message:    details.GetMessage(details.MongoDBFindEmpty),
			}
		}
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.DecoderFaulty,
			Message:           details.GetMessage(details.DecoderFaulty),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBFindSuccessful,
		Message:    details.GetMessage(details.MongoDBFindSuccessful),
	}
}
