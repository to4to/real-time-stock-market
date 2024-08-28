package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Message struct {
	MessageType int
	Data        []byte
}

func main() {
	//connect to remote websocket
	u := url.URL{Scheme: "ws", Host: "loalhost:3000", Path: "/ws"}

	fmt.Printf("Connecting To %s\n ", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial: ", err)
	}

	defer conn.Close()

	//Channels for managing messages

	send := make(chan Message)

	done := make(chan struct{})

}
