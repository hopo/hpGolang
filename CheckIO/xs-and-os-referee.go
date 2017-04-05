package main

import (
    "fmt"
)

func xoReferee(data []string) string {
    return "D" || "X" || "O"
}


func main() {
    fmt.Println(xoReferee([]string{
        "X.O",
        "XX.",
        "XOO"})) //"X", "Xs wins"

    fmt.Println(xoReferee([]string{
        "OO.",
        "XOX",
        "XOX"})) //"O", "Os wins"

    fmt.Println(xoReferee([]string{
        "OOX",
        "XXO",
        "OXX"}))    //"D", "Draw"

    fmt.Println(xoReferee([]string{
        "O.X",
        "XX.",
        "XOO"}))    //"X", "Xs wins again"
}
