// https://tour.golang.org/methods/18

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


//Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

//For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
