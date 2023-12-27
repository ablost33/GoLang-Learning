package exponential_backoff

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
	RetryCount = 3
	baseURL    = "https://api.coincap.io/v2/assets/"
)

type retryableTransport struct {
	transport http.RoundTripper
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

func shouldRetry(err error, response *http.Response) bool {
	if err != nil {
		return true
	}
	if response.StatusCode == http.StatusBadGateway ||
		response.StatusCode == http.StatusServiceUnavailable ||
		response.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}

func drainBody(resp *http.Response) {
	/* So really all we're doing here is discarding the io reader's body */
	if resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
}

func (t *retryableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	resp, err := t.transport.RoundTrip(req)

	retries := 0
	for shouldRetry(err, resp) && retries < RetryCount {
		time.Sleep(exponentialBackoff(retries))
		drainBody(resp)
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		resp, err = t.transport.RoundTrip(req)
		retries++
	}
	return resp, err
}

func NewRetryableClient() *http.Client {
	transport := &retryableTransport{
		transport: &http.Transport{},
	}
	return &http.Client{
		Transport: transport,
	}
}

func CallCoinexAPI() {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := NewRetryableClient()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	data := &CryptoData{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		fmt.Println(err)
	}

	fmt.Println("The crypto data is ", data)
}
