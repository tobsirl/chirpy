package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	cfg := &apiConfig{}

	mux := http.NewServeMux()
	// Create a new handler that writes the number of requests that have been counted as plain text in this format to the HTTP response:
// Hits: x
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		hits := cfg.fileserverHits.Load()
		w.Write([]byte(fmt.Sprintf("Hits: %d", hits)))
	})

	// Finally, create and register a handler on the /reset path that, when hit, will reset your fileserverHits back to 0.
	mux.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Store(0)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits: 0"))
	})
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	// configure the filepath so that the image is accessible from http://localhost:8080/assets/logo.png
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepathRoot+"/assets"))))	
	// The endpoint should be accessible at the /healthz path using any HTTP method.
	
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}