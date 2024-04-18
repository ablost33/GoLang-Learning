package http_practice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FunkyData struct {
	UserID  int64  `json:"userId"`
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Attempt int64
}

func MakeGetRequest() {
	res, err := http.Get(GETURL1)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var dataArray FunkyData
	if err := json.Unmarshal(dataBytes, &dataArray); err != nil {
		panic(err)
	}

	fmt.Println(dataArray)
}

func (f *FunkyData) UnmarshalJSON(data []byte) error {
	// Define an auxiliary type that has the same fields as FunkyData but doesn't have a custom UnmarshalJSON method.
	type Alias FunkyData

	var returnData Alias
	if err := json.Unmarshal(data, &returnData); err != nil {
		return err
	}

	// Convert Alias back to FunkyData and perform any additional processing.
	*f = FunkyData(returnData)
	f.Attempt = 6
	return nil
}
