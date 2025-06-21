package insert_handler

import (
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (insertHandler *insertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		insertHandler.Post(w, r)
	default:
		apiError := details.APIDetails{
			Error:      true,
			StatusCode: details.HTTPMethodNotFound,
			Message:    details.GetMessage(details.HTTPMethodNotFound),
		}
		utility.WriteHTTPJSON(w, http.StatusMethodNotAllowed, "Invalid Method", apiError)
	}
}

func newInsertHandler() *insertHandler {
	return &insertHandler{}
}

func SetRoutes(setterFunc func(string, http.Handler, bool)) {
	setterFunc("/insert", newInsertHandler(), false)
}
