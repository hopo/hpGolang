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
	tpl = template.Must(template.ParseFiles("temp/tpl-slice-struct.gohtml"))
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
	cars := []string{"stella", "tico", "poter", "hummer"}

	type data struct {
		Sages []sage
		Cars  []string
	}

	d := data{
		Sages: sages,
		Cars:  cars,
	}

	err := tpl.Execute(os.Stdout, d)
	if err != nil {
		log.Fatalln(err)
	}
}
