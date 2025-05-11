package data

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertOneDocument(db string, collection string, doc []byte, opts ...options.Lister[options.InsertOneOptions]) details.APIDetails {
	var result bson.M
	err := bson.UnmarshalExtJSON(doc, true, &result)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	mongoDetails := database.InsertOneDoc(db, collection, result, opts...)
	if mongoDetails.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.DocumentInsertionFailure,
			Message:           details.GetMessage(details.DocumentInsertionFailure),
			AdditionalDetails: mongoDetails,
		}
	}
	return details.APIDetails{
		StatusCode:        details.DocumentInsertionSuccessful,
		Message:           details.GetMessage(details.DocumentInsertionSuccessful),
		AdditionalDetails: mongoDetails.AdditionalDetails,
	}
}

func FindOneDocument(db string, collection string, decoder interface{}, filter []byte, opts ...options.Lister[options.FindOneOptions]) details.APIDetails {
	var result bson.M
	err := bson.UnmarshalExtJSON(filter, true, &result)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	mongoDetails := database.FindOneDoc(db, collection, decoder, result, opts...)
	if mongoDetails.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.DocumentFindFailure,
			Message:           details.GetMessage(details.DocumentFindFailure),
			AdditionalDetails: mongoDetails,
		}
	} else if mongoDetails.StatusCode == details.MongoDBFindEmpty {
		return details.APIDetails{
			StatusCode:        details.NoDocumentFound,
			Message:           details.GetMessage(details.NoDocumentFound),
			AdditionalDetails: mongoDetails.AdditionalDetails,
		}
	}

	return details.APIDetails{
		StatusCode:        details.DocumentFindSuccessful,
		Message:           details.GetMessage(details.DocumentFindSuccessful),
		AdditionalDetails: mongoDetails.AdditionalDetails,
	}
}
