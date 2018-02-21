package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	ex_multiwriter()
	// ex_newreader()
	// ex_iopipe()
	// ex_reader()
}

func ex_multiwriter() {
	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)

	fmt.Fprintln(mw, "hello")
	fmt.Println("from buffer: ", buf)
}

func ex_newreader() {
	// https://www.youtube.com/watch?v=LHZ2CAZE6Gs
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	footer := strings.NewReader("</msg>")
	// io.MultiReader(header, body, footer)

	for _, r := range []io.Reader{header, body, footer} {
		_, err := io.Copy(os.Stdout, r)
		echk(err)
	}
}

func ex_iopipe() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "hello")
		echk(err)
	}()

	_, err := io.Copy(os.Stdout, pr)
	echk(err)
}

func ex_reader() {
	src := "example source string"
	reader := strings.NewReader(src)
	fmt.Printf("src: %v\n", src)
	fmt.Printf("reader: %v\n", reader)

	b, _ := ioutil.ReadAll(reader)
	fmt.Printf("b: %v\n", b)

	read_string := string(b)
	fmt.Printf("read_string: %v\n", read_string)

	if src == read_string {
		fmt.Println("*** src == read_string")
	}
}

func echk(err error) {
	if err != nil {
		panic(err)
	}
}
