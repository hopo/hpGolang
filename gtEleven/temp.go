// https://www.youtube.com/watch?v=RhvOQdJZ_EE&index=4&list=PLSak_q1UXfPrPJ3s7v43CMH9iMa4Dvh4X&t=2s

package main

import (
    "net/http"
    "io"
    "html/template"
    "log"
)
func main() {
    http.HandleFunc("/", foo)
    http.HandleFunc("/dog", bar)
    http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, rq *http.Request) {
    io.WriteString(res, "foo ran!")
}

func bar(res http.ResponseWriter, rq *http.Request) {
    tpl, err := template.ParseFiles("index.gohtml")
    if err != nil {
        log.Fatalln(err)
    }
    tpl.ExecuteTemplate(res, "index.gohtml", 27)
}
