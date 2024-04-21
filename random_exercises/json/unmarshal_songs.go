package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	songPath = "/Users/alexblostein/go/src/github.com/ablost33/GoLang-Learning/random_exercises/json/songs.json"
)

type Song struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear int64  `json:"releaseYear,string"`
}

func UnmarshalSongs() {
	file, err := os.Open(songPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var songs []Song
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &songs); err != nil {
		log.Fatal(err)
	}

	for _, song := range songs {
		fmt.Printf("Title: %s, Artist: %s, ReleaseYear: %d\n", song.Title, song.Artist, song.ReleaseYear)
	}
}
