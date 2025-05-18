package auth_handler

import (
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/web/utility"
)

func (authHandler *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		authHandler.Post(w, r)
	default:
		apiError := details.APIDetails{
			Error:      true,
			StatusCode: details.HTTPMethodNotFound,
			Message:    details.GetMessage(details.HTTPMethodNotFound),
		}
		utility.WriteHTTPJSON(w, http.StatusMethodNotAllowed, "Invalid Method", apiError)
	}
}

func newAuthHandler() *authHandler {
	return &authHandler{}
}

func SetRoutes(setterFunc func(string, http.Handler, bool)) {
	setterFunc("/auth/token", newAuthHandler(), false)
}
