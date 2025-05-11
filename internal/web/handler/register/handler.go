package register

import (
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (registerHandler registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerHandler.Post(w, r)
	default:
		apiError := details.APIDetails{
			Error:      true,
			StatusCode: details.HTTPMethodNotFound,
			Message:    details.GetMessage(details.HTTPMethodNotFound),
		}
		utility.WriteHTTPJSON(w, http.StatusMethodNotAllowed, "Invalid Method", apiError)
	}
}

func NewRegisterHandler() registerHandler {
	return registerHandler{}
}

func SetRoutes(setterFunc func(string, http.Handler)) {
	setterFunc("/register", NewRegisterHandler())
}
