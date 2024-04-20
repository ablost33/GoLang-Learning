package json_to_csv

import (
	"bytes"
	"encoding/json"
	"github.com/yukithm/json2csv"
	"log"
	"os"
)

func JSONToCSV() {
	buffer := &bytes.Buffer{}
	writer := json2csv.NewCSVWriter(buffer)
	j, err := os.ReadFile("/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json_to_csv/input.json")
	if err != nil {
		log.Fatal(err)
	}

	var x []map[string]interface{}
	if err := json.Unmarshal(j, &x); err != nil {
		log.Fatal(err)
	}

	// 1. Convert json to csv; function returns byte slice containing csv data
	csv, err := json2csv.JSON2CSV(x)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Write csv data to a buffer
	if err := writer.WriteCSV(csv); err != nil {
		log.Fatal(err)
	}
	writer.Flush()

	// 3. Open the file to create
	f, err := os.Create("./output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 4. Write the buffer to the file
	if _, err := f.WriteString(buffer.String()); err != nil {
		log.Fatal(err)
	}
}
