package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) serveHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you wany in this func")
}

func main() {

}
