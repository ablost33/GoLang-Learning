package crypto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	url      = "https://api.coinex.com/v1/market/list"
	filepath = "crypto/financial_data.json"
)

type AllMarketList struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func GetCoinexData() error {
	httpRes, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var allMarketList AllMarketList
	if err := json.NewDecoder(httpRes.Body).Decode(&allMarketList); err != nil {
		fmt.Println(err)
		return err
	}
	defer httpRes.Body.Close()
	fmt.Println(allMarketList)

	return nil
}

type DailyMarketData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    TickerData
}

type TickerData struct {
	ServerTime CustomTime `json:"data"`
	Open       float64    `json:"open"`
	Close      float64    `json:"last"`
	Volume     float64    `json:"vol"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	timeValue := time.UnixMilli(i)
	*t = CustomTime{timeValue}
	return nil
}

func (d *TickerData) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var rawTicker struct {
		ServerTime CustomTime `json:"date"`
		Ticker     struct {
			Open   string `json:"open"`
			Close  string `json:"last"`
			Volume string `json:"vol"`
		} `json:"ticker"`
	}
	if err := json.Unmarshal(data, &rawTicker); err != nil {
		return err
	}
	fmt.Println("Raw ticker is: ", rawTicker)

	*d = TickerData{
		ServerTime: rawTicker.ServerTime,
	}
	d.Open, _ = strconv.ParseFloat(rawTicker.Ticker.Open, 64)
	d.Close, _ = strconv.ParseFloat(rawTicker.Ticker.Close, 64)
	d.Volume, _ = strconv.ParseFloat(rawTicker.Ticker.Volume, 64)

	return nil
}

func GetTodayData() error {
	rawData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	data := &DailyMarketData{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)
	return nil
}
