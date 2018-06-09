package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	doc := js.Global.Get("documnet")
	doc.Call("addEventListener",
		"click",
		func() {
			println("clicked")
		},
	)
}
