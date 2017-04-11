package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type city struct {
	Num   int
	Title string
	Desc  string
}

func main() {
	http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/topic", topicHandler)
	// http.HandleFunc("/topic/:Num", topicIdHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":4000", nil)
}

func rootHandler(res http.ResponseWriter, rq *http.Request) {
	io.WriteString(res, `<a href="/topic">Enter Topic</a>`)
}

/*
func indexHandler(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}
*/

func topicHandler(res http.ResponseWriter, rq *http.Request) {
	tpl, err := template.ParseFiles("doc/topic.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	seoul := city{1, "Seoul", "Kyeongki"}
	busan := city{2, "Busan", "KyeonNam"}
	daejeon := city{3, "Daejeon", "ChungNam"}
	osaka := city{4, "Osaka", "Japan-Asia"}
	vancouver := city{5, "Vancouver", "Canada-NorthAmerica"}

	c := []city{seoul, busan, daejeon, osaka, vancouver}

	tpl.ExecuteTemplate(res, "topic.gohtml", c)
}
