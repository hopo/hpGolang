package main

import (
    "fmt"
    "os"
)

func main() {
    openFile("docu/1.txt")
    println("Done") // 이 문장 실행됨
}

func openFile(fn string) {
    // defere 함수. panic 호출시 실행됨
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()

    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }

    // 파일 close 실행됨
    defer f.Close()
}
