package main

import (
	"fmt"
	"net/http"
)

func main() {
	// static files
	http.Handle("/", http.FileServer(http.Dir(".")))

	// HTTP handlers

	http.HandleFunc("/submit-request", handlers.handleRequestSubmission) // Add this line

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
