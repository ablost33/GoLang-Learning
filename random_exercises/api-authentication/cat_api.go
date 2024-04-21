package api_authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	maxRetries = 6
)

var (
	url = "https://api.thecatapi.com/v1/images/search"
)

type GetCatRequest struct {
	Limit     int64 `json:"limit,omitempty"`
	HasBreeds bool  `json:"has_breeds,omitempty"`
}

type GetCatResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type httpOp func(url string) (*http.Response, error)

func GetCatImages(in GetCatRequest) error {
	if in.Limit != 0 {
		limit := fmt.Sprintf("?limit=%d", in.Limit)
		url = url + limit
	}
	if in.HasBreeds == true {
		has_breeds := fmt.Sprintf("&has_breeds=true")
		url = url + has_breeds
	}
	client := &http.Client{}
	getCatOp := func(url string) (*http.Response, error) {
		newRequest, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		newRequest.Header.Set("'x-api-key", "live_QxtBblkQ6M31HYbWICIRakjr0EcdOBxnNrc8M3jRr844WwSf6eT3vdrP1PTLoPjo")
		return client.Do(newRequest)
	}

	catObj, err := callWithRetries(getCatOp, url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer catObj.Body.Close()
	byteCatObj, err := io.ReadAll(catObj.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var catObjects []GetCatResponse
	if err := json.Unmarshal(byteCatObj, &catObjects); err != nil {
		log.Fatal(err)
		return err
	}

	if err := saveImages(catObjects); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func callWithRetries(op httpOp, url string) (*http.Response, error) {
	for i := 1; i <= maxRetries; i++ {
		res, err := (op)(url)
		if shouldRetry(res, err) {
			continue
		}
		return res, err
	}
	return nil, fmt.Errorf("exceeded max retries")
}

func shouldRetry(res *http.Response, err error) bool {
	if err != nil {
		return true
	}
	if res.StatusCode == http.StatusBadGateway ||
		res.StatusCode == http.StatusServiceUnavailable ||
		res.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}

func saveImages(cats []GetCatResponse) error {
	for i, cat := range cats {
		fileName := fmt.Sprintf("catFile_%d.jpg", i)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer file.Close()

		client := &http.Client{}
		downloadImageOp := func(url string) (*http.Response, error) {
			newRequest, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal(err)
			}
			return client.Do(newRequest)
		}
		res, err := callWithRetries(downloadImageOp, cat.URL)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Body.Close()
		if _, err := io.Copy(file, res.Body); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
