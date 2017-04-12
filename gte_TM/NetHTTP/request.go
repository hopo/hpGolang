package main

import (
	"html/template"
	"net/http"
)

type hotdog int

func (m hotdog) ServerHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Pritln(err)
	}

	/*
			type Data struct {
				Submissions map[string][]string
			}
			data := Data{req.Form}
		    // equal code???
	*/
	data := struct {
		Submissions map[string][]string
	}{
		req.Form,
	}

	tpl.ExexuteTemplate(res, "request.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("request.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
