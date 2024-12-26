package main

import (
	"time"
)

var activeUsers = 0
var messagesPerSecond = 0
var analyticsBroadcast = make(chan Analytics)

type Analytics struct {
	ActiveUsers       int `json:"active_users"`
	MessagesPerSecond int `json:"messages_per_second"`
}

func trackAnalytics() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			mutex.Lock()
			analytics := Analytics{
				ActiveUsers:       len(clients),
				MessagesPerSecond: messagesPerSecond,
			}
			messagesPerSecond = 0
			mutex.Unlock()
			analyticsBroadcast <- analytics
		}
	}
}

func broadcastAnalytics() {
	for {
		analytics := <-analyticsBroadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(analytics)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
