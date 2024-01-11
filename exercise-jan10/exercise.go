package exercise_jan10

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	filepath = "./exercise-jan10/csv_sales_report_inconsistent.csv"
)

func readCSV() ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	rawData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rowArray := strings.Split(string(rawData), "\n")

	var cleanData [][]string
	for i := 1; i < len(rowArray)-1; i++ {
		if isCleanRow(rowArray[i]) {
			cleanData = append(cleanData, strings.Split(rowArray[i], ","))
		}
	}
	return cleanData, nil
}

func isCleanRow(row string) bool {
	rowPattern := regexp.MustCompile(`\d+,[^,]+,[\S,][^,]+,\d+,\d+`)
	return rowPattern.MatchString(row)
}

func revenuePerCategory(data [][]string) (map[string]float64, error) {
	result := make(map[string]float64)
	for _, transaction := range data {
		category := transaction[2]
		revenue, err := getTransactionRevenue(transaction[3], transaction[4])
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if _, contains := result[category]; !contains {
			result[category] = revenue
			continue
		}
		result[category] += revenue
	}

	prettyPrintResult(result)
	return result, nil
}

func getTransactionRevenue(quantity, price string) (float64, error) {
	quantityFloat, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return 0, err
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, err
	}

	return quantityFloat * priceFloat, nil
}

func Run() {
	data, err := readCSV()
	if err != nil {
		os.Exit(1)
	}
	revenuePerCategory(data)
}

func prettyPrintResult(revenuePerCategory map[string]float64) {
	for category, revenue := range revenuePerCategory {
		fmt.Printf("%s: %.2f\n", category, revenue)
	}
}
