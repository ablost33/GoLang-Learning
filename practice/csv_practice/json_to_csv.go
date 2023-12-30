package csv_practice

import (
	"bytes"
	"encoding/json"
	"github.com/yukithm/json2csv"
	"os"
)

const (
	jsonFilePath   = "./csv_practice/input.json"
	outputFilePath = "./csv_practice/output.csv"
)

func JSONToCSV() {
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}
	buffer := &bytes.Buffer{}
	csvwriter := json2csv.NewCSVWriter(buffer)
	var dataMap []map[string]interface{}

	if err := json.Unmarshal(jsonData, &dataMap); err != nil {
		panic(err)
	}

	csv, err := json2csv.JSON2CSV(dataMap)
	if err != nil {
		panic(err)
	}

	if err := csvwriter.WriteCSV(csv); err != nil {
		panic(err)
	}
	defer csvwriter.Flush()
	got := buffer.String()

	file, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(got); err != nil {
		panic(err)
	}
}
