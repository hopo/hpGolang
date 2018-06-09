package main

import (
	"fmt"
)

type Data struct {
	char    string
	natural int
	binary  string
}

func main() {
	dt := makeData(32, 128)
	dtPrnt(dt)
}

// differ between for and for
func makeData(a, b int) []Data {
	l := b - a
	dt := make([]Data, l)
	for i := 0; i < l; i++ {
		dt[i] = Data{
			char:    string(i + a),
			natural: i + a,
			binary:  fmt.Sprintf("%08b", i+a),
		}
	}
	return dt
}

func dtPrnt(dt []Data) {
	for i, _ := range dt {
		if dt[i].natural%8 == 0 {
			fmt.Println("")
		}
		fmt.Printf("<%v>\t%v,\t%v\n", dt[i].char, dt[i].natural, dt[i].binary)
	}
}
