package data

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateCollection(db string, collection string, validator bson.M) details.APIDetails {
	//TODO: Take schema and convert to validator of bson format
	collectionOptions := options.CreateCollection().SetValidator(validator)
	createResult := database.CreateCollection(db, collection, collectionOptions)

	if createResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.CollectionCreationFailure,
			Message:           details.GetMessage(details.CollectionCreationFailure),
			AdditionalDetails: createResult,
		}
	}

	return details.APIDetails{
		StatusCode: details.CollectionCreationSuccess,
		Message:    details.GetMessage(details.CollectionCreationSuccess),
	}
}

func CompareCollectionList(db string, collectionNames []string) details.APIDetails {
	mongoCollection, err := database.ListCollections(db)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.FindCollectionFailure,
			Message:           details.GetMessage(details.FindCollectionFailure),
			AdditionalDetails: err,
		}
	}
	set := make(map[string]struct{}, len(mongoCollection))
	for _, item := range mongoCollection {
		set[item] = struct{}{}
	}

	var matches []string
	for _, item := range collectionNames {
		if _, exists := set[item]; exists {
			matches = append(matches, item)
		}
	}

	if len(matches) > 0 {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.CollectionExists,
			Message:           details.GetMessage(details.CollectionExists),
			AdditionalDetails: matches,
		}
	}

	return details.APIDetails{
		StatusCode: details.CollectionNotFound,
		Message:    details.GetMessage(details.CollectionNotFound),
	}
}

func DeleteCollection(db string, collection string) details.APIDetails {
	deleteResult := database.DeleteCollection(db, collection)
	if deleteResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.CollectionDeletionFailure,
			Message:           details.GetMessage(details.CollectionDeletionFailure),
			AdditionalDetails: deleteResult,
		}
	}

	return details.APIDetails{
		StatusCode: details.CollectionDeletionSuccess,
		Message:    details.GetMessage(details.CollectionDeletionSuccess),
	}
}
