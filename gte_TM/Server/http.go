package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func tm(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("doc/http.gohtml")
	if err != nil {
		log.Fatalln("error parsin template", err)
	}

	// res.Header().Set("content-type", "text/plain; chatset=utf-8")

	err = tpl.ExecuteTemplate(res, "http.gohtml", "TM")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/tm/", tm)
	http.ListenAndServe(":4000", nil)
}
