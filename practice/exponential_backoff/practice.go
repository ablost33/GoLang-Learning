package exponential_backoff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	POSTurl     = "https://postman-echo.com/post"
	contentType = "application/json"
	maxAttempt  = 3
)

func loadCSVData() [][]string {
	file, err := os.Open("./exponential_backoff/write_practice.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rawData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	stringifiedData := string(rawData)
	rowArray := strings.Split(stringifiedData, "\n")

	cleanData := [][]string{}
	for i := 1; i < len(rowArray)-1; i++ {
		cleanData = append(cleanData, strings.Split(rowArray[i], ","))
	}
	return cleanData
}

func postCSVData(data [][]string) {
	data = cleanData(data)
	byteData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var op httpOp = func(url string) (*http.Response, error) {
		buffer := bytes.NewBuffer(byteData)
		resp, err := http.Post(url, contentType, buffer)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	resp, err := exponentialRetry(op, POSTurl)
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var cleanData map[string]interface{}
	if err = json.Unmarshal(respData, &cleanData); err != nil {
		panic(err)
	}
	fmt.Println("Data is: ", cleanData)
}

func LoadAndPOST() {
	data := loadCSVData()
	postCSVData(data)
}

func cleanData(data [][]string) [][]string {
	regexpression, err := regexp.Compile("md")
	if err != nil {
		panic(err)
	}
	var cleanedData [][]string
	for _, row := range data {
		if !regexpression.MatchString(strings.Join(row, ",")) {
			cleanedData = append(cleanedData, row)
		}
	}
	return cleanedData
}

type httpOp func(url string) (*http.Response, error)

func exponentialRetry(operation httpOp, url string) (*http.Response, error) {
	for i := 1; i < maxAttempt; i++ {
		resp, err := (operation)(url)
		if err == nil {
			return resp, err
		}
		if !shouldRetryError(resp) {
			return nil, err
		}
		backoffTime := time.Duration(math.Pow(2, float64(i))) * time.Second
		time.Sleep(backoffTime)
	}
	return nil, fmt.Errorf("operation exceeded max attemps")
}

func shouldRetryError(resp *http.Response) bool {
	if resp != nil {
		if resp.StatusCode == http.StatusBadGateway ||
			resp.StatusCode == http.StatusServiceUnavailable ||
			resp.StatusCode == http.StatusGatewayTimeout {
			return true
		}
	}
	return false
}
