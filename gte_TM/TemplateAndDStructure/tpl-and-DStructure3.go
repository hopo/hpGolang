package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp/tpl-map.gohtml"))
}

func main() {
	sages := map[string]string{
		"Ryk": "Harpa",
		"Tpe": "Amber",
		"Egl": "Riceboy",
		"Sel": "Sooki",
	}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
