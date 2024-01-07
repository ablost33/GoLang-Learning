package html_scraping

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
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
	bookFinder, err := regexp.Compile("<div class=\"book\">\n    <h3 class=\"title\">.* -")
	if err != nil {
		return nil, err
	}

	var cleanTitles []string
	titleIndexArray := bookFinder.FindAllStringIndex(string(data), -1)
	for _, pair := range titleIndexArray {
		title := string(data[pair[0]:pair[1]])
		title = strings.Trim(title, "<div class=\"book\">\n    <h3 class=\"title\">")
		title = strings.Trim(title, " -")
		cleanTitles = append(cleanTitles)
	}
	return cleanTitles, nil
}

func extractAuthors(data []byte) ([]string, error) {
	authorFinder, err := regexp.Compile("<span class=\"author\">([A-z][a-zA-Z'-]*\\.? )*[A-Z][a-zA-Z`-]*")
	if err != nil {
		return nil, err
	}
	authorIndexArray := authorFinder.FindAllStringIndex(string(data), -1)

	return nil, nil
}

func ParseWebPage() {
	data, err := loadHTMLFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = extractBookTitles(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	extractAuthors(data)
}
