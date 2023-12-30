package csv_practice

import (
	"fmt"
	"os"

	gocsv "github.com/gocarina/GoCsv"
)

type Person struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	Gender string `csv:"gender"`
}

const (
	filepath = "./csv_practice/write_practice.csv"
)

func GoCSVWriteToFile() {
	file, err := os.Create("./csv_practice/gocsv_practice.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	people := []*Person{
		{"Alice", 25, "Female"},
		{"Bob", 30, "Male"},
		{"Charlie", 35, "Male"},
	}

	if err := gocsv.MarshalFile(&people, file); err != nil {
		panic(err)
	}
}

func GoCSVReadFile() {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	person := []Person{}
	if err := gocsv.Unmarshal(file, &person); err != nil {
		fmt.Println(err)
	}

	for _, peep := range person {
		fmt.Println("Person is: ", peep)
	}
}
