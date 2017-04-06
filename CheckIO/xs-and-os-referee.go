package main

import (
    "fmt"
)

func xoReferee(data []string) interface{} {
    var dt []string
    for i := 0; i < 3; i++ {
        for j := 0; j <3; j++ {
            dt = append(dt, string(data[i][j]))
        }
    }
    var os, xs int
    fmt.Println(dt)
    for _, d := range dt {
        if "." == d {
            if "O" == d {
                os++
            } else if "X" == d {
                xs++            
            }        
        }
    }
    if os > xs { return "O" }
    if os < xs { return "X" }
    return "center"
}

func main() {
    // fmt.Println(xoReferee([]string{"X.O",
    //                                "XX.",
    //                                "XOO"})) //"X", "Xs wins"
    fmt.Println(xoReferee([]string{"OO.",
                                   "XOX",
                                   "XOX"}))	//"O", "Os wins"
    fmt.Println(xoReferee([]string{"OOX",
                                   "XXO",
                                   "OXX"})) //"D", "Draw"
    // fmt.Println(xoReferee([]string{"O.X",
    //                                "XX.",
    //                                "XOO"})) //"X", "Xs wins again"
}