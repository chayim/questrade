package questrade

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type MarketQuotes struct {
	Quotes []SymbolQuote
}

type SymbolQuote struct {
	Symbol         int64
	SymbolID       string
	Tier           string
	BidPrice       float64
	BidSize        int64
	AskPrice       float64
	AskSize        int64
	LastTradeTrHrs float64
	LastTradePrice float64
	LastTradeTick  string
	Volume         int64
	OpenPrice      float64
	ClosePrice     float64
	HighPrice      float64
	LowPrice       float64
	Delay          bool
	IsHalted       bool
}

type Options struct {
	Quotes []OptionQuote
}

type OptionQuote struct {
	Underlying     string
	UnderlyingID   int64
	Symbol         string
	SymbolID       int64
	BidPrice       float64
	BidSize        int64
	AskPrice       float64
	AskSize        int64
	LastTradeTrHrs float64
	LastTradePrice float64
	LastTradeSize  int64
	LastTradeTick  string
	LastTradeTime  time.Time
	Volume         int64
	OpenPrice      float64
	HighPrice      float64
	LowPrice       float64
	Volatility     float64
	Delta          float64
	Gamma          float64
	Theta          float64
	Vega           float64
	Rho            float64
	OpenInterest   int64
	Delay          int64
	IsHalted       bool
	VWAP           float64
}

func (qt Questrade) SymbolQuotes(symbol ...int64) []SymbolQuote {

	symbols := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(symbol)), ","), "[]")

	quotes := MarketQuotes{}
	url := fmt.Sprintf("v1/markets/quotes?ids=%s", symbols)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &quotes)
	return quotes.Quotes
}

// TODO get the options into the post
func (qt Questrade) OptionQuotes(option ...int64) []OptionQuote {
	url := "v1/markets/quotes/options"
	options := Options{}
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}

	json.Unmarshal(res, &options)
	return options.Quotes
}
