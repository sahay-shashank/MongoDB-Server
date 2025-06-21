package tenant

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
	"github.com/sahay-shashank/mongodb-server/internal/core/utils"
	"github.com/sahay-shashank/mongodb-server/internal/service/data"
)

func NewInsert(tenantID string, service string, httpData []byte) details.APIDetails {
	var insertRequest models.InsertRequest
	if err := json.Unmarshal(httpData, &insertRequest); err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	if err := utils.Validate.Struct(insertRequest); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Field '%s' failed validation: %s", err.Field(), err.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.ValidationFailed,
			Message:           details.GetMessage(details.ValidationFailed),
			AdditionalDetails: errorMessages,
		}
	}
	db := tenantID + "_" + service
	dataByte, err := json.Marshal(insertRequest.Data)
	if err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	documentResult := data.InsertOneDocument(db, insertRequest.Collection, dataByte)
	if documentResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.RegistrationFailure,
			Message:           details.GetMessage(details.RegistrationFailure),
			AdditionalDetails: documentResult,
		}
	}
	return details.APIDetails{
		StatusCode: details.InsertionSuccessful,
		Message:    details.GetMessage(details.InsertionSuccessful),
	}
}
