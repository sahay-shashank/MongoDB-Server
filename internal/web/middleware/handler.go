package middleware

import "net/http"

func SetMiddleware(setterFunc func(func(http.Handler) http.Handler)) {
	setterFunc(jwtMiddleware)
}

func ApplyMiddleware(handler http.Handler, middleware []func(http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
