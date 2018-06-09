// https://go-tour-kr.appspot.com/#58

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":4000", nil)
}

func rootHandler(res http.ResponseWriter, rq *http.Request) {
	io.WriteString(res, "Welcome")
}
