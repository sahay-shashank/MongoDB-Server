package tenant

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sahay-shashank/mongodb-server/internal/core/auth"
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/models"
	"github.com/sahay-shashank/mongodb-server/internal/service/data"
)

func NewAuth(httpData []byte) details.APIDetails {
	var request models.AuthRequest
	if err := json.Unmarshal(httpData, &request); err != nil {
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
	query := fmt.Sprintf(`{"api_key": "%s"}`, request.APIKey)
	var tenant models.Tenant
	findResult := data.FindOneDocument("tenants", "tenant_info", &tenant, []byte(query))
	if findResult.StatusCode == details.NoDocumentFound {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.AuthFailure,
			Message:           details.GetMessage(details.AuthFailure),
			AdditionalDetails: "API Key invalid",
		}
	}
	apiResult := auth.GenerateJWTToken(tenant.TenantID.Hex(), tenant.Service)
	if apiResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.AuthFailure,
			Message:           details.GetMessage(details.AuthFailure),
			AdditionalDetails: apiResult,
		}
	}
	return details.APIDetails{
		StatusCode:        details.AuthSuccessful,
		Message:           details.GetMessage(details.AuthSuccessful),
		AdditionalDetails: apiResult.AdditionalDetails,
	}
}
