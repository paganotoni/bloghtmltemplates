package server

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderPost)

	return mux
}
