package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    // API routes
    http.HandleFunc("/api/fortune", handleFortune)
    http.HandleFunc("/api/fortune/category", handleFortuneByCategory)

    // Add CORS middleware
    corsHandler := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            
            h.ServeHTTP(w, r)
        })
    }

    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, corsHandler(http.DefaultServeMux)); err != nil {
        log.Fatal(err)
    }
}