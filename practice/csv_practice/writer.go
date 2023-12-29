package csv_practice

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteFile() {
	file, err := os.Create("/csv_practice/write_practice.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

}
