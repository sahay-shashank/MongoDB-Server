package router

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", baseHandler)
	return mux
}
