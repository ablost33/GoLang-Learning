package main

import (
	"os"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
}
