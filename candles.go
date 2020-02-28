package questrade

import (
	"encoding/json"
	"fmt"
	"time"
)

type Candle struct {
	Start      time.Time
	End        time.Time
	OpenPrice  float64 `json:"open"`
	HighPrice  float64 `json:"high"`
	LowPrice   float64 `json:"low"`
	ClosePrice float64 `json:"close"`
	Volume     int64   `json:"volume"`
}

type Candles struct {
	Candles []Candle `json:"candles"`
}

func (qt Questrade) Candle(symbol int64, start time.Time, end time.Time) []Candle {
	jd := Candles{}
	startTime := start.Format(dateFmt)
	endTime := end.Format(dateFmt)
	url := fmt.Sprintf("v1/markets/candles/%d?startTime=%s&endTime=%s&interval=OneMinute", symbol, startTime, endTime)
	res, _ := qt.request(url)
	if res == nil {
		return []Candle{}
	}

	json.Unmarshal(res, &jd)
	return jd.Candles
}
