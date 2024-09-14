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
	Data []TradeData `json:"data"`
	Type string      `json:"type"` //trade
}

type TradeData struct {
	Close     []string  `json:"c"`
	Price     []float64 `json:"p"`
	Symbol    string    `json:"s"`
	TimeStamp int64     `json:"t"`
	Volume    int       `json:"v"`
}

//Data to write to clients connected

type BroadCastMessage struct {
	UpdateType UpdateType
}

type UpdateType string

const (
	Live   UpdateType = "live"
	Closed UpdateType = "closed"
)

type TempCandle struct {
	Symbol string

	OpenTime   time.Time
	CloseTime  time.Time
	OpenPrice  float64
	ClosePrice float64
	HighPrice  float64
	LowPrice   float64
	Volume     float64
}

//converts temp candle into a candle

func (tc *TempCandle) toCandle() *Candle {
	return &Candle{
		Symbol:    tc.Symbol,
		Open:      tc.OpenPrice,
		Close:     tc.ClosePrice,
		High:      tc.HighPrice,
		Low:       tc.LowPrice,
		TimeStamp: tc.CloseTime,
	}
}
