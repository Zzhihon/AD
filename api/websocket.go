package api

import (
	"AD/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	utils.RegisterClient(conn)
}
