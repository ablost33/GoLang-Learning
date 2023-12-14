package main

import "fmt"

/*
	Important array library functions:
	- How to delete
	- How to print all the elements of array quick
*/

func createFixedSizeArray(len int) [3]int {
	a := [...]int{1, 2, 3}
	a[2] = 69
	fmt.Println("Modified the array at index 1 to have value:", a[1])

	b := &a
	b[1] = 0

	// Notice you can pass an array directly into fmt.Println
	fmt.Println("a is:", a)
	fmt.Println(0, "b is:", b)

	return a
}

/*
	Key slice concepts:
	- a[start:end], ex: b := a[1:3], will create a slice from index start to index end-1
	- A slice does not own any data of its own. It is just a representation of the underlying array. See ex with theArray
*/
func createFlexibleSizeSlice() []float64 {
	a := []float64{1.0332, 1.2009, 1.43543, 4.432, 4.5}
	b := a[1:3]

	theArray := [3]float64{1, 2, 3}
	b = theArray[1:2]
	fmt.Println("theArray before b messed with it:", theArray)
	b[0] = 69
	fmt.Println("theArray after b messed with it:", theArray)

	// Inputs to make are, in order: type, length, capacity
	aTrueSlice := make([]int, 0, 0)
	aTrueSlice = append(aTrueSlice, 3)
	bTrueSlice := aTrueSlice
	bTrueSlice[0] = 2
	fmt.Println("a is:", aTrueSlice)
	fmt.Println("b is:", bTrueSlice)

	cTrueSlice := make([]int, len(aTrueSlice))
	copy(cTrueSlice, aTrueSlice)
	cTrueSlice[0] = 100
	fmt.Println("a again is:", aTrueSlice)
	fmt.Println("c is:", cTrueSlice)

	return a
}
