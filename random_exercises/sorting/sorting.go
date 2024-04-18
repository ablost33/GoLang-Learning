package sorting

import (
	"fmt"
	"sort"
)

/*
	Idea of exercise:
	1. Make a call to api
	2. Find a few piece in response to match using regex
	3. Put them in an array
	4. Write a custom sorter for them according to multiple criteria
*/

func BasicDataTypeSort() {
	slice := []int{4, 6, 78, 24, 1}
	sort.Ints(slice)
	fmt.Println("Integer sorting: ", slice)

	stringSlice := []string{"whois", "the", "goat"}
	sort.Strings(stringSlice)
	fmt.Println("String sorting: ", stringSlice)

	floatSlice := []float64{1.032, 2.094, 5.2414}
	sort.Float64s(floatSlice)
	fmt.Println("Float64 sorting: ", floatSlice)

	sort.SliceStable(slice, func(i, j int) bool {
		return i > j
	})
	fmt.Println("Integer reverse sorting: ", slice)
}

type Person struct {
	Age  int
	Name string
}

func CustomSort() {
	personSlice := []Person{
		{
			Name: "John Doe",
			Age:  12,
		},
		{
			Name: "John Dorian",
			Age:  16,
		},
		{
			Name: "Max Doe",
			Age:  25,
		},
		{
			Name: "Matt Doe",
			Age:  25,
		},
		{
			Name: "John Jacob",
			Age:  12,
		},
	}
	sort.SliceStable(personSlice, func(i, j int) bool {
		if personSlice[i].Age != personSlice[j].Age {
			return personSlice[i].Age < personSlice[j].Age
		}
		return personSlice[i].Name < personSlice[j].Name
	})
	fmt.Println("Print personSlice: ", personSlice)
}
