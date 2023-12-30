package csv_practice

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*	Exercise: Call an API endpoint, get the data, write it to a csv file
	- Try to incorporate some regex ?
*/

func WriteFile() {
	file, err := os.Create("./csv_practice/write_practice.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"name", "age", "gender"}
	data := [][]string{
		{"Alice", "25", "Female"},
		{"Bob", "30", "Male"},
		{"Charlie", "35", "Male"},
	}

	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}
}

func AppendToFile() {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	rowToAdd := []string{"Bobert", "55", "Male"}
	if err := writer.Write(rowToAdd); err != nil {
		panic(err)
	}
}
