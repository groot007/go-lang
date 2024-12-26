package main

import (
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

	http.HandleFunc("/", Handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":"+port, nil)
}
