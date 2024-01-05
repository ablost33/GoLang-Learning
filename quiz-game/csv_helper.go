package quiz_game

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rawData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	stringifiedData := string(rawData)
	return cleanData(stringifiedData)
}

func cleanData(data string) ([][]string, error) {
	rows := strings.Split(data, "\n")

	var questionsAndAnswers [][]string
	for i := 1; i < len(rows)-1; i++ {
		rowArray := strings.Split(rows[i], ",")
		if len(rowArray) < 2 {
			return nil, fmt.Errorf("invalid csv. Received row with only question")
		}

		var cleanRow []string
		if len(rowArray) > 2 {
			answer := rowArray[len(rowArray)-1]
			question := rowArray[0 : len(rowArray)-1]
			joinedQuestion := strings.Join(question, ",")
			cleanRow = append(cleanRow, joinedQuestion)
			cleanRow = append(cleanRow, answer)
			questionsAndAnswers = append(questionsAndAnswers, cleanRow)
			continue
		}
		cleanRow = append(cleanRow, rowArray[0])
		cleanRow = append(cleanRow, rowArray[1])
		questionsAndAnswers = append(questionsAndAnswers, cleanRow)
	}
	return questionsAndAnswers, nil
}
