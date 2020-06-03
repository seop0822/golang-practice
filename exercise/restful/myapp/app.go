package myapp

import "net/http"

// NewHnalder make a new myapp Handler
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	return mux
}
