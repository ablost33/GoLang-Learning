package transaction

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type TransactionData struct {
	lastDate    string
	occurrences int64
}

type MonthlyExpenses struct {
	totalSpend  float64
	numberOfTxn float64
}

func loadTransactionData() [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	stringifiedData := string(bytes)
	rowArr := strings.Split(stringifiedData, "\n")

	var data [][]string
	for i := 1; i < len(rowArr)-1; i++ {
		data = append(data, strings.Split(rowArr[i], ","))
	}
	return data
}

func findRecurringTransaction(data [][]string) map[string]int64 {
	reoccurenceMap := make(map[string]TransactionData)

	for _, transaction := range data {
		transactionName := transaction[0]
		transactionDate := transaction[1]

		value, exists := reoccurenceMap[transaction[0]]
		if !exists {
			reoccurenceMap[transactionName] = TransactionData{
				lastDate:    transaction[1],
				occurrences: 1,
			}
			continue
		}

		if isNewMonth(transactionDate, value.lastDate) {
			value.lastDate = transactionDate
			value.occurrences++
			reoccurenceMap[transactionName] = value
		}
	}

	res := make(map[string]int64)
	for key, val := range reoccurenceMap {
		res[key] = val.occurrences
	}
	return res
}

func LoadAndAnalyze() {
	data := loadTransactionData()
	dataMap := findRecurringTransaction(data)
	fmt.Println("Recurring transactions are ", dataMap)
	avgSpend := findAverageMonthlySpend(data)
	fmt.Println("Average spend is: ", avgSpend)
}

func isNewMonth(currentDate, lastDate string) bool {
	currentYearMonth := currentDate[:7]
	lastYearMonth := lastDate[:7]
	return currentYearMonth > lastYearMonth
}

func findAverageMonthlySpend(data [][]string) map[string]string {
	monthlyExpenseMap := make(map[string]MonthlyExpenses)

	for _, txn := range data {
		transactionDate, _ := time.Parse("2006-01-02", txn[1])
		transactionMonth := transactionDate.Month().String()
		transactionAmount, _ := strconv.ParseFloat(txn[2], 64)

		val, exists := monthlyExpenseMap[transactionMonth]
		if !exists {
			monthlyExpenseMap[transactionMonth] = MonthlyExpenses{
				totalSpend:  transactionAmount,
				numberOfTxn: 1,
			}
			continue
		}

		val.numberOfTxn++
		val.totalSpend += transactionAmount
		monthlyExpenseMap[transactionMonth] = val
	}

	monthlyAverageMap := make(map[string]string)
	for month, monthlyExpenseData := range monthlyExpenseMap {
		averageTransactionAmount := monthlyExpenseData.totalSpend / monthlyExpenseData.numberOfTxn
		formattedAverage := fmt.Sprintf("%.2f", averageTransactionAmount)
		monthlyAverageMap[month] = formattedAverage
	}
	return monthlyAverageMap
}
