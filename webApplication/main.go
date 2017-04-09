package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/index", index)
	http.HandleFunc("/topic", topic)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":4000", nil)
}

func root(res http.ResponseWriter, rq *http.Request) {
	log.Println("*** Request - localhost:4000")
	io.WriteString(res, `<a href="/index">Enter</a>`)
}

func index(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func topic(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/topic.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// tpl.Execute(res, "News")
	tpl.ExecuteTemplate(res, "topic.gohtml", "(from main.go)")

}
