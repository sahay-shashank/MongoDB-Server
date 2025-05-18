package router

import (
	"net/http"

	"github.com/sahay-shashank/mongodb-server/internal/web/middleware"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	routes := NewRoutes()
	middlewares := NewMiddleware()
	for routePath, route := range routes.getRoutes() {
		handler := route.handler
		if route.protected {
			handler = middleware.ApplyMiddleware(handler, middlewares.getMiddleware())
		}
		mux.Handle(routePath, handler)
	}
	return mux
}
