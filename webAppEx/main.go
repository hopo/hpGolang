package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/topic", topic)
	http.ListenAndServe("127.0.0.1:4000", nil)
}

func index(res http.ResponseWriter, rq *http.Request) {
	io.WriteString(res, `<a href="/topic">Welcome</a>`)
}

func topic(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("temp/topic.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// tpl.Execute(res, "News")
	tpl.ExecuteTemplate(res, "topic.gohtml", "News")

}
