package endpoint

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/hello", helloHandler)
	fileServer := http.FileServer(http.Dir("./endpoint/static/form.html"))
	http.Handle("/", fileServer)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(w, "ParseForm() err: %v", err)
	}
}
