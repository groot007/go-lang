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
    {Text: "A smile will lead to great opportunities.", Category: "motivational"},
    {Text: "Laughter is the best medicine.", Category: "funny"},
    {Text: "Your sense of humor will make someone happy today.", Category: "funny"},
    {Text: "A mysterious event will soon unfold.", Category: "mysterious"},
    {Text: "A strange visitor will bring unexpected news.", Category: "mysterious"},
}

func init() {
    // Modern way to seed random number generator
    rand.New(rand.NewSource(time.Now().UnixNano()))
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
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    fortune := getRandomFortune()
    if err := json.NewEncoder(w).Encode(fortune); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}

func handleFortuneByCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

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
    if err := json.NewEncoder(w).Encode(fortune); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}