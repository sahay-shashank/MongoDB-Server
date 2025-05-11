package tenant

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
	"github.com/sahay-shashank/mongodb-server/internal/service/data"
)

var validate = validator.New()

func NewRegister(HTTPData []byte) details.APIDetails {
	var request models.RegistrationRequest
	if err := json.Unmarshal(HTTPData, &request); err != nil {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	if err := validate.Struct(request); err != nil {
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
	query := fmt.Sprintf(`{"email": "%s"}`, request.Email)
	var tenant models.Tenant
	findResult := data.FindOneDocument("tenants", "tenant_info", &tenant, []byte(query))
	if findResult.StatusCode == details.NoDocumentFound {
		tenant = models.NewTenant(request)
		tenantBytes, err := json.Marshal(tenant)
		if err != nil {
			return details.APIDetails{
				Error:             true,
				StatusCode:        details.JSONInvalid,
				Message:           details.GetMessage(details.JSONInvalid),
				AdditionalDetails: err,
			}
		}

		indexResult := data.InsertIndex("tenants", "tenant_info", []byte(`{"email":1}`), true)
		if indexResult.Error {
			return details.APIDetails{
				Error:             true,
				StatusCode:        details.RegistrationFailure,
				Message:           details.GetMessage(details.RegistrationFailure),
				AdditionalDetails: indexResult,
			}
		}
		documentResult := data.InsertOneDocument("tenants", "tenant_info", tenantBytes)
		if documentResult.Error {
			return details.APIDetails{
				Error:             true,
				StatusCode:        details.RegistrationFailure,
				Message:           details.GetMessage(details.RegistrationFailure),
				AdditionalDetails: documentResult,
			}
		}
	}

	return details.APIDetails{
		StatusCode: details.RegistrationSuccessful,
		Message:    details.GetMessage(details.RegistrationSuccessful),
		AdditionalDetails: models.RegistrationResponse{
			APIKey: tenant.APIKey,
		},
	}
}
