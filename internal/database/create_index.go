package database

import (
	"context"
	"time"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateIndex(dbName string, collectionName string, keys bson.D, opts ...*options.IndexOptionsBuilder) details.APIDetails {
	collection := mongoClient.Database(dbName).Collection(collectionName)
	var indexOpts *options.IndexOptionsBuilder
	if len(opts) > 0 {
		indexOpts = opts[0]
	} else {
		indexOpts = options.Index() // default builder
	}
	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: indexOpts,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.MongoDBIndexCreationFailure,
			Message:           details.GetMessage(details.MongoDBDocumentInsertionSuccessful),
			AdditionalDetails: err,
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBIndexCreationSuccessful,
		Message:    details.GetMessage(details.MongoDBIndexCreationSuccessful),
	}

}
