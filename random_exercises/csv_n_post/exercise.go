package csv_n_post

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	pathToCsv = "./csv_n_post/transactions.csv"
)

func parseCSVFile() ([][]string, error) {
	file, err := os.Open(pathToCsv)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rawData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	csvEntries := strings.Split(string(rawData), "\n")

	var transactions [][]string
	for i := 1; i < len(csvEntries)-1; i++ {
		transactions = append(transactions, strings.Split(csvEntries[i], ","))
	}
	return transactions, nil
}

func Run() {
	data, err := parseCSVFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)
}
