package main

import (
	"fmt"
	"net/http"
)

func handleRequestSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm() // Parse the form data
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	studentName := r.Form.Get("studentName")
	description := r.Form.Get("description")

	// Rudimentary output for now:
	fmt.Println("Student Name:", studentName)
	fmt.Println("Description:", description)

	// Later: Replace this with your database saving logic

	fmt.Fprint(w, "Request Received! (Not yet saved)")
}
