package schema_handler

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	context_keys "github.com/sahay-shashank/mongodb-server/internal/core/context"
// 	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
// 	"github.com/sahay-shashank/mongodb-server/internal/service/tenant"
// 	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
// )

// func (schemaHandler *schemaHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	data, err := io.ReadAll(r.Body)
// 	if err != nil || len(data) == 0 {
// 		log.Printf("%v", err)
// 		apiError := details.APIDetails{
// 			Error:             true,
// 			StatusCode:        details.JSONInvalid,
// 			Message:           details.GetMessage(details.JSONInvalid),
// 			AdditionalDetails: err,
// 		}
// 		utility.WriteHTTPJSON(w, http.StatusBadRequest, "Error Reading JSON body", apiError)
// 		return
// 	}
// 	tenantIDctx := r.Context().Value(context_keys.TenantIDKey)
// 	servicectx := r.Context().Value(context_keys.ServiceKey)
// 	tenantID, oktenant := tenantIDctx.(string)
// 	service, okservice := servicectx.(string)
// 	if !oktenant || !okservice {
// 		apiError := details.APIDetails{
// 			Error:      true,
// 			StatusCode: details.ContextNotFound,
// 			Message:    details.GetMessage(details.ContextNotFound),
// 		}
// 		utility.WriteHTTPJSON(w, http.StatusInternalServerError, "Error during schema deletion", apiError)
// 		return
// 	}
// 	apiResult := tenant.DeleteSchema(tenantID, service, data)
// 	if apiResult.Error {
// 		utility.WriteHTTPJSON(w, http.StatusInternalServerError, "Error during schema deletion", apiResult)
// 		return
// 	}
// 	utility.WriteHTTPJSON(w, http.StatusOK, "Schema Registation Deleted", apiResult)
// }
