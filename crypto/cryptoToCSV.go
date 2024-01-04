package crypto

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coin struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Supply   float64 `json:"supply"`
	PriceUSD float64 `json:"priceUsd"`
}

type CoinData struct {
	Data      []Coin `json:"data"`
	Timestamp int    `json:"timestamp"`
}

func getCryptoData() []Coin {
	res, err := http.Get(baseURL)
	if err != nil {
		panic(err)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	coinData := CoinData{}
	if err := json.Unmarshal(byteData, &coinData); err != nil {
		panic(err)
	}
	return coinData.Data
}

func (coin *Coin) UnmarshalJSON(data []byte) error {
	var rawCoin struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Supply   string `json:"supply"`
		PriceUSD string `json:"priceUsd"`
	}
	if err := json.Unmarshal(data, &rawCoin); err != nil {
		return err
		panic(err)
	}
	supplyFloat, _ := strconv.ParseFloat(rawCoin.Supply, 64)
	priceUSDFloat, _ := strconv.ParseFloat(rawCoin.Supply, 64)
	*coin = Coin{
		Name:     rawCoin.Name,
		Symbol:   rawCoin.Symbol,
		Supply:   supplyFloat,
		PriceUSD: priceUSDFloat,
	}
	return nil
}

func writeToCSV(data []Coin) {
	file, err := os.Create("./crypto/crypto_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Name", "Symbol", "Supply", "PriceUSD"}
	_, err = writer.WriteString(strings.Join(headers, ",") + "\n")
	if err != nil {
		panic(err)
	}

	regex, err := regexp.Compile("Bitcoin|Ethereum|Cardano")
	for _, coin := range data {
		stringifiedCoin := fmt.Sprintf("%s,%s,%.4f, %.4f", coin.Name, coin.Symbol, coin.Supply, coin.PriceUSD) + "\n"
		if regex.MatchString(stringifiedCoin) {
			_, err := writer.WriteString(stringifiedCoin)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CryptoToCSV() {
	data := getCryptoData()
	writeToCSV(data)
}
