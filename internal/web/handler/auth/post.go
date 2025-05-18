package auth_handler

import (
	"io"
	"log"
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/service/tenant"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (authHandler *authHandler) Post(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		log.Printf("%v", err)
		apiError := details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
		utility.WriteHTTPJSON(w, http.StatusBadRequest, "Error Reading JSON body", apiError)
		return
	}
	apiResult := tenant.NewAuth(data)
	if apiResult.Error {
		utility.WriteHTTPJSON(w, http.StatusInternalServerError, "Error during Authentication", apiResult)
		return
	}
	utility.WriteHTTPJSON(w, http.StatusOK, "Authentication Completed", apiResult)
}
