package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintf(w, "Hello, gopher %s\n", match[0][1])
		return
	}
	fmt.Fprintf(w, "Hello, dear %s\n", path)
}
