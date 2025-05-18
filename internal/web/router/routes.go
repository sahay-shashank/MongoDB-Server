package router

import (
	"net/http"

	authWebHandler "github.com/sahay-shashank/mongodb-server/internal/web/handler/auth"
	registerWebHandler "github.com/sahay-shashank/mongodb-server/internal/web/handler/register"
)

type RoutesInterface interface {
	getRoutes() map[string]route
	setRoutes(string, http.Handler, bool)
}
type routeStruct struct {
	routes map[string]route
}

type route struct {
	path      string
	handler   http.Handler
	protected bool
	//TODO: Introduce specific middlewares that will be fed by the route itself
	//TODO: Introduce grouping api, e.g.: /api/v1
}

func (r *routeStruct) getRoutes() map[string]route {
	return r.routes
}

func (r *routeStruct) setRoutes(pathString string, handlerHTTP http.Handler, protectedBool bool) {
	r.routes[pathString] = route{
		path:      pathString,
		handler:   handlerHTTP,
		protected: protectedBool,
	}
}

func NewRoutes() RoutesInterface {
	routes := routeStruct{
		routes: make(map[string]route),
	}
	registerWebHandler.SetRoutes(routes.setRoutes)
	authWebHandler.SetRoutes(routes.setRoutes)
	return &routes
}
