package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"-"`
}

func (p *Person) MarshalJSON() ([]byte, error) {
	type PersonAlias Person
	return json.Marshal(&struct {
		*PersonAlias
		Email string `json:"email"`
	}{
		PersonAlias: (*PersonAlias)(p),
		Email:       strings.Repeat("*", 4) + "@mail.com",
	})
}

func MarshalPerson() error {
	person := &Person{
		Name:  "Alex Apply",
		Age:   12,
		Email: "alex@mail.com",
	}
	data, err := person.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("data is: ", string(data))
	return nil
}
