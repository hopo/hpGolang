package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp/tpl-slice.gohtml"))
}

func main() {
	sages := []string{"Harpa", "Amber", "Riceboy", "Sooki"}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
