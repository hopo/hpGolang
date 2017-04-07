package main

import (
    "fmt"
)

func xoReferee(data []string) string {
    x, y, z := data[0], data[1], data[2]

    // for i := 0; i < 3; i++ {
    //     if x[i] == y[i] && y[i] == z[i] { return string(x[i]) }
    //     for _, D := range []string{x, y, z} {
    //         if D[0] == D[1] && D[1] == D[2] { return string(D[0]) }
    //     }
    // }
    //
    // if x[0] == y[1] && y[1] == z[2] { return string(x[0]) }
    // if x[2] == y[1] && y[1] == z[0] { return string(z[0]) }

    for i, D := range []string{x, y, z} {
        if D[0] == D[1] && D[1] == D[2] { return string(D[0]) }
        if x[0] == y[1] && y[1] == z[2] { return string(x[0]) }
        if x[2] == y[1] && y[1] == z[0] { return string(z[0]) }
        if x[i] == y[i] && y[i] == z[i] { return string(x[i]) }
    }

    return "D"
}

func main() {
    fmt.Println(xoReferee([]string{"X.O",
                                   "XX.",
                                   "XOO"})) //"X", "Xs wins"
    fmt.Println(xoReferee([]string{"OO.",
                                   "XOX",
                                   "XOX"}))	//"O", "Os wins"
    fmt.Println(xoReferee([]string{"OOX",
                                   "XXO",
                                   "OXX"})) //"D", "Draw"
    fmt.Println(xoReferee([]string{"O.X",
                                   "XX.",
                                   "XOO"})) //"X", "Xs wins again"
    fmt.Println(xoReferee([]string{"O.O ",
                                   "OO.",
                                   "XXX"})) //"X", "hp check"
}
