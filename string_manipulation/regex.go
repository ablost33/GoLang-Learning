package string_manipulation

import (
	"fmt"
	"regexp"
)

func MatchString(input, pattern string) {
	match, err := regexp.MatchString(pattern, input)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Match: ", match)
}

func PlayWithFindStringIndex(message, expr string) {
	pattern := regexp.MustCompile(expr)
	firstMatchIndex := pattern.FindStringIndex(message)
	fmt.Println("firstMatchIndex: ", subString(firstMatchIndex[0], firstMatchIndex[1], message))
}

func subString(start, end int, word string) string {
	strArr := []byte(word)
	strArr = strArr[start:end]
	return string(strArr)
}

func FindTheSubstring(message, expr string) {
	pattern := regexp.MustCompile(expr)
	substring := pattern.FindString(message)
	fmt.Println("Substring is: ", substring)
}

func FindAllSubstring(message, expr string) {
	pattern := regexp.MustCompile(expr)
	stringArr := pattern.FindAllString(message, -1)
	fmt.Println("Array is: ", stringArr)
}

func FindAllStringIndex(message, expr string) {
	pattern := regexp.MustCompile(expr)
	indices := pattern.FindAllStringIndex(message, -1)
	fmt.Println("Indices are: ", indices)
}

func Questionning() {
	pattern, err := regexp.Compile("[A-z]ork")
	if err != nil {
		fmt.Println(err)
	}
	question := "Do you love new york city?"
	fmt.Println("Question:", pattern.MatchString(question))
}

func Splitting() {
	pattern, err := regexp.Compile("guys|york")
	if err != nil {
		fmt.Println(err)
	}
	welcomeMessage := "Hello guys, welcome to new york city"
	strArr := pattern.Split(welcomeMessage, 1)
	fmt.Println("array is ", strArr)
}

func ReplaceMatchedContent() {
	pattern := regexp.MustCompile("welcome")
	welcomeMessage := "Hello guys, welcome to new york city"
	replaced := pattern.ReplaceAllStringFunc(welcomeMessage, func(s string) string {
		return "I send my love"
	})
	replaced2 := pattern.ReplaceAllString(welcomeMessage, "asdasda")
	fmt.Println(replaced)
	fmt.Println(replaced2)
}

func CaseInsensitiveMatching() {
	pattern := `(?i)hello`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	fmt.Println(r.MatchString("Hello world")) // true
	fmt.Println(r.MatchString("hello world")) // true
	fmt.Println(r.MatchString("HeLlO world")) // true
	fmt.Println(r.MatchString("HELLO world")) // true
	fmt.Println(r.MatchString("goodbye world"))
}
