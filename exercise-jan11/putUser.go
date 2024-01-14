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
	url2        = "https://reqres.in/api/users/2"
	maxAttempts = 5
)

type EntityUpdateData struct {
	UpdatedTime time.Time `json:"updatedAt,string"`
}

type httpOperation func(url string) (*http.Response, error)

func ModifyData() error {
	data := make(map[string]string)
	data["name"] = "theGame"
	data["job"] = "rapper"

	reqData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	putOperation := func(url string) (*http.Response, error) {
		buffer := bytes.NewReader(reqData)
		client := &http.Client{}
		req, err := http.NewRequest("PUT", url, buffer)
		if err != nil {
			return nil, err
		}
		return client.Do(req)
	}

	response, err := callWithRetries(putOperation, url2)
	if err != nil {
		return err
	}

	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var entityUpdateData EntityUpdateData
	if err := json.Unmarshal(byteData, &entityUpdateData); err != nil {
		return err
	}
	fmt.Println(entityUpdateData)
	return nil
}

func callWithRetries(op httpOperation, url string) (*http.Response, error) {
	for i := 1; i <= maxAttempts; i++ {
		res, err := (op)(url)
		if !shouldRetryError(res, err) {
			return res, err
		}
		sleepTime := time.Duration(math.Pow(2, float64(i))) * time.Second
		time.Sleep(sleepTime)
	}
	return nil, fmt.Errorf("operation exceeded max number of attempts")
}

func shouldRetryError(r *http.Response, err error) bool {
	if err == nil {
		return false
	}
	if r.StatusCode == http.StatusBadGateway ||
		r.StatusCode == http.StatusServiceUnavailable ||
		r.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}
