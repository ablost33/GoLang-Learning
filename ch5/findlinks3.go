package ch5

import (
	"GoLang-Learning/ch5/links"
	"fmt"
	"log"
	"os"
)

/*
	Implementation of BFS. f is called for each item in the worklist, and all items returned by
	f are added to the worklist. f is called at most once for each item.
*/
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// f(item)... causes all the items in the list returned by f to be appended to the worklist
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

/*
	Here's a small crawl function which we'll use to supply as input to breadthFirst
*/
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
