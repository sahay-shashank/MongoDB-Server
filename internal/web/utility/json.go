package utility

import (
	"encoding/json"
	"net/http"

	details "github.com/sahay-shashank/mongodb-server/internal/core/details"
)

func WriteHTTPJSON(w http.ResponseWriter, status int, shortMessage string, apiDetails details.APIDetails) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	httpJSONDetails := details.HTTPJSONDetails{
		HTTPStatusCode: status,
		Message:        shortMessage,
		Details:        apiDetails,
	}
	err := json.NewEncoder(w).Encode(httpJSONDetails)
	if err != nil {
		http.Error(w, `{ "error" : "Failed to encode JSON" }`, http.StatusInternalServerError)
	}
}
