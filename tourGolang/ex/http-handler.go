// https://go-tour-kr.appspot.com/#58

package main

import (
    "net/http"
)

func main() {
    // your http.Handle calls here
    http.ListenAndServe("localhost:4000", nil)
}
