package questrade

import (
	"encoding/json"
	"fmt"
	"time"
)

type AccountOrders struct {
	AccountOrders []Order
}

type Order struct {
	Symbol                   string
	SymbolID                 int64
	Quantity                 float64
	Side                     string
	Price                    float64
	ID                       float64
	OrderID                  int64
	OrderChainID             int64
	ExchangeExecID           string
	Timestamp                time.Time
	Notes                    string
	Venue                    string
	TotalCost                float64
	OrderPlacementCommission float64
	Commission               float64
	ExecutionFee             float64
	SecFee                   float64
	CanadianExecutionFee     int64
	ParentID                 int64
}

func (qt Questrade) Orders(account string) []Order {

	orders := AccountOrders{}
	url := fmt.Sprintf("v1/accounts/%s/orders", account)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &orders)
	return orders.AccountOrders
}
