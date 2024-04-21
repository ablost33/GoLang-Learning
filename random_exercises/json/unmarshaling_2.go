package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	bidFilePath = "/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json/bids.json"
)

type Txns struct {
	Asks []float64 `json:"asks,[]string"`
	Bids []float64 `json:"bids,[]string"`
}

func UnmarshalBids() {
	file, err := os.Open(bidFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output Txns
	byteData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(byteData, &output); err != nil {
		log.Fatal(err)
	}

	for _, ask := range output.Asks {
		fmt.Printf("ask is %f\n", ask)
	}

	for _, bid := range output.Bids {
		fmt.Printf("bid is %f\n", bid)
	}
}
