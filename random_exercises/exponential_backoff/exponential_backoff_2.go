package exponential_backoff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"
)

const (
	url2       = "https://dummy.restapiexample.com/api/v1/create"
	maxRetries = 5
)

func CreateEmployee() {
	data := map[string]string{
		"name":   "watermelon",
		"salary": "333",
		"age":    "12",
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	postOperation := func(url string) (*http.Response, error) {
		buffer := bytes.NewReader(reqBody)
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, buffer)
		if err != nil {
			log.Fatal(err)
		}
		return client.Do(req)
	}
	res, err := callWithRetries(postOperation, url2)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rawData))
}

type httpFunc func(url string) (*http.Response, error)

func callWithRetries(call httpFunc, url string) (*http.Response, error) {
	for i := 1; i < maxRetries; i++ {
		res, err := (call)(url)
		if shouldRetryRequest(res, err) {
			sleepTime := time.Duration(math.Pow(2, float64(i))) * time.Second
			time.Sleep(sleepTime)
		}
	}
	return nil, fmt.Errorf("exceeded max retry attempt")
}

func shouldRetryRequest(res *http.Response, err error) bool {
	if err != nil {
		return true
	}
	if res.StatusCode == http.StatusBadGateway || res.StatusCode == http.StatusServiceUnavailable || res.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}
