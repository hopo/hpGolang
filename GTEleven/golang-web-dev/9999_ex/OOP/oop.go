package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("doc/oop.gohtml"))
}

func main() {
	p1 := doubleZero{
		person:        person{"Ian Fleming", 56},
		LicenseToKill: false,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
