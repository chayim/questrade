package questrade

import (
	"encoding/json"
	"fmt"
)

type AccountPositions struct {
	AccountPositions []Position
}

type Position struct {
	Symbol             string
	SymbolID           int64
	OpenQuantity       float64
	ClosedQuantity     float64
	CurrentMarketValue float64
	CurrentPrice       float64
	AverageEntryPrice  float64
	ClosedPnL          float64
	OpenPnL            float64
	TotalCost          float64
	IsRealTime         bool
	IsUnderReorg       bool
}

func (qt Questrade) Positions(account string) []Position {

	positions := AccountPositions{}
	url := fmt.Sprintf("v1/accounts/%s/activities", account)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &positions)
	return positions.AccountPositions
}
