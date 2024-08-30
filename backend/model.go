package main

import "time"

//Candle struct represents  a single OHLC(igh Low Open Close) candle

type Candle struct {
	Symbol   string    `json:"symbol"`
	Open     float64   `json:"open"`
	Close    float64   `json:"close"`
	High     float64   `json:"high"`
	Low      float64   `json"low"`
	imeStamp time.Time `json:"timestamp"`
}
