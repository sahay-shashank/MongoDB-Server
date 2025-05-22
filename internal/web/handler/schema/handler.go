package schema_handler

import (
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (schemaHandler *schemaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		schemaHandler.Post(w, r)
	// case http.MethodDelete:
	// 	schemaHandler.Delete(w, r)
	default:
		apiError := details.APIDetails{
			Error:      true,
			StatusCode: details.HTTPMethodNotFound,
			Message:    details.GetMessage(details.HTTPMethodNotFound),
		}
		utility.WriteHTTPJSON(w, http.StatusMethodNotAllowed, "Invalid Method", apiError)
	}
}

func newSchemaHandler() *schemaHandler {
	return &schemaHandler{}
}

func SetRoutes(setterFunc func(string, http.Handler, bool)) {
	setterFunc("/schema", newSchemaHandler(), true)
}
