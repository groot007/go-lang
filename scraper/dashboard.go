package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type DashboardMessage struct {
	Type    string        `json:"type"`
	URLs    []string      `json:"urls,omitempty"`
	Results []ScrapeResult `json:"results,omitempty"`
}

func handleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	for {
		var msg DashboardMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("ReadJSON: %v", err)
			break
		}

		if msg.Type == "start" {
			go func() {
				results := startScraping(msg.URLs, 5)
				response := DashboardMessage{Type: "results", Results: results}
				ws.WriteJSON(response)
			}()
		}
	}
}
