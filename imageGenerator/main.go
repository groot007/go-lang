package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/generate", handleGenerate)

	log.Println("HTTP server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	fontStyle := r.FormValue("fontStyle")
	bgColor := r.FormValue("bgColor")
	textColor := r.FormValue("textColor")
	textSizeStr := r.FormValue("textSize")

	if name == "" || fontStyle == "" || bgColor == "" || textColor == "" || textSizeStr == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		log.Println("Missing parameters")
		return
	}

	textSize, err := strconv.Atoi(textSizeStr)
	if err != nil {
		http.Error(w, "Invalid text size", http.StatusBadRequest)
		log.Printf("Invalid text size: %v", err)
		return
	}

	generateImage(w, name, fontStyle, bgColor, textColor, textSize)
}
