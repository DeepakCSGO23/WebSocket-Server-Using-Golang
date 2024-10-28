package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader configures the upgrade from HTTP to Websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handling websocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	// Log and echo the message back to the client
	log.Printf("Client connected!")
	// Infinite loop to keep reading messages
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Client disconnected")
			break
		}
		// Log and echo the message back to the client
		log.Printf("Received: %s\n", message)
		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
func main() {
	http.HandleFunc("/ws", handleConnections)
	// Start the server on port 5000
	fmt.Println("Websocket server started on port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
