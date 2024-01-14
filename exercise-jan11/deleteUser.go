package exercise_jan11

import (
	"fmt"
	"net/http"
)

const (
	deleteURL = "https://reqres.in/api/users/2"
)

func DeleteUser() error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	return nil
}
