package data_structures

import "fmt"

/*
	- Go over how to create infinite loop
	- while loops?
*/

func forever() {
	for {
		fmt.Println("Squishy")
	}
}

func forLooping() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileLooping() {
	i := -1
	for i < 0 {
		fmt.Println("In the loop")
	}
}

func aRangeLoop() {
	numbers := []string{"One", "Two", "Three"}
	for i, value := range numbers {
		fmt.Println("Index is %s", i, " and value is: %s", value)
	}
}

func breakAndContinue() {
	var i int32
	for {
		if i%2 == 0 {
			continue
		}
		if i == 0 {
			break
		}
	}
}
