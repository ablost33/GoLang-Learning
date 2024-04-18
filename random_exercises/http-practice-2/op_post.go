package http_practice_2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	url = "https://reqres.in/api/users/2"
)

// A simple PUT operation
type Job struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func PutJobData(job *Job) error {
	reqData, err := json.Marshal(job)
	if err != nil {
		return err
	}

	putOperation := func(url string) (*http.Response, error) {
		buffer := bytes.NewReader(reqData)
		httpReq, err := http.NewRequest("PUT", url, buffer)
		if err != nil {
			return nil, err
		}

		client := &http.Client{}
		return client.Do(httpReq)
	}

	res, err := (putOperation)(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	byteRes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf(string(byteRes))
	return nil
}
