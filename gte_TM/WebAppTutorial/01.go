// https://www.youtube.com/watch?v=iIztjjNTSjs

package main

import (
	"io"
	"net/http"
)

/*
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}
*/

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "some text hello world")
}
