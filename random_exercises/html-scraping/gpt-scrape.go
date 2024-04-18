package html_scraping

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func gptloadHTMLFile() ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func gptextractBookDetails(data []byte) ([][]string, error) {
	bookPattern := regexp.MustCompile(`(?s)<div class="book">.*?<h3 class="title">(.*?) - <span class="author">(.*?)</span></h3>.*?<span class="price">Price: (\$[\d\.]+).*?Discount: <span class="discount">(.*?)</span>.*?</div>`)
	matches := bookPattern.FindAllStringSubmatch(string(data), -1)

	var books [][]string
	for _, match := range matches {
		if len(match) == 5 {
			books = append(books, match[1:5])
		}
	}
	return books, nil
}

func gptprettyPrintResults(books [][]string) {
	for i, book := range books {
		fmt.Printf("%d. %s by %s, original price is: %s, but there is a discount of %s\n", i+1, book[0], book[1], book[2], book[3])
	}
}

func GPTmain() {
	data, err := gptloadHTMLFile()
	if err != nil {
		fmt.Println("Error loading HTML file:", err)
		return
	}

	books, err := gptextractBookDetails(data)
	if err != nil {
		fmt.Println("Error extracting book details:", err)
		return
	}

	gptprettyPrintResults(books)
}
