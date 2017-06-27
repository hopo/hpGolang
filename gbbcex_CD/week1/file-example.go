package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	srcFile, err := os.Open("docs/hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", line)
	}
}
