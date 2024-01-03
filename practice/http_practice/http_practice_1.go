package http_practice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	GETURL1     = "https://jsonplaceholder.typicode.com/posts/1"
	URL2        = "https://reqres.in/api/users/"
	contentType = "application/json"

	POSTURL1 = "https://postman-echo.com/post"
)

// Reference: https://blog.logrocket.com/making-http-requests-in-go/
/* @TODO: Practice cleaning data using the following link: https://www.makeuseof.com/go-make-http-requests/		ex: Remove the whitespaces, or commas */

func httpGET() error {
	resp, err := http.Get(GETURL1)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cleanResponse := string(body)
	fmt.Println("GET response is: %s", cleanResponse)

	return nil
}

func httpGETWithClient() error {
	req, err := http.NewRequest("GET", GETURL1, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	httpRes, err := client.Do(req)
	if err != nil {
		return err
	}

	byteResponse, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	response := string(byteResponse)
	fmt.Println("response is:\n", response)

	return nil
}

/*
	Follow-ups:

- Check what creating a struct and marshaling would look like
- Try parsing/extracting specific parts of the response (ex: data)
*/
func httpPOST() error {
	data := make(map[string]string)
	data["name"] = "Toby"
	data["email"] = "toby@example.com"

	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(reqBody)

	httpRes, err := http.Post(POSTURL1, contentType, buffer)
	if err != nil {
		return err
	}

	byteRes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	cleanResponse := string(byteRes)
	fmt.Println("Response is:\n", cleanResponse)

	return nil
}

func httpPOSTWithClient() error {
	data := make(map[string]string)
	data["name"] = "John"
	data["email"] = "john@mail.com"

	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}
	buffer := bytes.NewReader(reqBody)

	httpReq, err := http.NewRequest("POST", POSTURL1, buffer)
	if err != nil {
		return err
	}

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)

	byteRes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	strRes := string(byteRes)
	fmt.Println("response is:\n", strRes)

	return nil
}
