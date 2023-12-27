package exponential_backoff

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type CryptoData struct {
	Data []CoinData `json:"data"`
}

type CoinData struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Supply   float64 `json:"supply"`
	PriceUsd float64 `json:"priceUsd"`
}

func (crypto *CoinData) UnmarshalJSON(data []byte) error {
	var rawData struct {
		Name     string `json:"name"`
		Supply   string `json:"supply"`
		Symbol   string `json:"symbol"`
		PriceUsd string `json:"priceUsd"`
	}
	if err := json.Unmarshal(data, &rawData); err != nil {
		fmt.Println(err)
	}
	supply, _ := strconv.ParseFloat(rawData.Supply, 64)
	priceUSD, _ := strconv.ParseFloat(rawData.PriceUsd, 64)
	*crypto = CoinData{
		Name:     rawData.Name,
		Symbol:   rawData.Symbol,
		Supply:   supply,
		PriceUsd: priceUSD,
	}
	return nil
}

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
