package database

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/net/context"
)

func ListCollections(dbName string) ([]string, error) {
	db := mongoClient.Database(dbName)
	return db.ListCollectionNames(context.Background(), bson.M{})
}
