package register

import (
	"io"
	"log"
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/service/tenant"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (registerHandler registerHandler) Post(w http.ResponseWriter, r *http.Request) {
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
	apiResult := tenant.NewRegister(data)
	if apiResult.Error {
		utility.WriteHTTPJSON(w, http.StatusInternalServerError, "Error during registration", apiResult)
		return
	}
	utility.WriteHTTPJSON(w, http.StatusOK, "Registation Completed", apiResult)
}
