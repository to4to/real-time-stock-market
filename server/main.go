package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {

	//Todo : Upgrade he incoming GET request into a webocket connection

	conn, err := upgrader.Upgrade()

}
func main() {
	http.HandleFunc("/ws", handler)
}
