package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("letter.doc")
	if err != nil {
		fmt.Println("*** There is error parsing\n", err)
	}

	friends := []string{"Harpa", "Amber", "Riceboy", "Amber"}

	err = tpl.Execute(os.Stdout, friends)
	if err != nil {
		fmt.Println("***There was an error executing template\n", err)
	}
}
