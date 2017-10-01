package main

import (
	"fmt"
)

type Data struct {
	network      [][]string
	power_plants map[string]int
}

func powerSupply(d Data) (ret []string) {
	return
}

// powerSupply(d01).sort()
func main() {
	d01 := Data{
		[][]string{
			{"p1", "c1"},
			{"c1", "c2"}},
		map[string]int{"p1": 1},
	}
	fmt.Println(powerSupply(d01)) // ['c2'], "one blackout"

	d02 := Data{
		[][]string{
			{"c0", "c1"},
			{"c1", "p1"},
			{"c1", "c3"},
			{"p1", "c4"}},
		map[string]int{"p1": 1},
	}
	fmt.Println(powerSupply(d02)) // ['c0', 'c3'], "two blackout"

	d03 := Data{
		[][]string{
			{"p1", "c1"},
			{"c1", "c2"},
			{"c2", "c3"}},
		map[string]int{"p1": 3},
	}
	fmt.Println(powerSupply(d03)) // [], "no blackout"

	d04 := Data{
		[][]string{
			{"c0", "p1"},
			{"p1", "c2"}},
		map[string]int{"p1": 0},
	}
	fmt.Println(powerSupply(d04)) // ['c0', 'c2'].sort(), "weak power-plant"

	d05 := Data{
		[][]string{
			{"p0", "c1"},
			{"p0", "c2"},
			{"c2", "c3"},
			{"c3", "p4"},
			{"p4", "c5"}},
		map[string]int{"p0": 1, "p4": 1},
	}
	fmt.Println(powerSupply(d05)) // [], "cooperation"

	d06 := Data{
		[][]string{
			{"c0", "p1"},
			{"p1", "c2"},
			{"c2", "c3"},
			{"c2", "c4"},
			{"c4", "c5"},
			{"c5", "c6"},
			{"c5", "p7"}},
		map[string]int{"p1": 1, "p7": 1},
	}
	fmt.Println(powerSupply(d06)) // ['c3', 'c4', 'c6'].sort(), "complex cities 1"

	d07 := Data{
		[][]string{
			{"p0", "c1"},
			{"p0", "c2"},
			{"p0", "c3"},
			{"p0", "c4"},
			{"c4", "c9"},
			{"c4", "c10"},
			{"c10", "c11"},
			{"c11", "p12"},
			{"c2", "c5"},
			{"c2", "c6"},
			{"c5", "c7"},
			{"c5", "p8"}},
		map[string]int{"p0": 1, "p12": 4, "p8": 1},
	}
	fmt.Println(powerSupply(d07)) // ['c6', 'c7'].sort(), "complex cities 2"

	d08 := Data{
		[][]string{
			{"c1", "c2"},
			{"c2", "c3"}},
		map[string]int{},
	}
	fmt.Println(powerSupply(d08)) // ['c1', 'c2', 'c3'].sort(), "no power plants"

	d09 := Data{
		[][]string{
			{"p1", "c2"},
			{"p1", "c4"},
			{"c4", "c3"},
			{"c2", "c3"}},
		map[string]int{"p1": 1},
	}
	fmt.Println(powerSupply(d09)) // ['c3'], "circle"

	d10 := Data{
		[][]string{
			{"p1", "c2"},
			{"p1", "c4"},
			{"c2", "c3"}},
		map[string]int{"p1": 4},
	}
	fmt.Println(powerSupply(d10)) // [], "more than enough"
}
