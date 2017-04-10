package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/topic", topicHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":4000", nil)
}

func rootHandler(res http.ResponseWriter, rq *http.Request) {
	log.Println("*** Request - webApplication to port 4000")
	io.WriteString(res, `<a href="/index">Enter</a>`)
}

func indexHandler(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

type City struct {
	name string
	desc string
}

func topicHandler(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/topic.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	seoul := City{"Seoul", "Kyeongki"}
	busan := City{"Busan", "KyeonNam"}
	daejeon := City{"Daejeon", "ChungNam"}
	osaka := City{"Osaka", "Japan-Asia"}
	vancouver := City{"Vancouver", "Canada-NA"}

	cities := []City{seoul, busan, daejeon, osaka, vancouver}

	tpl.ExecuteTemplate(res, "topic.gohtml", cities)
}
