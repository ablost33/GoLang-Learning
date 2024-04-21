package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Boxer struct {
	Name  string  `json:"name"`
	Age   float64 `json:"age,string"`
	Email string  `json:"email"`
}

func (b *Boxer) UnmarshalJSON(data []byte) error {
	type Alias Boxer
	aux := &struct {
		Email string `json:"email"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	b.Email = strings.Repeat("*", 4) + "mail.com"
	return nil
}

func UnmarshalEmails() {
	file, err := os.Open("/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var boxer Boxer
	if err := json.Unmarshal(data, &boxer); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\n", boxer.Name)
	fmt.Printf("Age: %f\n", boxer.Age)
	fmt.Printf("Email: %s\n", boxer.Email)
}
