package exercise_jan11

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	url        = "https://reqres.in/api/users"
	maxAttempt = 5
)

func PostUserWithClient() error {
	data := `{"name":"bagman","job":"leader"}`

	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	postOp := func(url string) (response *http.Response, err error) {
		buffer := bytes.NewReader(reqBody)
		client := http.Client{}
		req, err := http.NewRequest("POST", url, buffer)
		if err != nil {
			return nil, err
		}
		return client.Do(req)
	}

	res, err := callWithRetry(postOp, url)
	if err != nil {
		println(err)
		return err
	}

	byteResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	println(res.StatusCode)

	println(string(byteResponse))
	return nil
}

type httpReq func(url string) (response *http.Response, err error)

func callWithRetry(aFunc httpReq, url string) (response *http.Response, err error) {
	for i := 1; i <= maxAttempt; i++ {
		fmt.Printf("attempt: %d\n", i)
		res, err := (aFunc)(url)
		if shouldRetry(res.StatusCode) {
			backoff := time.Duration(math.Pow(2, float64(i))) * time.Second
			time.Sleep(backoff)
			continue
		}
		return res, err
	}
	return nil, fmt.Errorf("Exceeded max number of attempts on API")
}

func shouldRetry(errorCode int) bool {
	if errorCode == http.StatusBadGateway ||
		errorCode == http.StatusServiceUnavailable ||
		errorCode == http.StatusInternalServerError ||
		errorCode == http.StatusBadRequest {
		return true
	}
	return false
}
