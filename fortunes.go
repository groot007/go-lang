package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Fortune struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

var fortunes = []Fortune{
	{Text: "You will have a great day!", Category: "motivational"},
	{Text: "Laughter is the best medicine.", Category: "funny"},
	{Text: "A mysterious event will soon unfold.", Category: "mysterious"},
	// Add more fortunes here...
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomFortune() Fortune {
	return fortunes[rand.Intn(len(fortunes))]
}

func getFortunesByCategory(category string) []Fortune {
	var filteredFortunes []Fortune
	for _, fortune := range fortunes {
		if fortune.Category == category {
			filteredFortunes = append(filteredFortunes, fortune)
		}
	}
	return filteredFortunes
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
	fortune := getRandomFortune()
	json.NewEncoder(w).Encode(fortune)
}

func handleFortuneByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("type")
	if category == "" {
		http.Error(w, "Missing category type", http.StatusBadRequest)
		return
	}

	fortunes := getFortunesByCategory(category)
	if len(fortunes) == 0 {
		http.Error(w, "No fortunes found for this category", http.StatusNotFound)
		return
	}

	fortune := fortunes[rand.Intn(len(fortunes))]
	json.NewEncoder(w).Encode(fortune)
}
