package main

import (
	html_scraping "github.com/ablost33/GoLang-Learning/html-scraping"
)

/* @TODO: Answer the following
1. Do I need my project path to be the same as the github path?
2. Do I need to initialize all my projects in the go directory? I'm inclined to say no because of the go mod tool
3. See this link: https://blog.logrocket.com/making-http-requests-in-go/
*/

/*
- go build main.go
This will output an executable file. If we want fast code for users to interact with we'd use the executable file. Run go build
from the directory (/pkg) you're trying to build

- go run main.go
This will combine both the compilation and execution of code. But this won't work for multiple files. if you have multiple files with
interdependencies, do: go run main.go arrays_and_slices.go
*/

/*	====================== FMT basics ======================
	fmt.Println():
	fmt.Printf(format string, a ...any): formats according to a specifier & print to stdout. ex: fmt.Printf("%d print this: %.2f", 69.70)
*/

func main() {
	html_scraping.GPTmain()
}
