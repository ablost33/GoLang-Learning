package csv_practice

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Resource: https://earthly.dev/blog/golang-csv-files/
*/

func ReadFile() {
	file, err := os.Open("/Users/alexblostein/desktop/practice.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()

	for _, record := range data {
		fmt.Println("Record is: ", record)
	}
}

func ManualRead() [][]string {
	file, err := os.Open("./csv_practice/gocsv_practice.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	stringifiedData := string(byteData)
	lines := strings.Split(stringifiedData, "\n")

	var rows [][]string
	for i := 1; i < len(lines)-1; i++ {
		rows = append(rows, strings.Split(lines[i], ","))
	}

	printRows(rows)
	return rows
}

func printRows(data [][]string) {
	for _, row := range data {
		fmt.Println(row)
	}
}
