package models

import (
	"encoding/json"
	"time"
)

type BaseResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type Market struct {
	Name          string `json:"name"`
	MoneyPrec     string `json:"moneyPrec"`
	Stock         string `json:"stock"`
	Money         string `json:"money"`
	StockPrec     string `json:"stockPrec"`
	FeePrec       string `json:"feePrec"`
	MinAmount     string `json:"minAmount"`
	TradesEnabled bool   `json:"tradesEnabled"`
	MinTotal      string `json:"minTotal"`
}

type MarketActivity struct {
	LastUpdateTimestamp time.Time `json:"lastUpdateTimestamp"`
	TradingPairs        string    `json:"tradingPairs"`
	LastPrice           string    `json:"lastPrice"`
	LowestAsk           string    `json:"lowestAsk"`
	HighestBid          string    `json:"highestBid"`
	BaseVolume24H       string    `json:"baseVolume24h"`
	QuoteVolume24H      string    `json:"quoteVolume24h"`
	TradesEnabled       bool      `json:"tradesEnabled"`
}

type Withdrawal struct {}
