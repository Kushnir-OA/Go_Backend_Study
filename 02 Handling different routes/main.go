package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the 'About' page. Here you can learn more about our project.")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the 'Contact' page. Write to us at email: example@example.com")
}

func main() {
	http.HandleFunc("/", homeHandler)           // Route for the Home page
	http.HandleFunc("/about", aboutHandler)     // Route for the About page
	http.HandleFunc("/contact", contactHandler) // Route for the Contact page

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
