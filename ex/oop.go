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
	tpl = template.Must(template.ParseFiles("doc/oop.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "1_Fall",
			Courses: []course{
				course{Number: "101", Name: "English", Units: "a"},
				course{Number: "102", Name: "Japanese", Units: "b"},
				course{Number: "103", Name: "Spanish", Units: "c"},
			},
		},
		Spring: semester{
			Term: "2_Spring",
			Courses: []course{
				course{Number: "201", Name: "Soccer", Units: "l"},
				course{Number: "205", Name: "Swimming", Units: "m"},
				course{Number: "206", Name: "Tennis", Units: "n"},
			},
		},
		Summer: semester{
			Term: "3_Summer",
			Courses: []course{
				course{Number: "315", Name: "Go", Units: "x"},
				course{Number: "333", Name: "JavaScript", Units: "y"},
				course{Number: "399", Name: "Python", Units: "z"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
