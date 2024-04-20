package json_to_csv

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	filepath     = "/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json_to_csv/output.csv"
	jsonfilepath = "/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json_to_csv/input.json"
)

func CSVToJSON() {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	var data []map[string]any
	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		m := make(map[string]interface{})
		width, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		m[headers[0]] = width
		m[headers[1]] = row[1]

		height, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		m[headers[2]] = height

		data = append(data, m)
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.OpenFile(jsonfilepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(csvFile)
	defer writer.Flush()

	if _, err := writer.Write(byteData); err != nil {
		log.Fatal(err)
	}
}
