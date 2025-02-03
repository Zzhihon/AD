package utils

import (
	"AD/storage"
	"encoding/json"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

func RegisterClient(conn *websocket.Conn) {
	clients[conn] = true
}

func BroadcastResult(result storage.Prediction) {
	data, _ := json.Marshal(result)
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, data)
	}
}
