// https://go-tour-kr.appspot.com/#58

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe("localhost:4000", nil)
}

func root(res http.ResponseWriter, rq *http.Request) {
	io.WriteString(res, "Welcome")
}
