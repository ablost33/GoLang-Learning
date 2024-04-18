package html_scraping

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

const (
	filepath = "./html-scraping/bookstore.html"
)

func loadHTMLFile() ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rawData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}

func extractBookTitles(data []byte) ([]string, error) {
	bookFinder, err := regexp.Compile("<div class=\"book\">\n    <h3 class=\"title\">(.*) -")
	if err != nil {
		return nil, err
	}

	matches := bookFinder.FindAllStringSubmatch(string(data), -1)

	var cleanTitles []string
	for _, match := range matches {
		if len(match) > 1 {
			cleanTitles = append(cleanTitles, match[1])
		}
	}
	return cleanTitles, nil
}

func extractAuthors(data []byte) ([]string, error) {
	authorFinder, err := regexp.Compile("<span class=\"author\">(.*)</span>")
	if err != nil {
		return nil, err
	}

	matches := authorFinder.FindAllStringSubmatch(string(data), -1)

	var authorArray []string
	for _, location := range matches {
		if len(matches) > 1 {
			authorArray = append(authorArray, location[1])
		}
	}
	return authorArray, nil
}

func extractPrices(data []byte) ([]string, error) {
	priceFinder, err := regexp.Compile("<span class=\"price\">Price: (\\$\\d+\\.?\\d{0,2})")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	matches := priceFinder.FindAllStringSubmatch(string(data), -1)

	var prices []string
	for _, match := range matches {
		if len(match) > 1 {
			prices = append(prices, match[1])
		}
	}
	return prices, nil
}

func extractDiscounts(data []byte) ([]string, error) {
	discountFinder, err := regexp.Compile("<span class=\"discount\">(\\d{0,2}%)")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	matches := discountFinder.FindAllStringSubmatch(string(data), -1)

	var discounts []string
	for _, match := range matches {
		if len(match) > 1 {
			discounts = append(discounts, match[1])
		}
	}
	return discounts, nil
}

func ParseWebPage() {
	data, err := loadHTMLFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	titles, err := extractBookTitles(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	authors, err := extractAuthors(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prices, err := extractPrices(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	discounts, err := extractDiscounts(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prettyPrintResults(titles, authors, prices, discounts)
}

func prettyPrintResults(titles, authors, prices, discounts []string) {
	for i := 0; i < len(titles); i++ {
		fmt.Println(titles[i], "by ", authors[i], "original price is: ", prices[i], "but there is a discount of", discounts[i])
	}
}
