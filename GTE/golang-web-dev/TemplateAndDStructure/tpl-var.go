package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("doc/tpl-var.gohtml"))
}

func main() {
	msg := `Release sel-focus; embrace other-focus.`
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-var.gohtml", msg)
	if err != nil {
		log.Fatalln(err)
	}
}
