package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // This is not prod code, but change this if it were...
}

var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	log.Printf("Client connected: %s\n", conn.LocalAddr().String())

	defer func() {
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		log.Println("Client disconnected")
	}()

	for {
		_, _, err := conn.ReadMessage() // Keep connection alive
		if err != nil {
			break
		}
	}
}

func BroadcastBoardRefresh() {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	log.Println("Sending refresh")
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("get-refresh"))
		if err != nil {
			log.Println("write:", err)
			delete(clients, client)
			client.Close()
		}
	}
}
