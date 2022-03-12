package main

import (
	"blog/server"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(">> Starting server on port 8080 <<")
	http.ListenAndServe(":8080", server.New())
}
