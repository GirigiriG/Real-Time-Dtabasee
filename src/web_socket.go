package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func UpgradeConnectionToWebsocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	return conn, err
}

func ReadInboundWebsocketMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		ErrRetrievingMessage := "There was an error reading from the websocket connection"
		handleError(ErrRetrievingMessage, err)

		fmt.Println(string(message))
		go func(conn *websocket.Conn) {
			conn.WriteMessage(1, []byte(fmt.Sprintf("Hello %s from the world of sockets", message)))
		}(conn)
	}
}
