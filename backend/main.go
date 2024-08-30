package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

var (
	symbols = []string{"AAPL", "AMZN", "BINANCE:BTCUSDT", "IC MARKETS:1"}
)

func main() {
	env := EnvConfig()

	db := DBConnection(env)

	//Connect to finhub WebSockets
	finnhubWSConnn := connectToFinnhub(env)
	defer finnhubWSConnn.Close()

	//Handle Finhub incoming messages

	//Broadcast candle updates to all clients connected

	///----- Endpoints

	//Connect to websocket

	//Fetch all past candles for all of the symbols

	//Fetch all past candles from a specific symbol

	//Serve the Endpoints

}

//Connecting to finnhub Websockets

func connectToFinnhub(env Env) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token=cr7m6s1r01qotnb3n960cr7m6s1r01qotnb3n96g", nil)
	if err != nil {
		panic(err)
	}

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msg)
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
