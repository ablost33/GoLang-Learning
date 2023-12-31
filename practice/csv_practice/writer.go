package csv_practice

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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

func ManualWriteToFile() {
	file, err := os.Create("./csv_practice/write_practice.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	headers := []string{"name", "age", "gender"}
	rawData := [][]string{
		{"Gordon", "25", "Female"},
		{"Dean", "30", "Male"},
		{"Dan", "35", "Male"},
	}

	modifiedHeaders := strings.Join(headers, ",") + "\n"
	_, err = io.WriteString(file, modifiedHeaders)
	if err != nil {
		panic(err)
	}

	for _, row := range rawData {
		modifiedRow := strings.Join(row, ",") + "\n"
		_, err := io.WriteString(file, modifiedRow)
		if err != nil {
			panic(err)
		}
	}
}

// ManualBufferedWriter Buffered writing is preferable because we make less calls to disk i/o since we accumulate it in the buffer
func ManualBufferedWriter() {
	file, err := os.Create("./csv_practice/write_practice.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	headers := []string{"name", "age", "gender"}
	rawData := [][]string{
		{"Gordon", "25", "Female"},
		{"Dean", "30", "Male"},
		{"Dan", "35", "Male"},
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(strings.Join(headers, ",") + "\n")
	if err != nil {
		panic(err)
	}

	for _, row := range rawData {
		_, err := writer.WriteString(strings.Join(row, ",") + "\n")
		if err != nil {
			panic(err)
		}
	}
}
