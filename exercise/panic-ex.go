package main

import "fmt"

func foo() int{
    a := make([]int, 4)
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("Recovered from this error:", err)
        }
    }
    a[7] = 59 // panic: out of bounds!
}

func main() {
    fmt.Println(foo())
}
