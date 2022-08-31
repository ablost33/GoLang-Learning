package ch5

import (
	"fmt"
	"net/http"
	"strings"

	"./html"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

/*
	Here we explore the usage of deferred functions. Deferred functions essentially guarantee that
	the function being deferred gets called after the function that contains the defer statement
	has finished.
*/
func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
}

/*
	Some more experimentation with deferred function calls; Here we'll use defer
	statement to make the function print its arguments and results each time
	it gets called
*/
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

/*
	Now let's use deferred anonymous function to change the values the enclosing
	function returns to its caller
*/
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
