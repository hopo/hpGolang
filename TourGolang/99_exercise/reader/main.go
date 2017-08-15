// https://tour.golang.org/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}


//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
