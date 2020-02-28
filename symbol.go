package questrade

import (
	"encoding/json"
	"fmt"
	"time"
)

type Symbol struct {
	Symbol []EquitySymbol
}

type EquitySymbol struct {
	Symbol          string
	SymbolID        int64
	Description     string
	SecurityType    string
	ListingExchange string
	IsQuotable      bool
	IsTradable      bool
	Currency        string
}

type Option struct {
	Option []OptionSymbol
}

type OptionSymbol struct {
	ExpiryDate         time.Time
	Description        string
	ListingExchange    string
	OptionExerciseType string
	ChainPerRoot       []OptionSymbolChain
}

type OptionSymbolChain struct {
	StrikePrice  float64
	CallSymbolID int64
	PutSymbolID  int64
}

func (qt Questrade) Symbols(symbol string) []EquitySymbol {

	symbols := Symbol{}
	url := fmt.Sprintf("v1/symbols/search?prefix=%s", symbol)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &symbols)
	return symbols.Symbol
}

func (qt Questrade) OptionSymbols(symbol int64) []OptionSymbol {
	options := Option{}
	url := fmt.Sprintf("v1/symbols/%d/options", symbol)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &options)
	return options.Option
}
