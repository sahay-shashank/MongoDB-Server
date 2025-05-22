package tenant

import (
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
	"github.com/sahay-shashank/mongodb-server/internal/service/data"
)

func NewSchema(tenantID string, service string, dataHTTP []byte) details.APIDetails {
	schemas, modelDetails := models.NewSchemaModel(dataHTTP)
	if modelDetails.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.SchemaCreationFailure,
			Message:           details.GetMessage(details.SchemaCreationFailure),
			AdditionalDetails: modelDetails,
		}
	}

	db := tenantID + "_" + service
	compareResult := data.CompareCollectionList(db, schemas.GetCollectionNames())
	if compareResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.SchemaRegistrationFailure,
			Message:           details.GetMessage(details.SchemaRegistrationFailure),
			AdditionalDetails: compareResult,
		}
	}

	validator, validationResult := schemas.ConvertToValidator()
	if validationResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.SchemaRegistrationFailure,
			Message:           details.GetMessage(details.SchemaRegistrationFailure),
			AdditionalDetails: validationResult,
		}
	}
	for collection, validatorBSON := range validator {
		collectionResult := data.CreateCollection(db, collection, validatorBSON)
		if collectionResult.Error {
			return details.APIDetails{
				Error:             true,
				StatusCode:        details.SchemaRegistrationFailure,
				Message:           details.GetMessage(details.SchemaRegistrationFailure),
				AdditionalDetails: collectionResult,
			}
		}
	}

	return details.APIDetails{
		StatusCode: details.SchemaRegistrationSuccessful,
		Message:    details.GetMessage(details.SchemaRegistrationSuccessful),
	}
}

// func DeleteSchema(tenantID string, service string, dataHTTP []byte) details.APIDetails {
// 	db := tenantID + "_" + service
// 	deleteResult := data.DeleteCollection(db,)
// 	return details.APIDetails{
// 		StatusCode: details.SchemaDeletionSuccessful,
// 		Message:    details.GetMessage(details.SchemaDeletionSuccessful),
// 	}
// }
