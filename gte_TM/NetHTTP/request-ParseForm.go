// https://www.youtube.com/watch?v=Ppw5UluP2R8&index=41&list=PLSak_q1UXfPpXj-q1BeucvBAlNdotQWVD

package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	/*
			type Data struct {
				Submissions map[string][]string
			}
			data := Data{req.Form}
		    // equal code???
	*/
	data := struct {
		Submissions url.Values //map[string][]string
	}{
		req.Form,
	}

	tpl.ExecuteTemplate(res, "request-ParseForm.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("doc/request-ParseForm.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":4000", h)
}
