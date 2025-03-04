package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Olesia! You have created your first HTTP server!")
}

func main() {
	http.HandleFunc("/", handler) // Handle requests to the root path "/"
	fmt.Println("Server started on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server on port 8080
}
