package router

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	routes := NewRoutes()
	for route, handler := range routes.getRoutes() {
		mux.Handle(route, handler)
	}
	return mux
}
