package exercise_jan11

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	timeURL = "http://worldtimeapi.org/api/ip"
)

type CurrentTime struct {
	DateTime time.Time `json:"datetime,string"`
}

func GetTime() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", timeURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var currTime CurrentTime
	if err := json.Unmarshal(byteResponse, &currTime); err != nil {
		log.Fatal(err)
	}

	fmt.Println(currTime)
}
