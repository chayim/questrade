package questrade

import (
	"encoding/json"
	"fmt"
	"time"
)

type AccountActivities struct {
	Activities []Activity
}

type Activity struct {
	TradeDate       time.Time
	TransactionDate time.Time
	SettlementDate  time.Time
	Action          string
	Symbol          string
	SymbolID        int64
	Description     string
	Currency        string
	Quantity        float64
	Price           float64
	GrossAmount     float64
	Commission      float64
	NetAmount       float64
	Type            string
}

func (qt Questrade) Activities(account string, start time.Time, end time.Time) []Activity {
	startTime := start.Format(dateFmt)
	endTime := end.Format(dateFmt)

	activities := AccountActivities{}
	url := fmt.Sprintf("v1/accounts/%s/activities?startTime=%orderss&endTime=%s", account, startTime, endTime)
	res, _ := qt.request(url)
	if res == nil {
		return nil
	}
	json.Unmarshal(res, &activities)
	return activities.Activities
}
