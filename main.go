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

    http.HandleFunc("/fortune", handleFortune)
    http.HandleFunc("/fortune/category", handleFortuneByCategory)

    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}