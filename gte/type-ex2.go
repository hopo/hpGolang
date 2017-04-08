// https://www.youtube.com/watch?v=QZ1dCAFA1ws&index=3&list=PLSak_q1UXfPrPJ3s7v43CMH9iMa4Dvh4X
// https://github.com/GoesToEleven/golang-web-dev/blob/master/000_temp/16_svcc/04/main.go

package main

import "fmt"

type Person struct {
    fname string
    lname string
}

type SecretAgent struct {
    Person
    licenseToKill bool
}

type Human interface {
    speak()
}

func (p Person) speak() {
    fmt.Println(p.fname, p.lname, `says, "bla bla bla"`)
}

func (sa SecretAgent) speak() {
    fmt.Println(sa.fname, sa.lname, `says, "I am hungry."`)
}

func saySomething(h Human) {
    h.speak()
}

func main() {
    p1 := Person{
        "Mr",
        "Anywords",
    }
    saySomething(p1)

    sa1 := SecretAgent{
        Person{
            "So",
            "Cratess",
        },
        true,
    }
    saySomething(sa1)
}
