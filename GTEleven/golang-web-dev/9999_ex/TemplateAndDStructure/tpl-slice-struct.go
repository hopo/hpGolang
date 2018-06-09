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
	tpl = template.Must(template.ParseFiles("doc/tpl-slice-struct.gohtml"))
}

func main() {
	pepper := sage{
		Name:  "Pepper",
		Motto: "I'm so Hot!",
	}
	carrot := sage{
		Name:  "Carrot",
		Motto: "Kill Horse.",
	}
	onion := sage{
		Name:  "Onion",
		Motto: "Watasi wa Tamanegi.",
	}

	sages := []sage{pepper, carrot, onion}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
