package csv_n_post

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	urlPath = "/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/csv_n_post/transactions.csv"
)

func CsvParsing2() error {
	file, err := os.Open(urlPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return err
	}

	expr, err := regexp.Compile("Groceries.?")
	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 1; i < len(rows); i++ {
		for _, field := range rows[i] {
			if expr.MatchString(field) {
				fmt.Println(field)
			}
		}
	}
	return nil
}
