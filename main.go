package main

import (
	"fmt"
	"net/http"
)

// handler function for the root URL
func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML content to the response
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Hello, Go!</title>
    </head>
    <body>
        <h1>Hello, Go Web Server!</h1>
        <p>This is a simple web server written in Go.</p>
    </body>
    </html>`

	fmt.Fprintln(w, html)
}

func main() {
	// Define the route and handler
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
