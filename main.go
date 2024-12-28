package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

var rooms = make(map[string]*Room)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		http.Error(w, "Missing room ID", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	if _, ok := rooms[roomID]; !ok {
		rooms[roomID] = &Room{
			clients:   make(map[*websocket.Conn]bool),
			broadcast: make(chan []byte),
		}
		go handleMessages(roomID)
	}

	room := rooms[roomID]
	room.clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(room.clients, ws)
			break
		}
		room.broadcast <- msg
	}
}

func handleMessages(roomID string) {
	room := rooms[roomID]
	for {
		msg := <-room.broadcast
		for client := range room.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(room.clients, client)
			}
		}
	}
}

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

	http.HandleFunc("/ws", handleConnections)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, corsHandler(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}