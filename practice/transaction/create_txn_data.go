package transaction

import (
	"bufio"
	"os"
	"strings"
)

const (
	filepath = "./transaction/txn_data.csv"
)

func CreateTransactionCSV() {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	headers := []string{"name", "date", "amount", "location"}
	data := [][]string{
		{"car repair", "2024-01-01", "79.99", "Montreal, Canada"},
		{"rent payment", "2024-01-01", "1000", "Montreal, Canada"},
		{"football tickets", "2024-01-11", "129.99", "Montreal, Canada"},
		{"subway sandwich", "2024-01-14", "5.99", "Montreal, Canada"},
		{"snowboard purchase", "2024-02-03", "239.99", "Stowe, Vermont"},
		{"car repair", "2024-02-01", "79.99", "Montreal, Canada"},
		{"rent payment", "2024-02-01", "1000", "Montreal, Canada"},
		{"gas", "2024-02-27", "82.42", "Montreal, Canada"},
		{"flowers", "2024-02-17", "9.99", "Montreal, Canada"},
		{"shoes", "2024-02-30", "256", "Montreal, Canada"},
		{"car repair", "2024-03-01", "79.99", "Montreal, Canada"},
		{"rent payment", "2024-03-01", "1000", "Montreal, Canada"},
		{"amazon.com", "2024-03-15", "2.99", "Montreal, Canada"},
		{"SAQ", "2024-03-22", "15.99", "Montreal, Canada"},
		{"gas", "2024-03-27", "82.42", "Montreal, Canada"},
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString(strings.Join(headers, ",") + "\n")
	for _, row := range data {
		writer.WriteString(strings.Join(row, ",") + "\n")
	}
}
