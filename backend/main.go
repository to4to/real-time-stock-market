package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var (
	symbols = []string{"AAPL", "AMZN", "BINANCE:BTCUSDT", "IC MARKETS:1"}
)

func main() {
	env := EnvConfig()

	db := DBConnection(env)

	//Connect to finhub WebSockets
	finnhubWSConnn := connectToFinnhub(*env)
	defer finnhubWSConnn.Close()

	//Handle Finhub incoming messages
	go handleFinnhubMeaasges(finnhubWSConnn, db)
	//Broadcast candle updates to all clients connected

	///----- Endpoints

	//Connect to websocket

	//Fetch all past candles for all of the symbols

	//Fetch all past candles from a specific symbol

	//Serve the Endpoints

}

//Connecting to finnhub Websockets

func connectToFinnhub(env Env) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://ws.finnhub.io?token=%s", env.API_KEY), nil)
	if err != nil {
		panic(err)
	}

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		ws.WriteMessage(websocket.TextMessage, msg)
	}

	// var msg interface{}
	// for {
	// 	err := w.ReadJSON(&msg)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Message from server ", msg)
	// }

	return ws
}

func handleFinnhubMeaasges(ws *websocket.Conn, db *gorm.DB) {
	finnhubMessage := &FinnhubMessage{}
	for {
		if err := ws.ReadJSON(finnhubMessage); err != nil {
			fmt.Println("Error Reading the message %e", err)
			continue
		}
	}

}
