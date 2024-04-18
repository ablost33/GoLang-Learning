package json

import (
	"encoding/json"
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
