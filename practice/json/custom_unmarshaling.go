package json

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	FunctionName string
	Temperature  float32
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type ConfigAlias Config
	tmp := struct {
		Temperature float32
		*ConfigAlias
	}{
		ConfigAlias: (*ConfigAlias)(c),
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Print(err)
		return err
	}
	c.Temperature = tmp.Temperature - 273.15
	return nil
}

type BidAsk struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type MarketDepth struct {
	Asks []BidAsk `json:"asks"`
	Bids []BidAsk `json:"bids"`
}

func (t *BidAsk) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var bidask []string
	if err := json.Unmarshal(data, &bidask); err != nil {
		return err
	}

	price, err := strconv.ParseFloat(bidask[0], 64)
	if err != nil {
		return err
	}
	amount, err := strconv.ParseFloat(bidask[1], 64)
	if err != nil {
		return err
	}
	*t = BidAsk{
		Price:  price,
		Amount: amount,
	}
	return nil
}

func GetPrices() error {
	rawData, err := os.ReadFile("json/bids.json")
	if err != nil {
		return err
	}

	marketData := &MarketDepth{}
	if err := json.Unmarshal(rawData, marketData); err != nil {
		return err
	}
	fmt.Println(marketData)
	return nil
}
