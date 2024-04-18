package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// Used this: https://blog.logrocket.com/using-json-go-guide/

type Seller struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

type Product struct {
	Id     int    `json:"ud"`
	Name   string `json:"name"`
	Seller Seller `json:"seller"`
	Price  int    `json:"price"`
}

func MarshalingComplexObjects() error {
	book := Product{
		Id:     50,
		Name:   "Writing Book",
		Seller: Seller{1, "ABC Company", "US"},
		Price:  100,
	}
	bytes, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

func MarshalToFile() error {
	window := Window{
		Width:  150,
		Height: 200,
		Title:  "Californication",
	}

	jsonData, err := json.MarshalIndent(window, "", "\t")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(jsonData))

	err = os.WriteFile("json/window2.json", jsonData, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
