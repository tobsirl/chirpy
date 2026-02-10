package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// Create a new http.Server struct.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Start the server.
	srv.ListenAndServe()
}