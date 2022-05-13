package main

import (
	"log"
	"net/http"
)

func handleWebSocketUpgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := UpgradeConnectionToWebsocket(w, r)

	ErrUpgrading := "error occured while trying to upgrade to websocket"
	handleError(ErrUpgrading, err)
	ReadInboundWebsocketMessages(conn)
}

func handleError(msg string, err error) {
	if err != nil {
		log.Panic(msg, err)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocketUpgrade)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello gideon"))
	})
	http.ListenAndServe(":3000", nil)
}
