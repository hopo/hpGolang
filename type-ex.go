package main

import "fmt"

type Person struct {
    fname string
    lname string
}

func main() {
    x := []int{7, 14, 12, 42}
    fmt.Println(x)

    y := map[string]int{"Harpa": 30, "Amber": 25}
    fmt.Println(y)
    fmt.Println(y["Harpa"])

    z := Person {
        fname: "Harpa",
        lname: "Reyk",
    }
    fmt.Println(z)
    fmt.Println(z.fname)
}
