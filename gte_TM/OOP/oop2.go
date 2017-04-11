package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("oop2.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Course: []course{
				course{"CSCI-40", "Introduction 40", "f"},
				course{"CSCI-130", "Introduction 30", "fa"},
				course{"CSCI-40", "Introduction 140", "d"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Course: []course{
				course{"CSCI-50", "Introduction 50", "f"},
				course{"CSCI-190", "Introduction 190", "f"},
				course{"CSCI-191", "Introduction 191", "f"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Course: []course{
				course{"Baseball", "Grand Slam!", "a"},
				course{"Soccer", "go Messi!", "s"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}

}
