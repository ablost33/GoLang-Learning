package string_manipulation

import (
	"fmt"
	"strings"
)

func equality(word1 string, word2 string) {
	fmt.Println(strings.EqualFold(word1, word2))
}

func iterate(word string) {
	for _, ch := range word {
		fmt.Println("Word is: ", ch)
	}
}

func Replace() {
	fmt.Println(strings.Replace("k, i, l, l, e, r, s", ", ", "", -1))
	fmt.Println(strings.Replace("foobarb", "b", "f", 2))
}

func ModifyCharacterAtIndex(word string) {
	for _, c := range word {
		fmt.Println("char is ", string(c))
	}

	/* Notice this piece */
	byteArray := []byte(word)
	byteArray[2] = 'w'
	fmt.Println(string(byteArray))
}

func ExperimentWithSplit() {
	arr := "one, two, three, four"

	/* This output: [one  two  three  four] */
	fmt.Println("Split is: ", strings.Split(arr, ","))

	/* This output: [one,  two,  three,  four]*/
	fmt.Println("SplitAfter is: ", strings.SplitAfter(arr, ","))
}
