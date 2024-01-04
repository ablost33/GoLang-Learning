package exponential_backoff

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	maxAttempts = 3
	url         = "https://postman-echo.com/post"
)

type httpOperation func(url string) (*http.Response, error)

func retryFunction(operation httpOperation) (*http.Response, error) {
	for i := 0; i < maxAttempts; i++ {
		res, err := (operation)(url)
		if !shouldRetryFunc(res, err) {
			return res, err
		}
		timeSpan := time.Duration(math.Pow(2, float64(i))) * time.Second
		time.Sleep(timeSpan)
	}
	return nil, errors.New("Operation failed after all attempts.")
}

func shouldRetryFunc(response *http.Response, err error) bool {
	if err == nil {
		return false
	}
	if response != nil {
		if response.StatusCode == http.StatusBadGateway ||
			response.StatusCode == http.StatusServiceUnavailable ||
			response.StatusCode == http.StatusGatewayTimeout {
			return true
		}
	}
	return false
}

func UpdateUser(name, email string) {
	client := &http.Client{}

	person := &Person{
		Name:  name,
		Email: email,
	}
	bytedPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	putOperation := func(url string) (*http.Response, error) {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytedPerson))
		if err != nil {
			fmt.Println(err)
		}
		return client.Do(req)
	}
	response, err := retryFunction(putOperation)
	if err != nil {
		fmt.Println(err)
	}

	byteBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println(string(byteBody))
}
