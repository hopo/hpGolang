package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.ExpandEnv("$USER lives in ${HOME}."))
}
