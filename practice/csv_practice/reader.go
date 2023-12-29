package csv_practice

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*
	Ressource: https://earthly.dev/blog/golang-csv-files/
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
