package router

import (
	"net/http"

	"github.com/sahay-shashank/mongodb-server/internal/web/middleware"
)

type MiddlewareInterface interface {
	getMiddleware() []func(http.Handler) http.Handler
	use(func(http.Handler) http.Handler)
}

type middlewareStruct struct {
	middleware []func(http.Handler) http.Handler
}

func (m *middlewareStruct) getMiddleware() []func(http.Handler) http.Handler {
	return m.middleware
}

func (m *middlewareStruct) use(mw func(http.Handler) http.Handler) {
	m.middleware = append(m.middleware, mw)
}

func NewMiddleware() MiddlewareInterface {
	mw := &middlewareStruct{
		middleware: make([]func(http.Handler) http.Handler, 0),
	}
	middleware.SetMiddleware(mw.use)
	return mw
}
