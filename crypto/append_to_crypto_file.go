package crypto

import (
	"bufio"
	"log"
	"os"
)

const pathToCSV = "./crypto/crypto_data.csv"

func RunAppendFUnction() {
	file, err := os.OpenFile(pathToCSV, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	row := "Alexcoin,ACE,542.21,38953.31\n"
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if _, err := writer.WriteString(row); err != nil {
		log.Fatal(err)
	}
}
