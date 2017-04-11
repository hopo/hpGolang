package main

import (
	"html/template"
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func ml(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("http.gohtml")
	if err != nil {
		log.Fatalln("error parsin template", err)
	}

	err = tpl.ExecuteTemplate(res, "http.gohtml", "ML")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/tm/", ml)
	http.ListenAndServe(":4000", nil)
}
