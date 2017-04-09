package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp/tpl-var.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-var.gohtml", `Release sel-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
