package crypto

import (
	"encoding/json"
	"flag"
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
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Supply   float64 `json:"supply"`
	PriceUsd float64 `json:"priceUsd"`
}

func (c CryptoCurrency) stringify() string {
	return fmt.Sprintf("Name: %s\nSymbol: %s\nSupply: %s\nPriceUsd: %s\n", c.Data.Name, c.Data.Symbol, c.Data.Supply, c.Data.PriceUsd)
}

func getCryptoInfo(currency string) (string, error) {
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

func Commander() {
	nameOfCrypto := flag.String(
		"crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)
	flag.Parse()

	crypto, err := getCryptoInfo(*nameOfCrypto)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(crypto)
}
