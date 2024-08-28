package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

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

	// Go Routine Reading messages

	go func() {

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}

			fmt.Printf("Received:  %s\n", message)
		}

	}()

	//Go routine for sending message

	go func() {

		for {

		}

	}()
	// Read input with the terminal and send it to the websocket server
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("type Something")

	for scanner.Scan() {
		text := scanner.Text()
		//send text to channel
		send <- Message{websocket.TextMessage, []byte(text)}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner err: ", err)
	}
}
