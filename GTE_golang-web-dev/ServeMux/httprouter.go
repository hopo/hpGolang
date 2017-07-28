// https://www.youtube.com/watch?v=eN-hZHz383k&list=PLSak_q1UXfPpXj-q1BeucvBAlNdotQWVD&index=47

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("doc/*"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/hello/:name", hello)
	http.ListenAndServe(":4000", r)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "httprouter.gohtml", nil)
	HandleError(res, err)
}

func hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

// HandleError func
func HandleError(res http.ResponseWriter, e error) {
	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}
}
