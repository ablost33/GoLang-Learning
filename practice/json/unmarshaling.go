package json

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	filepath = "json/window.json"
)

type Token struct {
	Token string `json:"token"`
}

type Window struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}

func UnMarshalSimpleStructure() error {
	jsonInput := `{
        "width": 500,
        "height": 200,
        "title": "Hello Go!"
    }`
	var window Window
	err := json.Unmarshal([]byte(jsonInput), &window)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(window) // {500 200 Hello Go!}
	return nil
}

func UnMarshalIntoMap() error {
	jsonInput := `{
        "apples": 10,
        "mangos": 20,
        "grapes": 20
    }`
	output := make(map[string]int)
	err := json.Unmarshal([]byte(jsonInput), &output)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Map is: ", output)
	return nil
}

func UnMarshalArray() error {
	jsonInput := `[
	{
		"token": "33829u4kfn"
	},
	{
		"token": "33829u4kfn"
	}]`
	tokens := []Token{}
	err := json.Unmarshal([]byte(jsonInput), &tokens)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("tokens are: ", tokens)
	return nil
}

func UnMarshalFile() error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	windows := []Window{}
	err = json.Unmarshal(bytes, &windows)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Sprintf("Windows", windows)
	return nil
}
