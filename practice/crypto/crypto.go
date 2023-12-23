package crypto

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.coincap.io/v2/assets/"
)

type CryptoCurrency struct {
	Data CryptoData `json: "data"`
}

/* @TODO: See how to convert supply and PriceUsd to another float*/
type CryptoData struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Supply   string `json:"supply"`
	PriceUsd string `json:"priceUsd"`
}

func (c CryptoCurrency) stringify() string {
	return fmt.Sprintf("Name: %s\nSymbol: %s\nSupply: %s\nPriceUsd: %s\n", c.Data.Name, c.Data.Symbol, c.Data.Supply, c.Data.PriceUsd)
}

func GetCryptoInfo(currency string) (string, error) {
	url := baseURL + currency
	httpRes, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer httpRes.Body.Close()

	var cryptoCurrency CryptoCurrency
	if err := json.NewDecoder(httpRes.Body).Decode(&cryptoCurrency); err != nil {
		return "", err
	}
	return cryptoCurrency.stringify(), nil
}
