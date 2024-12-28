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
    {Text: "You will achieve your goals if you stay focused.", Category: "motivational"},
    {Text: "Believe in yourself and all that you are.", Category: "motivational"},
    {Text: "Happiness is a journey, not a destination.", Category: "motivational"},
    {Text: "Your hard work will soon pay off.", Category: "motivational"},
    {Text: "A good laugh is sunshine in the house.", Category: "funny"},
    {Text: "You will find humor in unexpected places.", Category: "funny"},
    {Text: "A joke a day keeps the doctor away.", Category: "funny"},
    {Text: "Your laughter is contagious.", Category: "funny"},
    {Text: "An unexpected event will change your life.", Category: "mysterious"},
    {Text: "A secret will be revealed to you soon.", Category: "mysterious"},
    {Text: "You will encounter a mysterious stranger.", Category: "mysterious"},
    {Text: "A hidden talent will be discovered.", Category: "mysterious"},
    {Text: "You will inspire others with your actions.", Category: "motivational"},
    {Text: "Your positive attitude will lead to success.", Category: "motivational"},
    {Text: "You are capable of achieving great things.", Category: "motivational"},
    {Text: "Your determination will bring you success.", Category: "motivational"},
    {Text: "A funny situation will make you laugh today.", Category: "funny"},
    {Text: "You will find joy in the little things.", Category: "funny"},
    {Text: "A humorous encounter will brighten your day.", Category: "funny"},
    {Text: "Your wit will impress those around you.", Category: "funny"},
    {Text: "A mysterious opportunity will present itself.", Category: "mysterious"},
    {Text: "You will uncover a hidden truth.", Category: "mysterious"},
    {Text: "A surprising event will bring you joy.", Category: "mysterious"},
    {Text: "You will be intrigued by a new discovery.", Category: "mysterious"},
    {Text: "Your perseverance will lead to success.", Category: "motivational"},
    {Text: "You have the power to create your own destiny.", Category: "motivational"},
    {Text: "Your efforts will be rewarded.", Category: "motivational"},
    {Text: "You will overcome any obstacles in your path.", Category: "motivational"},
    {Text: "A funny story will make you smile.", Category: "funny"},
    {Text: "You will find humor in everyday situations.", Category: "funny"},
    {Text: "A lighthearted moment will lift your spirits.", Category: "funny"},
    {Text: "Your laughter will bring joy to others.", Category: "funny"},
    {Text: "A mysterious journey awaits you.", Category: "mysterious"},
    {Text: "You will be captivated by a new experience.", Category: "mysterious"},
    {Text: "A hidden message will be revealed to you.", Category: "mysterious"},
    {Text: "You will be drawn to the unknown.", Category: "mysterious"},
    {Text: "Your positive energy will attract success.", Category: "motivational"},
    {Text: "You have the strength to achieve your dreams.", Category: "motivational"},
    {Text: "Your hard work will lead to great rewards.", Category: "motivational"},
    {Text: "You will find success in your endeavors.", Category: "motivational"},
    {Text: "A funny coincidence will make you laugh.", Category: "funny"},
    {Text: "You will find joy in unexpected moments.", Category: "funny"},
    {Text: "A humorous situation will bring you happiness.", Category: "funny"},
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