package router

import (
	"net/http"

	registerWebHandler "github.com/sahay-shashank/mongodb-server/internal/web/handler/register"
)

type RoutesInterface interface {
	getRoutes() map[string]http.Handler
	setRoutes(string, http.Handler)
}
type routeStruct struct {
	routes map[string]http.Handler
}

func (r *routeStruct) getRoutes() map[string]http.Handler {
	return r.routes
}

func (r *routeStruct) setRoutes(path string, handler http.Handler) {
	r.routes[path] = handler
}

func NewRoutes() RoutesInterface {
	routes := routeStruct{
		routes: make(map[string]http.Handler),
	}
	registerWebHandler.SetRoutes(routes.setRoutes)
	return &routes
}
