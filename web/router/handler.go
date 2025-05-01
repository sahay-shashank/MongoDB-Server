package router

import (
	"net/http"
	"strings"
)

func webHandler(w http.ResponseWriter, r *http.Request, routeHandlers map[string]http.HandlerFunc) {
	uriPath := strings.Trim(r.URL.Path, "/")
	segments := strings.Split(uriPath, "/")
	if len(segments) == 0 || segments[0] == "" {
		http.NotFound(w, r)
		return
	}

	if handler, ok := routeHandlers[segments[0]]; ok {
		remainingPath := "/" + strings.Join(segments[1:], "/")
		r.URL.Path = remainingPath
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}