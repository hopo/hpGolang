package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("temp/tpl-struct.gohtml"))
}

func main() {
	pepper := sage{
		Name:  "Pepper",
		Motto: "I'm so Hot!",
	}

	err := tpl.Execute(os.Stdout, pepper)
	if err != nil {
		log.Fatalln(err)
	}
}
