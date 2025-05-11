package data

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertIndex(db string, collection string, indexData []byte, unique bool) details.APIDetails {
	var data bson.D

	// Convert JSON []byte to bson.D (ordered map)
	if err := bson.UnmarshalExtJSON(indexData, true, &data); err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.IndexInsertionFailure,
			Message:           details.GetMessage(details.IndexInsertionFailure),
			AdditionalDetails: err,
		}
	}
	var uniqueOpts *options.IndexOptionsBuilder
	if unique {
		uniqueOpts = options.Index().SetUnique(true)
	}

	output := database.CreateIndex(db, collection, data, uniqueOpts)
	if output.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.IndexInsertionFailure,
			Message:           details.GetMessage(details.IndexInsertionFailure),
			AdditionalDetails: output,
		}
	}
	return details.APIDetails{
		StatusCode:        details.IndexInsertionSuccessful,
		Message:           details.GetMessage(details.IndexInsertionSuccessful),
		AdditionalDetails: output,
	}
}
