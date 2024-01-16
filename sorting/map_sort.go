package sorting

import (
	"fmt"
	"sort"
)

var populationMap = map[string]int64{
	"Australia": 8,
	"Qatar":     5,
	"Wales":     4,
	"Burundi":   3,
	"Guinea":    2,
}

type pair struct {
	key   string
	value int64
}

func SortKeys() {
	fmt.Println(populationMap)
	//countryArray := make([]string, len(populationMap))
	//for key := range populationMap {
	//	countryArray = append(countryArray, key)
	//}
	//sort.Strings(countryArray)
	//fmt.Println(countryArray)
}

func SortByValues() {
	pairSlice := make([]pair, len(populationMap))
	for k, v := range populationMap {
		pairSlice = append(pairSlice, pair{key: k, value: v})
	}
	sort.SliceStable(pairSlice, func(i, j int) bool {
		return pairSlice[i].value < pairSlice[j].value
	})
	for _, p := range pairSlice {
		fmt.Println("The key is: ", p.key, " with value: ", p.value)
	}
}
