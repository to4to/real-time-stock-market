package main

import "time"

//Candle struct represents  a single OHLC(igh Low Open Close) candle

type Candle struct {
	Symbol    string    `json:"symbol"`
	Open      float64   `json:"open"`
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json"low"`
	TimeStamp time.Time `json:"timestamp"`
}

// structure of data that comes from api
type FinnhubMessage struct {
	Data []TradeData
	Type string `json:"type"` //trade
}

type TradeData struct {
	Close     []string  `json:"c"`
	Price     []float64 `json:"p"`
	Symbol    string    `json:"s"`
	TimeStamp int64     `json:"t"`
	Volume    int       `json:"c"`
}
