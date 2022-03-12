package main

import (
	"blog/server"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", server.New())
}
