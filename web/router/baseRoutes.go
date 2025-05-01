package router

import (
	"net/http"
)

// type baseRouteHandlers struct {
// 	routes map[string]http.HandlerFunc
// }

var baseRouteHandlers = map[string]http.HandlerFunc{
	// add more routes here as needed
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	webHandler(w, r, baseRouteHandlers)
}
