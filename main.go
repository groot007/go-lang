package main

import (
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/fortune":
		handleFortune(w, r)
	case "/fortune/category":
		handleFortuneByCategory(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/fortune", handleFortune)
	http.HandleFunc("/fortune/category", handleFortuneByCategory)

	log.Println("HTTP server started on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
